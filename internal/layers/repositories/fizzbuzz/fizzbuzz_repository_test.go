package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories/fizzbuzz"
	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
	"github.com/danhspe/fizz-buzz-rest-server/internal/storage/cache"
)

const defaultSortedSetName = "fizzbuzz"

var (
	testArgument = arguments.Arguments{
		Int1:  0,
		Int2:  0,
		Limit: 0,
		Str1:  "",
		Str2:  "",
	}
	testArgumentAsJson, _ = testArgument.AsJson()
)

func TestFizzBuzzRepository_AddArgument_WhenAddingArgumentToCacheFails_ThenReturnError(t *testing.T) {
	mockIncrement := 1.0
	mockError := cache.ErrWriteSortedSet
	mockCache := cache.MockCache{}
	mockCache.On("SortedSetAdd", defaultSortedSetName, testArgumentAsJson, mockIncrement).Return(mockIncrement, mockError)

	fizzBuzzRepository := fizzbuzz.NewFizzBuzzRepository(&mockCache)
	expectedError := repositories.ErrAddArgument

	err := fizzBuzzRepository.AddArgument(testArgument)

	mockCache.AssertExpectations(t)
	assert.Equal(t, expectedError, err, "Expected error %s - got %s", expectedError, err)
}

func TestFizzBuzzRepository_AddArgument_WhenAddingArgumentToCacheSucceeds_ThenErrorShouldBeNil(t *testing.T) {
	var scoreIncrement float64 = 1
	mockCache := cache.MockCache{}
	mockCache.On("SortedSetAdd", defaultSortedSetName, testArgumentAsJson, scoreIncrement).Return(scoreIncrement, nil)

	fizzBuzzRepository := fizzbuzz.NewFizzBuzzRepository(&mockCache)

	err := fizzBuzzRepository.AddArgument(testArgument)

	mockCache.AssertExpectations(t)
	assert.Nil(t, err, "Expected nil error - got %s", err)
}
