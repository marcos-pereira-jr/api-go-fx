package https

import "github.com/gofiber/fiber/v2"

type TransacaoRouter struct {
}

func (p *TransacaoRouter) Load(r *fiber.App) {
	r.Get("/actuator", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func NewTransacaoRouter() *TransacaoRouter {
	return &TransacaoRouter{}
}
