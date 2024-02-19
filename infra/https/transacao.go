package https

import (
	"errors"

	"github.com/marcos-pereira-jr/rinha-go-fx/app"
	"github.com/marcos-pereira-jr/rinha-go-fx/infra/datasource"

	"github.com/gofiber/fiber/v2"
)

type TransacaoRouter struct {
	repository *datasource.TransacaoRepository
}

func (p *TransacaoRouter) Load(r *fiber.App) {
	r.Post("/clientes/:id/transacoes", func(c *fiber.Ctx) error {
		var body app.Transacao
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": "bad request"})
		}
		result, err := p.repository.InsertTransaction(c.Params("id"), body)
		if err != nil {
			var notFound *app.ErrorApp
			if errors.As(err, &notFound) {
				return c.Status(fiber.StatusNotFound).Next()
			}
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	r.Get("/clientes/:id/extrato", func(c *fiber.Ctx) error {
		result, err := p.repository.FindUser(c.Params("id"))
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
	repository *datasource.TransacaoRepository,
) *TransacaoRouter {
	return &TransacaoRouter{
		repository: repository,
	}
}
