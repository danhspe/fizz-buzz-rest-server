package cache

import "errors"

var ErrNoConnection = errors.New("no connection")
var ErrCloseConnection = errors.New("failed to close connection")
var ErrReadSortedSet = errors.New("failed to read sorted set")
var ErrWriteSortedSet = errors.New("failed to write sorted set")

type Cache interface {
	Connect() error
	Close() error

	SortedSetAdd(key string, member string, increment float64) (float64, error)
	SortedSetScore(key string, member string) (float64, error)
	SortedSetRangeWithScores(key string, start int64, stop int64) (map[string]int, error)
	SortedSetRangeByScoreWithScores(key string, min string, max string) (map[string]int, error)
}
