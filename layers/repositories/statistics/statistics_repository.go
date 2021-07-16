package statistics

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/danhspe/fizz-buzz-rest-server/internal/cache"
	"github.com/danhspe/fizz-buzz-rest-server/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

const defaultSortedSetName = "fizzbuzz"

type statisticsRepository struct {
	cache cache.Cache
}

var _ repositories.Statistics = (*statisticsRepository)(nil)

func NewStatisticsRepository(cache cache.Cache) repositories.Statistics {
	return &statisticsRepository{cache: cache}
}

func (s *statisticsRepository) HighestScore() (int, error) {
	entries, err := s.cache.SortedSetRangeWithScores(defaultSortedSetName, -1, -1)
	if err != nil {
		log.Printf("Failed to get highest scores from cache: %s\n", err.Error())
		return 0, err
	}

	var highestScore int
	for member, score := range entries {
		log.Printf("member: %s score: %+d\n", member, score)
		highestScore = score
	}
	return highestScore, nil
}

func (s *statisticsRepository) MostFrequentEntriesWithScores() (map[arguments.Arguments]int, error) {

	highestCount, err := s.HighestScore()
	if err != nil {
		return nil, err
	}

	entries, err := s.cache.SortedSetRangeByScoreWithScores(defaultSortedSetName, strconv.Itoa(highestCount), strconv.Itoa(highestCount))
	if err != nil {
		log.Printf("Failed to get most frequent entries from cache: %s\n", err.Error())
	}

	argumentsWithScores := make(map[arguments.Arguments]int)
	for entry, score := range entries {
		var arg arguments.Arguments
		if err := json.Unmarshal([]byte(entry), &arg); err != nil {
			log.Printf("Failed to unmarshal arguments: %s\n", err.Error())
			return argumentsWithScores, err
		}
		argumentsWithScores[arg] = score
	}

	return argumentsWithScores, nil
}