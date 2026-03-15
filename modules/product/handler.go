package product

import "ecommerce/config"

type Handler struct {
	cnf *config.Config
}

func NewHandler(cnf *config.Config) *Handler {
	return &Handler{
		cnf: cnf,
	}
}
