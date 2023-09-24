package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
	gorilla_mux "github.com/gorilla/mux"
)

// Mux defines the standard Multiplexer for http Request
type (
	Mux interface {
		// ServeHTTP
		http.Handler

		// URLParser returns the parsing method used by
		// the multiplexer
		URLParser() URLParser

		// Default Handler Method is all that is needed
		Handler(method, url string, fn http.Handler)
	}
)

type (
	URLParams map[string]string
	URLParser interface {
		Parse(r *http.Request) URLParams
	}
)

func (up URLParams) Param(key string) (value string, found bool) {
	value, found = up[key]
	return
}

// go-chi muxer which is default multiplexer for go-base
type (
	chiMuxer         struct{ *chi.Mux }
	chiURLParser     struct{}
	DefaultMuxOption func(*chiMuxer)
)

func (cup *chiURLParser) Parse(r *http.Request) URLParams {
	rcx := chi.RouteContext(r.Context())
	par := make(URLParams)

	for ix, key := range rcx.URLParams.Keys {
		par[key] = rcx.URLParams.Values[ix]
	}

	return par
}

func (mx *chiMuxer) URLParser() URLParser {
	return &chiURLParser{}
}

func (mx *chiMuxer) Handler(method, url string, fn http.Handler) {
	mx.Method(method, url, fn)
}

func NewDefaultMux(opts ...DefaultMuxOption) Mux {
	mx := &chiMuxer{chi.NewMux()}
	for _, o := range opts {
		o(mx)
	}
	return mx
}

// gorilla mux implementation
// this here is given as an example on how to use a custom mux
// the expectation is that the application will provide its own multiplexer
// if required and not the default chi multiplexer

type (
	gmux          struct{ *gorilla_mux.Router }
	gmuxURLParser struct{}
)

func (gu *gmuxURLParser) Parse(r *http.Request) URLParams {
	return mux.Vars(r)
}

func (gm *gmux) URLParser() URLParser {
	return &gmuxURLParser{}
}

func (gm *gmux) Handler(method, url string, fn http.Handler) {
	gm.Router.Handle(url, fn).Methods(method)
}

func NewGorillaMux() Mux { return &gmux{gorilla_mux.NewRouter()} }
