package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/danhspe/fizz-buzz-rest-server/golib/fizzbuzz"
	grpc2 "github.com/danhspe/fizz-buzz-rest-server/internal/layers/api/grpc"
	fizzbuzz3 "github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories/fizzbuzz"
	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories/statistics"
	fizzbuzz2 "github.com/danhspe/fizz-buzz-rest-server/internal/layers/usecases/fizzbuzz"
	statistics2 "github.com/danhspe/fizz-buzz-rest-server/internal/layers/usecases/statistics"
	"github.com/danhspe/fizz-buzz-rest-server/internal/storage/cache"
	"github.com/danhspe/fizz-buzz-rest-server/internal/storage/cache/redis"
)

const (
	defaultHttpEndpoint       = ":8080"
	defaultGrpcEndpoint       = ":58080"
	defaultRedisEndpoint      = "localhost:6379"
	defaultShouldWaitForRedis = false
)

const retryTimeout = time.Second * 5

var (
	dataCache     cache.Cache
	redisEndpoint *string
	waitForRedis  *bool
)

func init() {
	log.SetPrefix("fizz-buzz-rest-server ")
	initDefaults()
}

func initDefaults() {
	redisEndpoint = flag.String("redisEndpoint", defaultRedisEndpoint, "Redis endpoint host:port")
	waitForRedis = flag.Bool("waitForRedis", defaultShouldWaitForRedis, "Wait for Redis to be ready")
	flag.Parse()
	log.Printf("Redis endpoint: %s", *redisEndpoint)
	log.Printf("Wait for Redis: %t", *waitForRedis)
}

// initCache returns a redis client and optionally waits until the connection has been established
func initCache(address string, waitForRedis bool) cache.Cache {
	redisCache := redis.NewRedisCache(address)
	if redisCache.Connect() != nil {
		log.Printf("Warning: Redis endpoint not reachable at %s", address)
	}
	for waitForRedis && redisCache.Connect() != nil {
		log.Printf("Waiting for Redis endpoint at %s", address)
		time.Sleep(retryTimeout)
	}
	return redisCache
}

func main() {
	dataCache = initCache(*redisEndpoint, *waitForRedis)

	//
	// start gRPC server in background
	//

	listener, err := net.Listen("tcp", defaultGrpcEndpoint)
	if err != nil {
		log.Fatalf("failed to create gRPC listener: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	fizzbuzz.RegisterFizzBuzzServiceServer(grpcServer, grpc2.NewFizzBuzzServiceServer(
		fizzbuzz2.NewFizzBuzzUseCase(fizzbuzz3.NewFizzBuzzRepository(dataCache)),
		statistics2.NewStatisticsUseCases(statistics.NewStatisticsRepository(dataCache))),
	)

	go func() {
		log.Printf("gRPC server listening at %s", listener.Addr().String())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to start gRPC server: %s", err.Error())
		}
	}()

	//
	// start gRPC Gateway
	//

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts2 := []grpc.DialOption{grpc.WithInsecure()}
	if err = fizzbuzz.RegisterFizzBuzzServiceHandlerFromEndpoint(ctx, mux, defaultGrpcEndpoint, opts2); err != nil {
		log.Fatalf("Failed to register service handler: %v", err)
	}

	log.Printf("HTTP server listening at %s", defaultHttpEndpoint)
	log.Fatal(http.ListenAndServe(defaultHttpEndpoint, mux))
}
