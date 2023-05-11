package middleware

import (
	"net/http"
)

type (
	// Middleware receives and returns http.HandlerFunc
	Middleware = func(http.HandlerFunc) http.HandlerFunc

	// Set of Middleware
	Set struct {
		middlewares []Middleware
	}
)

// NewSet creates a new middleware set
func NewSet(middlewares ...Middleware) *Set {
	if middlewares != nil {
		return &Set{append(make([]Middleware, 0, len(middlewares)), middlewares...)}
	}
	return &Set{}
}

// Use adds the given standard middlewares to middleware set
func (s *Set) Use(middlewares ...Middleware) {
	s.middlewares = append(s.middlewares, middlewares...)
}

// Append the given params to a safe copy of current middleware set
// and returns it as new middleware set
func (s *Set) Append(middlewares ...Middleware) *Set {
	newSet := &Set{make([]Middleware, 0, len(s.middlewares)+len(middlewares))}
	newSet.middlewares = append(newSet.middlewares, s.middlewares...)
	newSet.middlewares = append(newSet.middlewares, middlewares...)
	return newSet
}

// Prepend the given params to a copy of current middleware set
// and returns it as a new middleware set
func (s *Set) Prepend(middlewares ...Middleware) *Set {
	newSet := &Set{make([]Middleware, 0, len(s.middlewares)+len(middlewares))}
	newSet.middlewares = append(newSet.middlewares, middlewares...)
	newSet.middlewares = append(newSet.middlewares, s.middlewares...)
	return newSet
}

// HandlerFunc wraps the given http.HandlerFunc with middlewares and returns it as http.Handler
func (s *Set) HandlerFunc(handler http.HandlerFunc) http.Handler {
	if len(s.middlewares) == 0 {
		return handler
	}
	// wrap handler with all middlewares starting from the last added
	final := handler
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		final = s.middlewares[i](final)
	}
	return final
}

// Handler wraps the given http.Handler with middlewares and returns it
func (s *Set) Handler(handler http.Handler) http.Handler {
	if hf, ok := handler.(http.HandlerFunc); ok {
		return s.HandlerFunc(hf)
	}
	return s.HandlerFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}
