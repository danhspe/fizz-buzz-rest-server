package statistics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories"
	"github.com/danhspe/fizz-buzz-rest-server/internal/layers/repositories/statistics"
	"github.com/danhspe/fizz-buzz-rest-server/internal/storage/cache"
)

const defaultSortedSetName = "fizzbuzz"

func TestStatisticsRepository_HighestScore_WhenRetrievingHighestScoreFromCacheFails_ThenReturnError(t *testing.T) {
	var mockResult map[string]int = nil
	mockError := cache.ErrReadSortedSet
	mockCache := cache.MockCache{}
	mockCache.On("SortedSetRangeWithScores", defaultSortedSetName, int64(-1), int64(-1)).Return(mockResult, mockError)

	statisticsRepository := statistics.NewStatisticsRepository(&mockCache)
	expectedScore := 0
	expectedError := repositories.ErrGetHighestScore

	score, err := statisticsRepository.HighestScore()

	mockCache.AssertExpectations(t)
	assert.Equal(t, expectedScore, score, "Expected score %d - got %d", expectedScore, score)
	assert.Equal(t, expectedError, err, "Expected error %s - got %s", expectedError, err)
}

func TestStatisticsRepository_HighestScore_WhenRetrievingHighestScoreFromCacheSucceeds_ThenReturnScore(t *testing.T) {
	mockResult := map[string]int{"mockEntry": 1}
	mockCache := cache.MockCache{}
	mockCache.On("SortedSetRangeWithScores", defaultSortedSetName, int64(-1), int64(-1)).Return(mockResult, nil)

	statisticsRepository := statistics.NewStatisticsRepository(&mockCache)
	expectedScore := 1

	score, err := statisticsRepository.HighestScore()

	mockCache.AssertExpectations(t)
	assert.Equal(t, expectedScore, score, "Expected score %d - got %d", expectedScore, score)
	assert.Nil(t, err, "Expected nil error - got %s", err)
}
