package inmemory_queue_map

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewQueue(t *testing.T)  {
	assert.NotNil(t, NewQueue())
}

func Test_Enqueue(t *testing.T)  {
	queue := NewQueue()
	queue.Enqueue(time.Now().Unix())
	assert.Equal(t, int64(1), queue.ItemsRecordedAfter(time.Now().Add(-1* time.Minute).Unix()))
}

func Test_Dequeue(t *testing.T)  {
	queue := NewQueue()
	expire := time.Now().Add(-1* time.Minute).Unix()
	queue.Enqueue(time.Now().Unix())
	assert.Equal(t, int64(1), queue.ItemsRecordedAfter(expire))

	queue.Dequeue(time.Now().Unix())
	assert.Equal(t, int64(0), queue.ItemsRecordedAfter(expire))
}

func Test_ItemsRecordedAfter(t *testing.T)  {
	queue := NewQueue()
	queue.Enqueue(time.Now().Unix())
	assert.Equal(t, int64(1), queue.ItemsRecordedAfter(time.Now().Add(-1* time.Minute).Unix()))
}

func Test_LastUpdatedSince(t *testing.T)  {
	queue := NewQueue()
	timeNow := time.Now().Unix()
	queue.Enqueue(timeNow)
	assert.Equal(t, timeNow, queue.LastUpdatedSince(time.Now().Add(-1* time.Minute).Unix()))
}