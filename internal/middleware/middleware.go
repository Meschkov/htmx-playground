package middleware

import (
	"net/http"
)

// Middleware is a function that wraps an http.Handler with custom logic.
type Middleware func(http.Handler) http.Handler

// Chain is a helper to build up a pipeline of middlewares, then apply them to a
// final handler.
type Chain struct {
	middlewares []Middleware
}

// Use appends a middleware to the chain.
func (c *Chain) Use(m Middleware) {
	c.middlewares = append(c.middlewares, m)
}

// Then applies the entire chain of middlewares to the final handler in reverse
// order.
func (c *Chain) Then(h http.Handler) http.Handler {
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		h = c.middlewares[i](h)
	}
	return h
}
