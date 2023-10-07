package v1

import (
	"strings"

	"github.com/audryus/miniature-octo-tribble/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) handleVendaCoberta(c *fiber.Ctx) error {
	empresas := c.Query("empresas", "")
	if len(empresas) == 0 {
		return c.Status(fiber.StatusNotAcceptable).SendString(
			"Parametro 'empresas' obrigatorio.")
	}
	uc := usecase.GetOpcoesUC(h.ctx)
	papeis, err := uc.FindVendaCoberta(strings.Split(empresas, ","))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(
			"Empresa não encontrada.")
	}

	return c.JSON(papeis)
}
