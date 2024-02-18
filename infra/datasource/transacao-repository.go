package datasource

import (
	"context"
	"fmt"
	"time"

	"github.com/marcos-pereira-jr/rinha-go-fx/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransacaoRepository struct {
	dbClient *mongo.Client
}

func calcularDebito(user *app.User, transacao app.Transacao) (int, error) {
	saldo := user.Saldo - transacao.Valor
	if saldo < -(user.Limite) {
		return 0, &app.ErrorApp{Message: "Not Found"}
	}
	return saldo, nil
}

func calcularCredito(user *app.User, transacao app.Transacao) int {
	return user.Saldo + transacao.Valor
}

func (t *TransacaoRepository) InsertTransaction(id string, transacao app.Transacao) (*app.Result, error) {
	var result app.Result
	transacao.RealizadoEm = time.Now()
	user, errorFind := t.FindUser(id)
	if errorFind != nil {
		return nil, errorFind
	}
	user.Transactions = append(user.Transactions, transacao)
	var saldo int
	if transacao.Tipo == "c" {
		saldo = calcularCredito(user, transacao)
	}

	if transacao.Tipo == "d" {
		var errorCredito error
		saldo, errorCredito = calcularDebito(user, transacao)
		if errorCredito != nil {
			return nil, errorCredito
		}
	}

	result.Limite = user.Limite
	result.Saldo = saldo
	user.Saldo = saldo

	coll := t.dbClient.Database("user").Collection("user")
	update := bson.M{
		"$push": bson.M{"Transactions": transacao},
		"$set":  bson.M{"Saldo": saldo},
	}

	filter := bson.M{"id": id}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedDocument bson.M
	err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		fmt.Print("Erro", err)
	}
	return &result, nil
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
	err := coll.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&result)
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
