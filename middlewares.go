package middlewares

import "net/http"

type Middlewares struct {
	middlewares []*Middleware
}
type MiddlewareHandler func(http.Handler) http.Handler

type Middleware struct {
	Name    string
	Handler MiddlewareHandler

	middlewares *Middlewares
	before      string
	after       string
}

// Use use middleware
func (middlewares *Middlewares) Use(name string, handler MiddlewareHandler) {
	middlewares.middlewares = append(middlewares.middlewares, &Middleware{
		middlewares: middlewares,
		Name:        name,
		Handler:     handler,
	})
}

// Before insert middleware before name
func (middlewares *Middlewares) Before(name string) Middleware {
	return Middleware{
		middlewares: middlewares,
		before:      name,
	}
}

// After insert middleware after name
func (middlewares *Middlewares) After(name string) Middleware {
	return Middleware{
		middlewares: middlewares,
		after:       name,
	}
}

// Use use middleware
func (middleware Middleware) Use(name string, handler MiddlewareHandler) {
	middleware.Name = name
	middleware.Handler = handler
	if middleware.middlewares != nil {
		middleware.middlewares.middlewares = append(middleware.middlewares.middlewares, &middleware)
	}
}