package cache

import (
	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Connect() error {
	panic("implement me")
}

func (m *MockCache) Close() error {
	panic("implement me")
}

func (m *MockCache) SortedSetAdd(key string, member string, increment float64) (float64, error) {
	panic("implement me")
}

func (m *MockCache) SortedSetScore(key string, member string) (float64, error) {
	panic("implement me")
}

func (m *MockCache) SortedSetRangeWithScores(key string, start int64, stop int64) (map[string]int, error) {
	args := m.Called(key, start, stop)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(map[string]int), args.Error(1)
	}
}

func (m *MockCache) SortedSetRangeByScoreWithScores(key string, min string, max string) (map[string]int, error) {
	panic("implement me")
}

var _ Cache = (*MockCache)(nil)
