package https

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/marcos-pereira-jr/rinha-go-fx/app"
)

type TransacaoRouter struct {
	repository app.TransacaoRepository
}

func (p *TransacaoRouter) Load(r *fiber.App) {
	r.Post("/clientes/:id/transacoes", func(c *fiber.Ctx) error {
		var body app.Transacao
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"erro": "bad request"})
		}
		result, err := p.repository.AddTransacao(c.Params("id"), body)
		if err != nil {
			var notFound *app.ErrorApp
			if errors.As(err, &notFound) {
				return c.Status(fiber.StatusNotFound).Next()
			}
			var errorCredit *app.ErrorCredit
			if errors.As(err, &errorCredit) {
				return c.Status(fiber.StatusUnprocessableEntity).SendString("Bad Request")
			}

		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	r.Get("/clientes/:id/extrato", func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := p.repository.Extrato(id)
		if err != nil {
			var notFound *app.ErrorApp
			if errors.As(err, &notFound) {
				return c.Status(fiber.StatusNotFound).Next()
			}
		}
		return c.Status(fiber.StatusOK).JSON(app.Extrair(result))
	})
}

func NewTransacaoRouter(
	repository app.TransacaoRepository,
) *TransacaoRouter {
	return &TransacaoRouter{
		repository: repository,
	}
}
