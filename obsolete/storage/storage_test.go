package storage_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
	"github.com/danhspe/fizz-buzz-rest-server/obsolete/storage"
)

func TestNew(t *testing.T) {
	if storage.New() == nil {
		t.Fatal("Failed to create statistics")
	}
}

func TestStorage_Count_WhenEntryDoesNotExist_ThenCountReturnsZero(t *testing.T) {
	args := arguments.New(0, 0, 0, "", "")
	db := storage.New()

	count := db.Count(args)

	if count != 0 {
		t.Errorf("Expected 0, got %d", count)
	}
}

func TestStorage_Count_WhenNewEntryIsIncremented_ThenCountIsOne(t *testing.T) {
	args := arguments.New(0, 0, 0, "", "")
	db := storage.New()
	newCount := db.Increment(args)
	t.Log(fmt.Sprintf("New count: %d", newCount))

	count := db.Count(args)

	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestStorage_Increment(t *testing.T) {
	args := arguments.New(0, 0, 0, "", "")
	db := storage.New()
	expectedCount := 3

	_ = db.Increment(args)
	_ = db.Increment(args)
	newCount := db.Increment(args)

	if newCount != expectedCount {
		t.Errorf("Expected count %d, got %d", expectedCount, newCount)
	}

	if newCount != db.Count(args) {
		t.Errorf("Expected count %d, got %d", expectedCount, newCount)
	}
}

func TestStorage_HighestCount(t *testing.T) {
	args1 := arguments.New(0, 0, 0, "", "")
	args2 := arguments.New(1, 1, 1, "", "")

	db := storage.New()
	expectedHighestCount := 2

	_ = db.Increment(args1)
	_ = db.Increment(args1)
	_ = db.Increment(args2)

	highestCount := db.HighestCount()
	if highestCount != expectedHighestCount {
		t.Errorf("Expected highest count %d, got %d", expectedHighestCount, highestCount)
	}
}

func TestStorage_MostFrequentEntries_InitiallyThereAreNoEntries(t *testing.T) {
	db := storage.New()

	mostFrequentEntries := db.MostFrequentEntries()
	numberOfMostFrequentEntries := len(mostFrequentEntries)

	if numberOfMostFrequentEntries != 0 {
		t.Errorf("Expected zero most frequent entries, got %d", numberOfMostFrequentEntries)
	}
	t.Log(fmt.Sprintf("%+v", mostFrequentEntries))
}

func TestStorage_MostFrequentEntries_OneMostFrequentEntry(t *testing.T) {
	args1 := arguments.New(0, 0, 0, "", "")
	args2 := arguments.New(1, 1, 1, "", "")

	db := storage.New()
	expectedNumberOfMostFrequentEntries := 1

	_ = db.Increment(args1)
	_ = db.Increment(args1)
	_ = db.Increment(args2)

	mostFrequentEntries := db.MostFrequentEntries()
	numberOfMostFrequentEntries := len(mostFrequentEntries)

	if numberOfMostFrequentEntries != expectedNumberOfMostFrequentEntries {
		t.Errorf("Expected most frequent entries %d, got %d", expectedNumberOfMostFrequentEntries, numberOfMostFrequentEntries)
	}
	t.Log(fmt.Sprintf("%+v", mostFrequentEntries))
}

func TestStorage_MostFrequentEntries_MultipleMostFrequentEntry(t *testing.T) {
	args1 := arguments.New(0, 0, 0, "", "")
	args2 := arguments.New(1, 1, 1, "", "")

	db := storage.New()
	expectedNumberOfMostFrequentEntries := 2

	_ = db.Increment(args1)
	_ = db.Increment(args2)

	mostFrequentEntries := db.MostFrequentEntries()
	numberOfMostFrequentEntries := len(mostFrequentEntries)

	if numberOfMostFrequentEntries != expectedNumberOfMostFrequentEntries {
		t.Errorf("Expected most frequent entries %d, got %d", expectedNumberOfMostFrequentEntries, numberOfMostFrequentEntries)
	}
	t.Log(fmt.Sprintf("%+v", mostFrequentEntries))
}

func TestStorage_Counter(t *testing.T) {
	var uniqueArgs = []arguments.Arguments{
		arguments.New(0, 0, 0, "", ""),
		arguments.New(3, 5, 100, "fizz", "buzz"),
	}

	db := storage.New()
	repetitions := 5

	for _, tt := range uniqueArgs {
		for reps := 0; reps < repetitions; reps = reps + 1 {
			oldValue := db.Count(tt)
			newValue := db.Increment(tt)
			t.Log(fmt.Sprintf("Increment count for %v from %d to %d", tt, oldValue, newValue))
		}
	}
	t.Log(fmt.Sprintf("DB: %v", db))

	assert.Equal(t, len(uniqueArgs), len(db.MostFrequentEntries()))
	assert.Equal(t, repetitions, db.HighestCount())
	for i := range uniqueArgs {
		assert.Equal(t, repetitions, db.Count(uniqueArgs[i]))
	}
}
