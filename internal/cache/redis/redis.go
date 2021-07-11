package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/danhspe/fizz-buzz-rest-server/internal/cache"
	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	client *redis.Client
	ctx    context.Context
}

var _ cache.Cache = (*redisCache)(nil)

func NewRedisCache(address string) cache.Cache {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	return &redisCache{
		client: client,
		ctx:    ctx,
	}
}

func (c *redisCache) Connect() bool {
	statusCmd := c.client.Ping(c.ctx)
	if _, err := statusCmd.Result(); err != nil {
		log.Printf(fmt.Sprintf("Failed to connect/ping Redis at %s: %s\n", c.client.Options().Addr, err.Error()))
		return false
	}
	return true
}

func (c *redisCache) Close() bool {
	err := c.client.Close()
	if err != nil {
		log.Printf(fmt.Sprintf("Failed to close Redis client for %s: %s\n", c.client.Options().Addr, err.Error()))
		return false
	}
	return true
}

// SortedSetAdd adds the member and increments it's score.
// If the member exists already, the score is updated and the element reinserted at the right position to ensure the correct ordering.
func (c *redisCache) SortedSetAdd(key string, member string, increment float64) (float64, error) {
	floatCmd := c.client.ZAddArgsIncr(c.ctx, key, redis.ZAddArgs{
		Members: []redis.Z{{Score: increment, Member: member}},
	})

	result, err := floatCmd.Result()
	if err != nil {
		log.Printf("ZAddArgsIncr failed : %s", err.Error())
	}
	log.Printf("ZAddArgsIncr result: %+v", result)
	return result, nil
}

// SortedSetScore returns the score of the member.
func (c *redisCache) SortedSetScore(key string, member string) (float64, error) {
	floatCmd := c.client.ZScore(c.ctx, key, member)
	result, err := floatCmd.Result()
	if err != nil {
		log.Printf("ZScore failed: %s", err.Error())
	}
	log.Printf("ZScore result: %+v", result)
	return result, err
}

// SortedSetRangeWithScores returns the members with indices from start to stop sorted by score; returns nil on error.
// Returns only the last member in the sorted set when start and stop is set to -1.
func (c *redisCache) SortedSetRangeWithScores(key string, start int64, stop int64) (map[string]int, error) {
	zSliceCmd := c.client.ZRangeWithScores(c.ctx, key, start, stop)
	result, err := zSliceCmd.Result()
	if err != nil {
		log.Printf("ZRangeWithScores failed: %s", err.Error())
		return nil, err
	}
	log.Printf("ZRangeWithScores result: %+v", result)

	entries := make(map[string]int)
	for i, v := range result {
		score := v.Score
		member := v.Member
		log.Printf("index: %d score: %f member: %+v", i, score, member)
		if str, ok := member.(string); ok {
			entries[str] = int(score)
		}
	}
	return entries, err
}

// SortedSetRangeByScoreWithScores returns the members with scores between min and max; returns nil on error.
func (c *redisCache) SortedSetRangeByScoreWithScores(key string, min string, max string) (map[string]int, error) {
	stringSliceCmd := c.client.ZRangeByScoreWithScores(c.ctx, key, &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: 0,
		Count:  0,
	})
	result, err := stringSliceCmd.Result()
	if err != nil {
		log.Printf("ZRangeByScoreWithScores failed: %s", err.Error())
		return nil, err
	}
	log.Printf("ZRangeByScoreWithScores result: %+v", result)

	entries := make(map[string]int)
	for k, v := range result {
		score := v.Score
		member := v.Member
		log.Printf("index: %d score: %f member: %+v", k, score, member)
		if str, ok := member.(string); ok {
			entries[str] = int(score)
		}
	}
	return entries, err
}
