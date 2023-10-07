package app

import (
	"context"
	"fmt"
	"log"

	"github.com/audryus/miniature-octo-tribble/config"
	"github.com/audryus/miniature-octo-tribble/internal/controller"
)

func Init(ctx context.Context) {
	handlers := controller.NewHandler(ctx)
	app := handlers.Init()
	cfg := config.Get(ctx)

	log.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", cfg.HTTP.Port)))
}
