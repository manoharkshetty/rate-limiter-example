package main

import (
	"github.com/manoharkshetty/rate-limiter-example/api"
	"github.com/manoharkshetty/rate-limiter-example/rate-limiter/config"
	"github.com/manoharkshetty/rate-limiter-example/rate-limiter/impl"
	"log"
	"net/http"
)

func main() {
	conf := map[string]*config.Config{
		"requester_a": {
			MaxReqAllowed:     5,
			TimeIntervalInSec: 10,
		},
		"requester_b": {
			MaxReqAllowed:     5,
			TimeIntervalInSec: 10,
		},
	}
	rateLimiter := ratelimiterimpl.NewFactory().Get(conf)
	home := api.New(rateLimiter)
	http.HandleFunc("/", home.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
