package https

import (

	"github.com/gofiber/fiber/v2"
)

type Router interface {
  Load()
}


func MakeRouter(transacaoRouter *TransacaoRouter) *fiber.App{
  cfg := fiber.Config{
    AppName: "rinha-go by @MarcosPereira expirado @leorcvagas",
    CaseSensitive: true,
  }

  r := fiber.New(cfg)

  transacaoRouter.Load(r)

  return r
}
