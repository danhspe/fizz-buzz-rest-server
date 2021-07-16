package repositories

import (
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

type Statistics interface {
	// HighestScore returns the score of the most frequent entries; returns zero on error.
	HighestScore() (int, error)

	// MostFrequentEntriesWithScores returns the most frequent entries with their scores; returns nil on error.
	MostFrequentEntriesWithScores() (map[arguments.Arguments]int, error)
}
