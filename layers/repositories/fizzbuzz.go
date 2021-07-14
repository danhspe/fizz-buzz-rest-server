package repositories

import (
	"errors"

	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

var ErrFailedToSerializeArgument = errors.New("failed to serialize argument")
var ErrFailedToSaveArgument = errors.New("failed to save argument")

type FizzBuzzRepository interface {
	AddArgument(argument arguments.Arguments) error
}
