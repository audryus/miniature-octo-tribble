package usecase_test

import (
	"context"
	"testing"

	"github.com/audryus/miniature-octo-tribble/config"
	"github.com/audryus/miniature-octo-tribble/internal/usecase"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo"
	"github.com/audryus/miniature-octo-tribble/pkg/mongodb"
	"github.com/audryus/miniature-octo-tribble/types"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, types.Config, &config.Config{
		Mongo: config.Mongo{
			URL:      "mongodb://root:example@localhost:27017/",
			DATABASE: "tribble",
		},
	})
	ctx = mongodb.New(ctx)
	ctx = repo.NewEmpresaRepo(ctx)
	ctx = usecase.NewEmpresaUC(ctx)
	uc := usecase.GetEmpresaUC(ctx)
	r, err := uc.GetSemTipo()
	assert.NoError(t, err, "Erro ao tentar obter empresas")
	assert.NotNil(t, r, "Sem empresas")
}
