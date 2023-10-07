package main

import (
	"context"

	"github.com/audryus/miniature-octo-tribble/cmd/app"
	"github.com/audryus/miniature-octo-tribble/config"
	"github.com/audryus/miniature-octo-tribble/internal/scheduler"
	"github.com/audryus/miniature-octo-tribble/internal/usecase"
	"github.com/audryus/miniature-octo-tribble/internal/usecase/repo"
	"github.com/audryus/miniature-octo-tribble/pkg/mongodb"
)

func main() {
	ctx := context.Background()
	ctx = config.New(ctx)
	ctx = mongodb.New(ctx)
	ctx = repo.NewEmpresaRepo(ctx)
	ctx = repo.NewVencimentoRepo(ctx)
	ctx = repo.NewSerieRepo(ctx)
	ctx = usecase.NewSerieUC(ctx)
	ctx = usecase.NewEmpresaUC(ctx)
	ctx = usecase.NewOpcoesUC(ctx)

	r := repo.GetVencimentoRepo(ctx)
	sr := usecase.GetSerieUC(ctx)

	vencimento := r.DeleteVencimentoPassado()

	if vencimento != nil {
		sr.DeleteByVencimento(vencimento)
	}

	scheduler.NewJobScheduler(ctx)

	app.Init(ctx)
}
