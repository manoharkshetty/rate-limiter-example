package inmemory_queue_map

import (
	"sync"
)

type QueueMap struct {
	lock   *sync.Mutex
	queues map[string]Queue
}

func NewQueueMap() *QueueMap {
	return &QueueMap{&sync.Mutex{}, map[string]Queue{}}
}

func (q *QueueMap) Add(user string, timeStamp int64, expiryTime int64) error {
	queue := q.queues[user]
	if queue == nil {
		q.lock.Lock()
		queue = NewQueue()
		q.queues[user] = queue
		q.lock.Unlock()
	}
	queue.Enqueue(timeStamp)
	go q.expire(user, expiryTime)
	return nil
}

func (q *QueueMap) GetLastRequestTimeInWindow(user string, timeStamp int64) (int64, error) {
	queue := q.queues[user]
	if queue == nil {
		return 0, nil
	}
	return queue.LastUpdatedSince(timeStamp), nil
}

func (q *QueueMap) GetRequestCountInWindow(user string, timeStamp int64) (int64, error) {
	queue := q.queues[user]
	if queue == nil {
		return 0, nil
	}
	return queue.ItemsRecordedAfter(timeStamp), nil
}

func (q *QueueMap) expire(user string, timeStamp int64) {
	defer DoPanicRecover()
	queue := q.queues[user]
	if queue == nil {
		return
	}
	queue.Dequeue(timeStamp)
	return
}
