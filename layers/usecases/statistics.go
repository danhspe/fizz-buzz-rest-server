package usecases

import (
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

type Statistics interface {
	GetStatistics() (highestScore int, mostFrequentArguments []arguments.Arguments)
}
