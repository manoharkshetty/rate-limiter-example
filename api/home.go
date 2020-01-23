package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	limiter "github.com/manoharkshetty/rate-limiter-example/rate-limiter"
)

type HomeHandler struct {
	rateLimiter limiter.RateLimiter
}

func New(rateLimiter limiter.RateLimiter) Handler {
	return HomeHandler{
		rateLimiter: rateLimiter,
	}
}

func (handler HomeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Path[1:]
	limitReached, tryAfter := handler.rateLimiter.IsLimitReached(context.Background(), user)
	if limitReached {
		outputJSON(w, http.StatusTooManyRequests, fmt.Sprintf("Try after %d second(s)", tryAfter))
		return
	}
	outputJSON(w, http.StatusOK, "success")
}

func outputJSON(w http.ResponseWriter, respCode int, payload interface{}) {
	output, err := json.Marshal(payload)
	if err != nil {
		respCode = http.StatusInternalServerError
	}
	w.WriteHeader(respCode)
	_, _ = w.Write(output)
}
