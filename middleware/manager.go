package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler // type function as like middleware signature

type Manager struct {
	globalMiddlewares []Middleware // nil array of middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0), // empty slice
	}
}

func (manager *Manager) Use(middlewares ...Middleware) {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
}

func (manager *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	// need to pass the data reverse order [Preflight, HandleCors]
	// HandleCors(Preflight(next))
	// for _, middleware := range middlewares {
	// 	n = middleware(n)j
	// }

	for i := len(middlewares) - 1; i >= 0; i-- {
		n = middlewares[i](n)
	}

	return n
}

func (manager *Manager) WrapMux(next http.Handler) http.Handler {
	n := next

	// need to pass the data reverse order [Preflight, HandleCors]
	// HandleCors(Preflight(next))
	// for _, middleware := range manager.globalMiddlewares {
	// 	n = middleware(n)
	// }

	// from top to bottom - [HandleCors, Preflight]
	// HandleCors(Preflight(next))
	for i := len(manager.globalMiddlewares) - 1; i >= 0; i-- {
		n = manager.globalMiddlewares[i](n)
	}

	return n
}
