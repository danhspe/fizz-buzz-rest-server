package repositories

import (
	"errors"

	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
)

var ErrSerializeArgument = errors.New("failed to serialize argument")
var ErrAddArgument = errors.New("failed to add argument")

type FizzBuzzRepository interface {
	AddArgument(argument arguments.Arguments) error
}
