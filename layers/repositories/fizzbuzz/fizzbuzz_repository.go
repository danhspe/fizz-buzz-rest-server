package fizzbuzz

import (
	"log"

	"github.com/danhspe/fizz-buzz-rest-server/internal/cache"
	"github.com/danhspe/fizz-buzz-rest-server/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/models/arguments"
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
		log.Printf("AddArgument error: %s\n", err.Error())
		return repositories.ErrFailedToSerializeArgument
	}

	if _, err = f.cache.SortedSetAdd(defaultSortedSetName, entry, 1); err != nil {
		log.Printf("AddArgument error: %s\n", err.Error())
		return repositories.ErrFailedToSaveArgument
	}

	return nil
}
