package datasource

import (
	"context"
	"fmt"
	"github.com/marcos-pereira-jr/rinha-go-fx/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransacaoRepository struct {
	dbClient *mongo.Client
}

func (t *TransacaoRepository) InsertTransaction(id string, transacao app.Transacao) error {
	user, errorFind := t.FindUser(id)
	if errorFind != nil {
		return errorFind
	}
	user.Transactions = append(user.Transactions, transacao)
	coll := t.dbClient.Database("user").Collection("user")
	update := bson.M{"$push": bson.M{"Transactions": transacao}}
	filter := bson.M{"id": id}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedDocument bson.M
	err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		fmt.Print("Erro", err)
	}
	return nil
}

func (t *TransacaoRepository) Insert(id string, user app.User) error {
	coll := t.dbClient.Database("user").Collection("user")
	_, err := coll.InsertOne(
		context.TODO(),
		user,
	)
	if err != nil {
		fmt.Print("Erro", err)
	}
	return nil
}

func (t *TransacaoRepository) FindUser(id string) (*app.User, error) {
	coll := t.dbClient.Database("user").Collection("user")
	var result app.User
	err := coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	if err != nil {
		return nil, &app.ErrorApp{Message: "Not Found"}
	}
	return &result, nil
}

func NewTransacaoRepository(
	dbClient *mongo.Client,
) *TransacaoRepository {
	return &TransacaoRepository{
		dbClient: dbClient,
	}
}
