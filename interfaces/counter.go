package interfaces

import (
	"github.com/danhspe/fizz-buzz-rest-server/models/argumentStore"
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
)

type Counter interface {

	// Count returns the counter of the arguments.
	Count(arguments arguments.Arguments) int

	// Increment increments the counter of the arguments and returns the value.
	// The arguments will be added if they do not exist yet.
	Increment(arguments arguments.Arguments) int

	// HighestCount returns the counter of the most frequent entry.
	HighestCount() int

	// MostFrequentEntries returns the most frequent entries.
	MostFrequentEntries() argumentStore.ArgumentStore
}
