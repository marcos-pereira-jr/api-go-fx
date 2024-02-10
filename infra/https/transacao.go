package https

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransacaoRouter struct {
  dbClient *mongo.Client 
}

func (p *TransacaoRouter) Load(r *fiber.App) {
	r.Get("/actuator", func(c *fiber.Ctx) error {
	coll := p.dbClient.Database("user").Collection("user")
  fmt.Print("TEST")
  _, err := coll.InsertOne(
    context.TODO(),
    bson.D{
        {"animal", "Dog"},
        {"breed", "Beagle"},
    },
  )
    if(err != nil) {
      fmt.Print("Erro",err);
    }
		return c.SendString("teste")
	})
}

func NewTransacaoRouter(
  dbClient *mongo.Client,
) *TransacaoRouter {
	return &TransacaoRouter{
    dbClient: dbClient,
  }
}
