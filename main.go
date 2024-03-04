package main

import (
	"github.com/marcos-pereira-jr/rinha-go-fx/infra/config"
	"github.com/marcos-pereira-jr/rinha-go-fx/infra/https"
	"github.com/marcos-pereira-jr/rinha-go-fx/infra/inmemory"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		https.Module,
		config.Module,
		inmemory.Module,
		fx.Invoke(func(scripts *inmemory.Scripts) {
			scripts.Run()
		}),
		fx.Invoke(func(*fasthttp.Server) {}),
	)
	app.Run()
}
