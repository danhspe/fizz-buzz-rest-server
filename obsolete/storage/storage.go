package storage

import (
	"sync"

	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
	"github.com/danhspe/fizz-buzz-rest-server/obsolete/argumentStore"
)

type Storage struct {
	mu            sync.Mutex
	argumentStore argumentStore.ArgumentStore
}

func New() *Storage {
	return &Storage{
		argumentStore: argumentStore.New(),
	}
}

func (s *Storage) Count(arguments arguments.Arguments) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.argumentStore[arguments]
}

func (s *Storage) Increment(arguments arguments.Arguments) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentCount := s.argumentStore[arguments]
	newCount := currentCount + 1
	s.argumentStore[arguments] = newCount

	return newCount
}

func (s *Storage) HighestCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	highestCount := 0
	for _, count := range s.argumentStore {
		if count > highestCount {
			highestCount = count
		}
	}
	return highestCount
}

func (s *Storage) MostFrequentEntries() argumentStore.ArgumentStore {
	highestCount := s.HighestCount()

	s.mu.Lock()
	defer s.mu.Unlock()

	mostFrequentEntries := argumentStore.New()
	for key, value := range s.argumentStore {
		if value == highestCount {
			mostFrequentEntries[key] = value
		}
	}
	return mostFrequentEntries
}

var _ Counter = (*Storage)(nil)
