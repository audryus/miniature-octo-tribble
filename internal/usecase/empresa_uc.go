package usecase

import (
	"context"

	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo/entity"
	"github.com/audryus/miniature-octo-tribble/types"
)

type EmpresaUC interface {
	GetSemTipo() ([]entity.Empresa, error)
	UpdateSemTipo([]entity.Empresa) error
	FindAll() []entity.Empresa
	FindByID(symbol string) (*entity.Empresa, error)
	UpdatePrice(string, float64) error
}

type empresaUC struct {
	repo repo.EmpresaRepo
}

func GetEmpresaUC(ctx context.Context) EmpresaUC {
	return ctx.Value(types.EmpresaUC).(EmpresaUC)
}

func NewEmpresaUC(ctx context.Context) context.Context {
	uc := &empresaUC{
		repo: repo.GetEmpresaRepo(ctx),
	}
	return context.WithValue(ctx, types.EmpresaUC, uc)
}
func (u *empresaUC) UpdatePrice(symbol string, price float64) error {
	return u.repo.UpdatePrice(symbol, price)
}

func (u *empresaUC) GetSemTipo() ([]entity.Empresa, error) {
	return u.repo.FindAllWithoutTipo()
}
func (u *empresaUC) UpdateSemTipo(empresas []entity.Empresa) error {
	for _, emp := range empresas {
		if err := u.repo.UpdateSemTipo(emp); err != nil {
			return err
		}
	}

	return nil
}
func (u *empresaUC) FindByID(symbol string) (*entity.Empresa, error) {
	return u.repo.FindByID(symbol)
}
func (u *empresaUC) FindAll() []entity.Empresa {
	return u.repo.FindAll()
}
