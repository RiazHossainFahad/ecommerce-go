package middleware

import (
	"net/http"
)

type Middlware func(http.Handler) http.Handler // type function as like middleware signature

type Manager struct {
	globalMiddlewares []Middlware // nil array of middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middlware, 0), // empty slice
	}
}

func (manager *Manager) Use(middlewares ...Middlware) {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
}

func (manager *Manager) With(next http.Handler, middlewares ...Middlware) http.Handler {
	n := next

	// middleware(middleware(next))
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range manager.globalMiddlewares {
		n = globalMiddleware(n)
	}

	return n
}
