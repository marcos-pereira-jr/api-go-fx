package https

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcos-pereira-jr/rinha-go-fx/infra/config"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func NewServer(
	lifecycle fx.Lifecycle,
	router *fiber.App,
	config *config.Config,
) *fasthttp.Server {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := router.Listen("localhost:" + config.Port); err != nil {
					fmt.Printf("Error starting the server: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Print("Stopping the server...")

			return router.ShutdownWithContext(ctx)
		},
	})
	return router.Server()
}
