package inmemory_queue_map

import (
	"log"
	"sync"
	"time"
)

type queue struct {
	lock *sync.Mutex
	Values []int64
}

type Queue interface {
	Enqueue(x int64)
	Dequeue(timeStamp int64)
	ItemsRecordedAfter(timeStamp int64) int64
	LastUpdatedSince(timeStamp int64) int64
}

func NewQueue() Queue {
	return &queue{&sync.Mutex{}, make([]int64, 0)}
}

func (q *queue) Enqueue(x int64) {
	q.lock.Lock()
	q.Values = append(q.Values, x)
	q.lock.Unlock()
	return
}

func (q *queue) Dequeue(timeStamp int64) {
	if len(q.Values) == 0 {
		return
	}
	index := q.searchNearestItem(timeStamp)
	if index > 0 {
		log.Print("dequeuing elements, count: ", index)
		q.lock.Lock()
		q.Values = q.Values[index:]
		q.lock.Unlock()
	}
}

func (q *queue) ItemsRecordedAfter(timeStamp int64) int64 {
	index := q.searchNearestItem(timeStamp)
	return int64(len(q.Values[(index):]))
}

func (q *queue) LastUpdatedSince(timeStamp int64) int64 {
	index := q.searchNearestItem(timeStamp)
	if index >= int64(len(q.Values)) {
		return time.Now().Unix()
	}
	return q.Values[index]
}

func (q *queue) searchNearestItem(timeStamp int64) int64  {
	index := 0
	for _, val := range q.Values {
		if val <= timeStamp {
			index += 1
		}
		continue
	}
	return int64(index)
}

