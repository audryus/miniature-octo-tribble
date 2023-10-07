package v1

import (
	"github.com/audryus/miniature-octo-tribble/internal/usecase"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo"
	"github.com/gofiber/fiber/v2"
)

type ativo struct {
	Symbol string `json:"symbol"`
}

func (h handler) handleAtivos(c *fiber.Ctx) error {
	uc := usecase.GetEmpresaUC(h.ctx)
	sUC := usecase.GetSerieUC(h.ctx)
	vUC := repo.GetVencimentoRepo(h.ctx)

	ativos := make([]ativo, 0)

	empresas := uc.FindAll()
	for _, empresa := range empresas {
		ativos = append(ativos, ativo{
			Symbol: empresa.Symbol,
		})
		vencimentos := vUC.GetNextTwo()

		for _, venc := range vencimentos {
			series := sUC.FindAllByEmpresaAndVencimento(empresa, venc)
			for _, serie := range series {
				ativos = append(ativos, ativo{
					Symbol: serie.Serie,
				})
			}
		}
	}
	return c.JSON(ativos)
}
