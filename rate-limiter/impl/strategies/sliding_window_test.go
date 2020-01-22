package ratelimiterimpl

import (
	"context"
	"github.com/manoharkshetty/rate-limiter-example/rate-limiter/config"
	"github.com/manoharkshetty/rate-limiter-example/rate-limiter/storages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func Test_NewSlidingWindowImpl(t *testing.T)  {
	storage := &storages.MockStorage{}
	impl := NewSlidingWindowImpl(map[string]*config.Config{}, storage)
	assert.NotNil(t, impl)
}

func Test_IsLimitReached(t *testing.T)  {
	scenarios := []struct{
		expectedLimitReached bool
		expectedTryAfter int64
		requestCount int64
		lastRequestAt int64
		description string
		config map[string]*config.Config
	}{
		{
			expectedLimitReached: false,
			description: "limit has not reached",
			config: map[string]*config.Config{
				"client_1": {
					MaxReqAllowed: 5,
					TimeIntervalInSec: 10,
				},
			},
			requestCount: 1,
		},
		{
			expectedLimitReached: true,
			expectedTryAfter: int64(10),
			config: map[string]*config.Config{
				"client_1": {
					MaxReqAllowed: 2,
					TimeIntervalInSec: 10,
				},
			},
			requestCount: 5,
			lastRequestAt: time.Now().Unix(),
			description: "limit has reached",
		},
		{
			expectedLimitReached: false,
			description: "invalid config",
		},
	}
	for _,  sc := range scenarios {
		storage := storages.MockStorage{}
		storage.On("GetRequestCountInWindow", mock.Anything, mock.Anything).Return(sc.requestCount, nil)
		storage.On("GetLastRequestTimeInWindow", mock.Anything, mock.Anything).Return(sc.lastRequestAt, nil)
		storage.On("Add", mock.Anything, mock.Anything, mock.Anything).Return( nil)
		impl := slidingWindowImpl{
			configMap: sc.config,
			storage: &storage,
		}
		limitReached, tryAfter := impl.IsLimitReached(context.Background(), "client_1")
		assert.Equal(t, sc.expectedLimitReached, limitReached, sc.description)
		assert.Equal(t, sc.expectedTryAfter, tryAfter, sc.description)
	}

}