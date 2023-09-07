package http

import (
	net_http "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/metrics"
)

// Mux defines the standard Multiplexer for http Request
type Mux interface {
	// ServeHTTP
	net_http.Handler

	// Default Handler Method is all that is needed
	Handler(method, url string, fn net_http.Handler)
}

type MuxOption func(*muxer)

type muxer struct {
	*chi.Mux
}

func (mx *muxer) Handler(method, url string, fn net_http.Handler) {
	mx.Method(method, url, fn)
}

func NewDefaultMux(opts ...MuxOption) Mux {
	mx := &muxer{chi.NewMux()}

	for _, o := range opts {
		o(mx)
	}

	return mx
}

// Metricser is wrapper for supported metrics agents
type Metricser interface {
	Counter(prefix, name string) metrics.Counter

	Histogram(prefix, name string) metrics.Histogram

	Handler() net_http.Handler
}

// ContextKey is key for context
type ContextKey int

// ContextKeys
const (
	ContextKeyRequestMethod ContextKey = iota
	ContextKeyRequestURI
	ContextKeyRequestPath
	ContextKeyRequestProto
	ContextKeyRequestHost
	ContextKeyRequestRemoteAddr
	ContextKeyRequestXForwardedFor
	ContextKeyRequestXForwardedProto
	ContextKeyRequestAuthorization
	ContextKeyRequestReferer
	ContextKeyRequestUserAgent
	ContextKeyRequestXRequestID
	ContextKeyRequestAccept
	ContextKeyResponseHeaders
	ContextKeyResponseSize
)

// Headers
const (
	HeaderAllowHeaders = "Access-Control-Allow-Headers"
	HeaderAllowMethods = "Access-Control-Allow-Methods"
	HeaderAllowOrigin  = "Access-Control-Allow-Origin"
	HeaderExposeHeader = "Access-Control-Expose-Headers"
	HeaderAccessMaxAge = "Access-Control-Max-Age"
)
