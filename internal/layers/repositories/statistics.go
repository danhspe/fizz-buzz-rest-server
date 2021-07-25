package repositories

import (
	"errors"

	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
)

var ErrDeserializeArgument = errors.New("failed to deserialize argument")
var ErrGetHighestScore = errors.New("failed to get highest score")
var ErrGetMostFrequentEntriesWithScore = errors.New("failed to get most frequent entries with score")

type Statistics interface {
	// HighestScore returns the score of the most frequent entries; returns zero on error.
	HighestScore() (int, error)

	// MostFrequentEntriesWithScores returns the most frequent entries with their scores; returns nil on error.
	MostFrequentEntriesWithScores() (map[arguments.Arguments]int, error)
}
