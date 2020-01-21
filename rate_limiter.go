package rate_limiter

import "context"


//go:generate mockery -name=API -inpkg -case=underscore
type RateLimiter interface {
	// IsLimitReached will return true if a service user requests certain workflow/action with over high frequency within a time period.
	IsLimitReached(ctx context.Context, userID int64, tag string) bool

	// IsLimitReachedOnClient will return true if a Client requests certain workflow/action with over high frequency within a time period.
	// Not implementing this function. I have added this just to declare what rate limiter can do. This can be extended to client IP level rate limiting
	IsLimitReachedOnClient(ctx context.Context, client string, tag string) bool
}

//// receiving the config directly for simplicity.
//// TODO: receive interface to storage where we can fetch the configs from.
//func New(config Config) RateLimiter {
//
//}