package api

import (
	limiter "github.com/manoharkshetty/rate-limiter-example/rate-limiter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_New(t *testing.T) {
	mockLimiter := &limiter.MockRateLimiter{}
	assert.NotNil(t, New(mockLimiter))
}

func Test_Handle(t *testing.T) {
	mockLimiter := &limiter.MockRateLimiter{}
	mockLimiter.On("IsLimitReached", mock.Anything, mock.Anything).Return(false, int64(0))
	handler := New(mockLimiter)
	handler.Handle(httptest.NewRecorder(), &http.Request{URL: &url.URL{Path: "/requester"}})
}