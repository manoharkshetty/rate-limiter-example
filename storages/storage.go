package storages

import "time"

type Storage interface {
	Set(user string, time time.Time) error
	Get(user string, time time.Time) (int64, error)
	Expire(user string, time time.Time) error
}
