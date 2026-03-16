package middleware

import "ecommerce/config"

type MiddlewareHandler struct {
	cnf *config.Config
}

func NewMiddlewareHandler(cnf *config.Config) *MiddlewareHandler {
	return &MiddlewareHandler{
		cnf: cnf,
	}
}
