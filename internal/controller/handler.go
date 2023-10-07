package controller

import (
	"context"

	v1 "github.com/audryus/miniature-octo-tribble/internal/controller/http/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Handler interface {
	Init() *fiber.App
}

type handler struct {
	ctx context.Context
}

func NewHandler(ctx context.Context) Handler {
	return &handler{
		ctx: ctx,
	}
}

func (h *handler) Init() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	h.api(app)
	return app
}

func (h *handler) api(app *fiber.App) {
	v1 := v1.NewHandler(h.ctx)
	api := app.Group("/api")
	h.healthz(api)
	{
		v1.Init(api)
	}
}
