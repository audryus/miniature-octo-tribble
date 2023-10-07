package v1

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Init(api fiber.Router)
}

type handler struct {
	ctx context.Context
}

func NewHandler(ctx context.Context) Handler {
	return &handler{
		ctx: ctx,
	}
}

func (h *handler) Init(api fiber.Router) {
	v1 := api.Group("/v1")
	v1.Get("/empresas", h.handleEmpresasSemTipo)
	v1.Patch("/empresas", h.handleEmpresasTipo)
	v1.Get("/ativos", h.handleAtivos)
	v1.Get("/vendaCoberta", h.handleVendaCoberta)
}
