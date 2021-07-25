package fizzbuzz

import (
	"log"

	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
	"github.com/danhspe/fizz-buzz-rest-server/internal/storage/cache"
)

const defaultSortedSetName = "fizzbuzz"

type fizzBuzzRepository struct {
	cache cache.Cache
}

func NewFizzBuzzRepository(cache cache.Cache) repositories.FizzBuzzRepository {
	return &fizzBuzzRepository{cache: cache}
}

func (f *fizzBuzzRepository) AddArgument(argument arguments.Arguments) error {

	entry, err := argument.AsJson()
	if err != nil {
		return repositories.ErrSerializeArgument
	}

	_, err = f.cache.SortedSetAdd(defaultSortedSetName, entry, 1)
	if err != nil && err == cache.ErrWriteSortedSet {
		log.Printf("AddArgument error: %s\n", err.Error())
		return repositories.ErrAddArgument
	}

	return nil
}
