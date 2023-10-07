package v1

import (
	"encoding/json"

	"github.com/audryus/miniature-octo-tribble/internal/usecase"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo/entity"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) handleEmpresasSemTipo(c *fiber.Ctx) error {
	uc := usecase.GetEmpresaUC(h.ctx)
	empresas, err := uc.GetSemTipo()
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).SendString(err.Error())
	}
	return c.JSON(empresas)
}

func (h *handler) handleEmpresasTipo(c *fiber.Ctx) error {
	uc := usecase.GetEmpresaUC(h.ctx)

	empresas := make([]entity.Empresa, 0)
	if err := json.Unmarshal(c.Body(), &empresas); err != nil {
		code := fiber.StatusBadRequest
		return c.Status(code).SendString(err.Error())
	}
	if len(empresas) > 0 {
		uc.UpdateSemTipo(empresas)
	}

	empresas, err := uc.GetSemTipo()
	if err != nil {
		code := fiber.StatusInternalServerError
		return c.Status(code).SendString(err.Error())
	}
	return c.JSON(empresas)
}
