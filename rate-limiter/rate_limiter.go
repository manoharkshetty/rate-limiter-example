package rate_limiter

import (
	"context"
)


//go:generate mockery -name=RateLimiter -inpkg -case=underscore
type RateLimiter interface {
	// IsLimitReached will return true if a requester requests certain action with over high frequency within a time period.
	IsLimitReached(ctx context.Context, requester string)(reached bool, tryAfter int64)

	// IsLimitReachedOnClient will return true if a Client requests certain workflow/action with over high frequency within a time period.
	// Not implementing this function. I have added this just to declare what rate limiter can do.
	// This can be extended to client geo level rate limiting
	IsLimitReachedOnClient(ctx context.Context, clientIP string, requester string)(reached bool, tryAfter int64)
}