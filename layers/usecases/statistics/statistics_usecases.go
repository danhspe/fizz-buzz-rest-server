package statistics

import (
	"log"

	"github.com/danhspe/fizz-buzz-rest-server/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/layers/usecases"
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

type statisticsUseCases struct {
	statistics repositories.Statistics
}

var _ usecases.Statistics = (*statisticsUseCases)(nil)

func NewStatisticsUseCases(statistics repositories.Statistics) usecases.Statistics {
	return &statisticsUseCases{statistics: statistics}
}

func (s *statisticsUseCases) GetStatistics() (highestScore int, mostFrequentArguments []arguments.Arguments) {

	highestScore, _ = s.statistics.HighestScore()

	argumentsWithScores, err := s.statistics.MostFrequentEntriesWithScores()
	if err != nil {
		log.Printf("GetStatistics error: %s\n", err.Error())
	}

	argumentsList := make([]arguments.Arguments, 0)
	for argument, score := range argumentsWithScores {
		if score != highestScore {
			log.Printf("Highest score %d does not match arguments score: %d\n", highestScore, score)
		} else {
			argumentsList = append(argumentsList, argument)
		}
	}

	return highestScore, argumentsList
}
