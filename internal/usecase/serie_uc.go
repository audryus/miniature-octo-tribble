package usecase

import (
	"context"

	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo/entity"
	"github.com/audryus/miniature-octo-tribble/types"
)

type SerieUC interface {
	SaveAll(series []entity.Serie)
	UpdatePrice(serie string, price float64) error
	DeleteByVencimento(vencimento *entity.Vencimento) error
	FindAllByEmpresaAndVencimento(empresa entity.Empresa, vencimento entity.Vencimento) []entity.Serie
}
type serieUC struct {
	ctx context.Context
	r   repo.SerieRepo
}

func GetSerieUC(ctx context.Context) SerieUC {
	return ctx.Value(types.SerieUC).(SerieUC)
}
func NewSerieUC(ctx context.Context) context.Context {
	uc := &serieUC{
		ctx: ctx,
		r:   repo.GetSerieRepo(ctx),
	}

	return context.WithValue(ctx, types.SerieUC, uc)
}

func (u *serieUC) UpdatePrice(serie string, price float64) error {
	return u.r.UpdatePrice(serie, price)
}
func (u *serieUC) DeleteByVencimento(vencimento *entity.Vencimento) error {
	return u.r.DeleteByVencimento(vencimento)
}

func (u *serieUC) SaveAll(series []entity.Serie) {
	vRepo := repo.GetVencimentoRepo(u.ctx)
	vRepo.SaveAll(series)
	u.r.SaveAll(series)
}

func (u *serieUC) FindAllByEmpresaAndVencimento(empresa entity.Empresa, vencimento entity.Vencimento) []entity.Serie {
	return u.r.FindAllByEmpresaAndVencimento(empresa, vencimento)
}
