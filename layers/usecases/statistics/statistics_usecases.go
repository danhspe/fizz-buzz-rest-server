package statistics

import (
	"github.com/danhspe/fizz-buzz-rest-server/internal/statistics"
	"github.com/danhspe/fizz-buzz-rest-server/layers/usecases"
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

type statisticsUseCases struct {
	statistics statistics.Statistics
}

var _ usecases.Statistics = (*statisticsUseCases)(nil)

func NewStatisticsUseCases(statistics statistics.Statistics) usecases.Statistics {
	return &statisticsUseCases{statistics: statistics}
}

func (s *statisticsUseCases) GetStatistics() (highestScore int, mostFrequentArguments []arguments.Arguments) {

	highestScore, _ = s.statistics.HighestScore()

	_, _ = s.statistics.MostFrequentEntriesWithScores()

	return highestScore, []arguments.Arguments{{
		Int1:  3,
		Int2:  5,
		Limit: 100,
		Str1:  "Fizz",
		Str2:  "Buzz",
	}}
}
