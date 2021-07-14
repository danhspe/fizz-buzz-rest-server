package fizzbuzz

import (
	"fmt"
	"log"

	"github.com/danhspe/fizz-buzz-rest-server/internal/fizzBuzz"
	"github.com/danhspe/fizz-buzz-rest-server/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/layers/usecases"
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

type fizzBuzzUseCase struct {
	repository repositories.FizzBuzzRepository
}

var _ usecases.FizzBuzz = (*fizzBuzzUseCase)(nil)

func NewFizzBuzzUseCase(repository repositories.FizzBuzzRepository) usecases.FizzBuzz {
	return &fizzBuzzUseCase{repository: repository}
}

func (f *fizzBuzzUseCase) GetFizzBuzz(arguments arguments.Arguments) string {

	result := fizzBuzz.FizzBuzz(arguments.Int1, arguments.Int2, arguments.Limit, arguments.Str1, arguments.Str2)

	err := f.repository.AddArgument(arguments)
	if err != nil {
		switch err { // TODO return error
		case repositories.ErrFailedToSerializeArgument:
			log.Fatalf("%s\n", err.Error())
		case repositories.ErrFailedToSaveArgument:
			log.Fatalf("%s\n", err.Error())
		default:
			log.Fatalf("%s\n", err.Error())
		}
	}

	return fmt.Sprintf("%+v", result)
}
