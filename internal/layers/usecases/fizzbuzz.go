package usecases

import (
	"errors"

	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
	"github.com/danhspe/fizz-buzz-rest-server/internal/models/health"
)

var (
	ErrWrongFizzBuzzArguments = errors.New("wrong fizz buzz arguments")
	ErrSaveFizzBuzzArguments  = errors.New("failed to save fizz buzz arguments")
)

type FizzBuzz interface {
	health.Health
	GetFizzBuzz(arguments arguments.Arguments) (string, error)
}
