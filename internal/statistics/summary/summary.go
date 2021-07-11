package summary

import (
	"log"
	"strconv"

	"github.com/danhspe/fizz-buzz-rest-server/internal/cache"
	"github.com/danhspe/fizz-buzz-rest-server/internal/statistics"
)

type summary struct {
	cache         cache.Cache
	sortedSetName string
}

var _ statistics.Statistics = (*summary)(nil)

func NewSummary(cache cache.Cache, sortedSetName string) statistics.Statistics {
	return &summary{cache: cache, sortedSetName: sortedSetName}
}

func (s *summary) HighestScore() (int, error) {
	entries, err := s.cache.SortedSetRangeWithScores(s.sortedSetName, -1, -1)
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

func (s *summary) MostFrequentEntriesWithScores() (map[string]int, error) {
	highestCount, err := s.HighestScore()
	if err != nil {
		return nil, err
	}

	entries, err := s.cache.SortedSetRangeByScoreWithScores(s.sortedSetName, strconv.Itoa(highestCount), strconv.Itoa(highestCount))
	if err != nil {
		log.Printf("Failed to get most frequent entries from cache: %s\n", err.Error())
	}
	return entries, err
}
