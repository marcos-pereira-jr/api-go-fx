package main

import (
	"github.com/marcos-pereira-jr/rinha-go-fx/infra/datasource"
	"github.com/marcos-pereira-jr/rinha-go-fx/infra/https"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func main(){
  app := fx.New(
    https.Module,
    datasource.Module,
    fx.Invoke(func(*fasthttp.Server){}),
    )
  app.Run()
}
