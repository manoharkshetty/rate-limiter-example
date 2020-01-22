package inmemory_queue_map

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewQueueMap(t *testing.T) {
	assert.NotNil(t, NewQueueMap())
}

func Test_Add(t *testing.T) {
	store := NewQueueMap()
	client := "client_1"
	expireTime := time.Now().Add(-10 * time.Minute).Unix()
	err := store.Add(client, time.Now().Unix(), expireTime)
	assert.Nil(t, err)

	count, err := store.GetRequestCountInWindow(client, expireTime)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), count)
}

func Test_GetRequestCountInWindow(t *testing.T) {
	store := NewQueueMap()
	client := "client_1"
	timeNow := time.Now().Unix()
	expireTime := time.Now().Add(-10 * time.Minute).Unix()
	err := store.Add(client, timeNow, expireTime)
	assert.Nil(t, err)

	// happy path
	count, err := store.GetRequestCountInWindow(client, expireTime)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), count)

	// expired
	count, err = store.GetRequestCountInWindow(client, timeNow)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), count)
}

func Test_GetLastRequestTimeInWindow(t *testing.T) {
	store := NewQueueMap()
	client := "client_1"
	timeNow := time.Now().Unix()
	expireTime := time.Now().Add(-10 * time.Minute).Unix()
	err := store.Add(client, timeNow, expireTime)
	assert.Nil(t, err)

	// happy path
	count, err := store.GetLastRequestTimeInWindow(client, expireTime)
	assert.Nil(t, err)
	assert.Equal(t, timeNow, count)

}
