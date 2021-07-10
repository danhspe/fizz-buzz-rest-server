package cache

type Cache interface {
	Connect() bool
	Close() bool

	SortedSetAdd(key string, member string, increment float64) (float64, error)
	SortedSetScore(key string, member string) (float64, error)
	SortedSetRangeWithScores(key string, start int64, stop int64) (map[string]int, error)
	SortedSetRangeByScoreWithScores(key string, min string, max string) (map[string]int, error)
}
