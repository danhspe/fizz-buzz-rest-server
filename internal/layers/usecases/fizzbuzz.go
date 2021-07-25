package usecases

import (
	"errors"

	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

var ErrSaveFizzBuzzArguments = errors.New("failed to save fizz buzz arguments")

type FizzBuzz interface {
	GetFizzBuzz(arguments arguments.Arguments) (string, error)
}
