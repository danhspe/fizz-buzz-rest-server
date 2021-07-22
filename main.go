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
	"github.com/danhspe/fizz-buzz-rest-server/internal/cache"
	"github.com/danhspe/fizz-buzz-rest-server/internal/cache/redis"
	grpc2 "github.com/danhspe/fizz-buzz-rest-server/layers/grpc"
	fizzbuzz3 "github.com/danhspe/fizz-buzz-rest-server/layers/repositories/fizzbuzz"
	"github.com/danhspe/fizz-buzz-rest-server/layers/repositories/statistics"
	fizzbuzz2 "github.com/danhspe/fizz-buzz-rest-server/layers/usecases/fizzbuzz"
	statistics2 "github.com/danhspe/fizz-buzz-rest-server/layers/usecases/statistics"
)

const (
	defaultHttpEndpoint  = ":8080"
	defaultGrpcEndpoint  = ":58080"
	defaultRedisEndpoint = "localhost:6379"
)

const retryTimeout = time.Second * 5

var (
	dataCache     cache.Cache
	redisEndpoint *string
)

func init() {
	log.SetPrefix("fizz-buzz-rest-server ")
	initDefaults()
}

func initDefaults() {
	redisEndpoint = flag.String("redisEndpoint", defaultRedisEndpoint, "Redis endpoint host:port")
	flag.Parse()
	log.Printf("Redis endpoint: %s", *redisEndpoint)
}

// initCache returns a redis client and optionally waits until the connection has been established
func initCache(address string, retry bool) cache.Cache {
	redisCache := redis.NewRedisCache(address)
	if redisCache.Connect() != nil {
		log.Printf("Warning: Redis endpoint not reachable at %s", address)
	}
	for retry && redisCache.Connect() != nil {
		log.Printf("Waiting for Redis endpoint at %s", address)
		time.Sleep(retryTimeout)
	}
	return redisCache
}

func main() {
	dataCache = initCache(*redisEndpoint, true)

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
