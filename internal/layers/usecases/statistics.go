package usecases

import (
	"errors"

	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
)

var ErrGetStatistics = errors.New("failed to get fizz buzz statistics")

type Statistics interface {
	GetStatistics() (highestScore int, mostFrequentArguments []arguments.Arguments, err error)
}
