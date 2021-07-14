package usecases

import "github.com/danhspe/fizz-buzz-rest-server/models/arguments"

type FizzBuzz interface {
	GetFizzBuzz(arguments arguments.Arguments) string
}
