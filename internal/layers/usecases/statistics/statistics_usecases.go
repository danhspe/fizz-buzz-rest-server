package statistics

import (
	"log"

	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/usecases"
	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
)

type statisticsUseCases struct {
	statistics repositories.Statistics
}

var _ usecases.Statistics = (*statisticsUseCases)(nil)

func NewStatisticsUseCases(statistics repositories.Statistics) usecases.Statistics {
	return &statisticsUseCases{statistics: statistics}
}

func (s *statisticsUseCases) GetStatistics() (highestScore int, mostFrequentArguments []arguments.Arguments, err error) {

	highestScore, err = s.statistics.HighestScore()
	if err != nil {
		return 0, nil, usecases.ErrGetStatistics
	}

	argumentsWithScores, err := s.statistics.MostFrequentEntriesWithScores()
	if err != nil {
		return 0, nil, usecases.ErrGetStatistics
	}

	mostFrequentArguments = make([]arguments.Arguments, 0)
	for argument, score := range argumentsWithScores {
		if score != highestScore {
			log.Printf("expected highest score %d to match arguments score: %d\n", highestScore, score)
		} else {
			mostFrequentArguments = append(mostFrequentArguments, argument)
		}
	}

	return highestScore, mostFrequentArguments, err
}
