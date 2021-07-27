package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/danhspe/fizz-buzz-rest-server/golib/fizzbuzz"
	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/api"
	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/usecases"
	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
)

type grpcServer struct {
	fizzBuzzUseCases   usecases.FizzBuzz
	statisticsUseCases usecases.Statistics
}

var _ api.GRPC = (*grpcServer)(nil)

func NewFizzBuzzServiceServer(fizzBuzzUseCases usecases.FizzBuzz, statisticsUseCases usecases.Statistics) api.GRPC {
	return &grpcServer{fizzBuzzUseCases: fizzBuzzUseCases, statisticsUseCases: statisticsUseCases}
}

func (s *grpcServer) GetFizzBuzz(ctx context.Context, request *fizzbuzz.FizzBuzzRequest) (*fizzbuzz.FizzBuzzResponse, error) {

	args := arguments.New(int(request.Int1), int(request.Int2), int(request.Limit), request.Str1, request.Str2)
	log.Printf("GetFizzBuzz arguments: %+v", args)

	fizzBuzz, err := s.fizzBuzzUseCases.GetFizzBuzz(args)
	if err != nil {
		if err == usecases.ErrWrongFizzBuzzArguments {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("%s: %+v", err.Error(), args))
		} else if err == usecases.ErrSaveFizzBuzzArguments {
			return nil, status.Error(codes.Internal, "failed to save arguments for statistics")
		}
	}

	return &fizzbuzz.FizzBuzzResponse{
		Result: fizzBuzz,
	}, nil
}

func (s *grpcServer) GetStatistics(ctx context.Context, empty *emptypb.Empty) (*fizzbuzz.StatisticsResponse, error) {

	highestScore, mostFrequentArguments, err := s.statisticsUseCases.GetStatistics()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get statistics")
	}

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
