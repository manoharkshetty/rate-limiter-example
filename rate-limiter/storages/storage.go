package storages

import queueMap "github.com/manoharkshetty/rate-limiter-example/rate-limiter/storages/inmemory-queue-map"

// Storage is an interface to store the requests date
//go:generate mockery -name=Storage -inpkg -case=underscore
type Storage interface {
	// Add stores the timestamp data for given user as key and expires the key on expireTime
	// For storage like redis, it should be easy task as setting expiry time on key
	// Other implementations should expire the keys older than expireTime
	Add(user string, timeStamp int64, expireTime int64) error

	// GetLastRequestTimeInWindow returns the last request time in a given time window of timeStamp and time.now
	GetLastRequestTimeInWindow(user string, timeStamp int64) (int64, error)

	// GetRequestCountInWindow returns the number of requests made in a given time window of timeStamp and time.now
	GetRequestCountInWindow(user string, timeStamp int64) (int64, error)
}

func GetStorage() Storage {
	return queueMap.NewQueueMap()
}
