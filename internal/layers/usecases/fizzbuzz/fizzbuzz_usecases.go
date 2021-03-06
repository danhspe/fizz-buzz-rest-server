package fizzbuzz

import (
	"log"
	"strings"

	"github.com/danhspe/fizz-buzz-rest-server/internal/fizzBuzz"
	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/usecases"
	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
)

type fizzBuzzUseCase struct {
	repository repositories.FizzBuzzRepository
}

var _ usecases.FizzBuzz = (*fizzBuzzUseCase)(nil)

func NewFizzBuzzUseCase(repository repositories.FizzBuzzRepository) usecases.FizzBuzz {
	return &fizzBuzzUseCase{repository: repository}
}

func (f *fizzBuzzUseCase) Healthy() (string, error) {
	return f.repository.Healthy()
}

func (f *fizzBuzzUseCase) Ready() (string, error) {
	return f.repository.Ready()
}

func (f *fizzBuzzUseCase) GetFizzBuzz(arguments arguments.Arguments) (string, error) {

	result := fizzBuzz.FizzBuzz(arguments.Int1, arguments.Int2, arguments.Limit, arguments.Str1, arguments.Str2)
	if len(result) == 0 {
		return "", usecases.ErrWrongFizzBuzzArguments
	}
	fizzBuzzString := strings.Join(result, ", ")

	if err := f.repository.AddArgument(arguments); err != nil {
		switch err {
		case repositories.ErrSerializeArgument:
			return "", usecases.ErrSaveFizzBuzzArguments
		case repositories.ErrAddArgument:
			return fizzBuzzString, usecases.ErrSaveFizzBuzzArguments
		default: // should not occur
			log.Printf("Unexpected error: %s\n", err.Error())
		}
	}

	return fizzBuzzString, nil
}
