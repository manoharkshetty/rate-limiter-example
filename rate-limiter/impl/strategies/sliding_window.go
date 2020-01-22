package ratelimiterimpl

import (
	"context"
	"log"
	"time"

	"github.com/manoharkshetty/rate-limiter-example/rate-limiter/config"
	"github.com/manoharkshetty/rate-limiter-example/rate-limiter/storages"
)

const logTag = "slidingWindowImpl"

// Sliding window strategy keeps track of requests made in a given time window to make decisions about throttling
// if requests made in given time frame are less than maxRequestAllowed then it allows requests
// if not the requests are throttled
type slidingWindowImpl struct {
	configMap map[string]*config.Config
	storage   storages.Storage
}

func NewSlidingWindowImpl(configMap map[string]*config.Config, storage storages.Storage) *slidingWindowImpl {
	return &slidingWindowImpl{
		configMap: configMap,
		storage:   storage,
	}
}

// IsLimitReached will return true if a requester requests certain action with over high frequency within a time period.
func (i *slidingWindowImpl) IsLimitReached(ctx context.Context, requester string) (bool, int64) {
	return i.isLimitReached(requester)
}

// Not implemented
func (i *slidingWindowImpl) IsLimitReachedOnClient(ctx context.Context, client string, requester string) (reached bool, tryAfter int64) {
	log.Println("called unimplemented function", client, requester)
	return false, 0
}

func (i *slidingWindowImpl) isLimitReached(identifier string) (bool, int64) {
	var maxReqAllowed int64
	var timeIntervalInSec int64

	rateLimitConf := i.configMap[identifier]
	if rateLimitConf != nil {
		maxReqAllowed = rateLimitConf.MaxReqAllowed
		timeIntervalInSec = rateLimitConf.TimeIntervalInSec
	}
	// If any config var is given default value 0 (which make no sense) we don't apply rate-limiting to this workflow
	if timeIntervalInSec == 0 || maxReqAllowed == 0 {
		return false, 0
	}

	currentTime := time.Now().Unix()
	startTime := currentTime - timeIntervalInSec
	reqCount, err := i.storage.GetRequestCountInWindow(identifier, startTime)
	if err != nil {
		log.Println(logTag, "[isLimitReached][error:%s] Failed to check request count in storage", err)
		return false, 0
	}

	// has not reached the limit yet
	if reqCount < maxReqAllowed {
		go i.updateRequestCount(identifier, currentTime, startTime)
		return false, 0
	}

	// get the lastRequested time within the window
	lastRequestAt, err := i.storage.GetLastRequestTimeInWindow(identifier, startTime)
	if err != nil {
		log.Println(logTag, "[isLimitReached][error:%s] Failed to get last updated time, setting it to default", err)
		lastRequestAt = currentTime
	}

	// user should try after the time where last request expires
	tryAfter := lastRequestAt - startTime
	return true, tryAfter
}

func (i *slidingWindowImpl) updateRequestCount(identifier string, currentTime int64, startTime int64) {
	if err := i.storage.Add(identifier, currentTime, startTime); err != nil {
		log.Println("error Adding new key", identifier)
	}
}

