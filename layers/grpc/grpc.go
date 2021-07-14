package grpc

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/danhspe/fizz-buzz-rest-server/golib/fizzbuzz"
	"github.com/danhspe/fizz-buzz-rest-server/layers/usecases"
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

type grpcServer struct {
	fizzBuzzUseCases   usecases.FizzBuzz
	statisticsUseCases usecases.Statistics
}

var _ fizzbuzz.FizzBuzzServiceServer = (*grpcServer)(nil)

func NewFizzBuzzServiceServer(fizzBuzzUseCases usecases.FizzBuzz, statisticsUseCases usecases.Statistics) fizzbuzz.FizzBuzzServiceServer {
	return &grpcServer{fizzBuzzUseCases: fizzBuzzUseCases, statisticsUseCases: statisticsUseCases}
}

func (s *grpcServer) GetFizzBuzz(ctx context.Context, request *fizzbuzz.FizzBuzzRequest) (*fizzbuzz.FizzBuzzResponse, error) {
	args := arguments.New(int(request.Int1), int(request.Int2), int(request.Limit), request.Str1, request.Str2)
	log.Printf("GetFizzBuzz arguments: %+v", args)

	fizzBuzz := s.fizzBuzzUseCases.GetFizzBuzz(args)
	return &fizzbuzz.FizzBuzzResponse{
		Result: fizzBuzz,
	}, nil
}

func (s *grpcServer) GetStatistics(ctx context.Context, empty *emptypb.Empty) (*fizzbuzz.StatisticsResponse, error) {
	highestScore, mostFrequentArguments := s.statisticsUseCases.GetStatistics()

	var fizzBuzzRequests []*fizzbuzz.FizzBuzzRequest

	for _, argument := range mostFrequentArguments {
		request := fizzbuzz.FizzBuzzRequest{
			Int1:  int64(argument.Int1),
			Int2:  int64(argument.Int2),
			Limit: int64(argument.Limit),
			Str1:  argument.Str1,
			Str2:  argument.Str2,
		}
		fizzBuzzRequests = append(fizzBuzzRequests, &request)
	}

	return &fizzbuzz.StatisticsResponse{
		HighestScore: int64(highestScore),
		Requests:     fizzBuzzRequests,
	}, nil
}
