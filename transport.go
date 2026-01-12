package ratelimit

import (
	"net/http"

	"go.uber.org/ratelimit"
)

// Option is a type alias for the ratelimit package's option type.
type Option = ratelimit.Option

// Common ratelimit options are exported for convenience.
var (
	Per          = ratelimit.Per
	WithClock    = ratelimit.WithClock
	WithSlack    = ratelimit.WithSlack
	WithoutSlack = ratelimit.WithoutSlack
)

// Transport is a http.RoundTripper that limits the rate of HTTP requests.
type Transport struct {
	Base    http.RoundTripper
	Limiter ratelimit.Limiter
}

// RoundTrip implements the http.RoundTripper interface.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Limiter != nil {
		t.Limiter.Take()
	}
	return t.Base.RoundTrip(req)
}

// New creates a new Transport with the given base transport and rate.
// If no base transport is provided, http.DefaultTransport is used.
// If the rate is zero or negative, the transport will not limit the rate.
func New(base http.RoundTripper, rate int, opts ...Option) *Transport {
	if base == nil {
		base = http.DefaultTransport
	}
	var limiter ratelimit.Limiter
	if rate > 0 {
		limiter = ratelimit.New(rate, opts...)
	}
	return &Transport{
		Base:    base,
		Limiter: limiter,
	}
}
