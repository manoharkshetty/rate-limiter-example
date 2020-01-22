package ratelimiterimpl

import (
	rate_limiter "github.com/manoharkshetty/rate-limiter-example/rate-limiter"
	"github.com/manoharkshetty/rate-limiter-example/rate-limiter/config"
	impl "github.com/manoharkshetty/rate-limiter-example/rate-limiter/impl/strategies"
	"github.com/manoharkshetty/rate-limiter-example/rate-limiter/storages"
)

type factory struct {
	storage  storages.Storage
}

//go:generate mockery -name=Factory -inpkg -case=underscore
// Factory can return any implementation based on the parameters or configuration(feature flags)
// Currently only implemented SlidingWindowImpl, Can be extended for leaky bucket, token etc
type Factory interface {
	Get(configMap  map[string]*config.Config) rate_limiter.RateLimiter
}

func NewFactory() Factory {
	return &factory{
		storage: storages.GetStorage(),
	}
}

func (f *factory) Get(configMap  map[string]*config.Config) rate_limiter.RateLimiter  {
	return impl.NewSlidingWindowImpl(configMap, f.storage)
}
