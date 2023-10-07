package scheduler_test

import (
	"context"
	"testing"

	"github.com/audryus/miniature-octo-tribble/config"
	"github.com/audryus/miniature-octo-tribble/internal/scheduler"
	"github.com/audryus/miniature-octo-tribble/internal/usecase"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo"
	"github.com/audryus/miniature-octo-tribble/pkg/mongodb"
	"github.com/audryus/miniature-octo-tribble/types"
	"github.com/stretchr/testify/assert"
)

var ctx context.Context

func contextLoad() {
	ctx = context.Background()
	ctx = context.WithValue(ctx, types.Config, &config.Config{
		Mongo: config.Mongo{
			URL:      "mongodb://root:example@localhost:27017/",
			DATABASE: "tribble",
		},
	})
	ctx = mongodb.New(ctx)
	ctx = repo.NewEmpresaRepo(ctx)
	ctx = repo.NewSerieRepo(ctx)
	ctx = usecase.NewEmpresaUC(ctx)
	ctx = usecase.NewSerieUC(ctx)
}

func TestCopyFile(t *testing.T) {
	contextLoad()
	job := scheduler.NewPriceJob(ctx)
	err := job.CopyFile()
	assert.NoError(t, err, "Erro ao tentar copiar arquivo")
	job.CleanFile()
}
func TestReadFile(t *testing.T) {
	contextLoad()
	job := scheduler.NewPriceJob(ctx)
	err := job.CopyFile()
	assert.NoError(t, err, "Erro ao tentar copiar arquivo")
	records, err := job.ReadFile()

	assert.NoError(t, err, "Erro ao tentar ler arquivo")
	assert.Greater(t, len(records), 0, "Sem linhas carregadas")
	job.CleanFile()
}

func TestIsAtivo(t *testing.T) {
	contextLoad()
	job := scheduler.NewPriceJob(ctx)

	symbols := make([]string, 0)
	symbols = append(symbols, "ABEV3")

	assert.False(t, job.IsAtivo(symbols, "ABEVI842"), "Derivado")
	assert.True(t, job.IsAtivo(symbols, "ABEV3"), "Ativo")
}

func TestRun(t *testing.T) {
	contextLoad()
	job := scheduler.NewPriceJob(ctx)
	job.Run()
}
