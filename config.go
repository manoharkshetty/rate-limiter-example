package rate_limiter

type Config struct {
	TimeIntervalInSec int64
	MaxReqAllowed int64
}
