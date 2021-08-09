package usecases

import (
	"errors"

	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
	"github.com/danhspe/fizz-buzz-rest-server/internal/models/health"
)

var ErrGetStatistics = errors.New("failed to get fizz buzz statistics")

type Statistics interface {
	health.Health
	GetStatistics() (highestScore int, mostFrequentArguments []arguments.Arguments, err error)
}
