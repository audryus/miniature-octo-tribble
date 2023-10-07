package usecase

import (
	"context"
	"sort"

	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo"
	"github.com/audryus/miniature-octo-tribble/types"
)

type opcoesUC struct {
	empUC    EmpresaUC
	serieUC  SerieUC
	repoVenc *repo.VencimentoRepo
}

type ativoDTO struct {
	Symbol string
	Price  float64
	Opcoes []derivadoDTO
}

type derivadoDTO struct {
	Symbol     string
	Strike     float64
	Price      float64
	Percent    float64
	Vencimento string
}

func NewOpcoesUC(ctx context.Context) context.Context {
	empUC := GetEmpresaUC(ctx)
	serieUC := GetSerieUC(ctx)
	repoVenc := repo.GetVencimentoRepo(ctx)
	uc := &opcoesUC{
		empUC:    empUC,
		serieUC:  serieUC,
		repoVenc: repoVenc,
	}
	return context.WithValue(ctx, types.OpcoesUC, uc)
}
func GetOpcoesUC(ctx context.Context) *opcoesUC {
	return ctx.Value(types.OpcoesUC).(*opcoesUC)
}

func (u *opcoesUC) FindVendaCoberta(empresas []string) ([]ativoDTO, error) {
	acoes := make([]ativoDTO, 0)
	vencimentos := u.repoVenc.GetNextTwo()

	for _, empresa := range empresas {
		emp, err := u.empUC.FindByID(empresa)
		if err != nil {
			return nil, err
		}
		opcoes := make([]derivadoDTO, 0)

		for i, vencimento := range vencimentos {
			series := u.serieUC.FindAllByEmpresaAndVencimento(*emp, vencimento)
			for _, serie := range series {
				ePrice := emp.Price
				sPrice := serie.Price
				sStrike := serie.Strike

				if sStrike > ePrice {
					continue
				}

				cut := 2.0

				if i > 0 {
					cut = cut * float64(i)
				}

				percent := (((sStrike - ePrice + sPrice) / ePrice) * 100)

				if percent < cut {
					continue
				}

				opcoes = append(opcoes, derivadoDTO{
					Symbol:     serie.Serie,
					Strike:     serie.Strike,
					Price:      serie.Price,
					Percent:    percent,
					Vencimento: serie.Vencimento,
				})
			}
		}
		sort.SliceStable(opcoes, func(i, j int) bool {
			return opcoes[i].Percent > opcoes[j].Percent
		})

		acoes = append(acoes, ativoDTO{
			Symbol: emp.Symbol,
			Price:  emp.Price,
			Opcoes: opcoes,
		})
	}

	return acoes, nil

}
