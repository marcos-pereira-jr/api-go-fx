package datasource

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/marcos-pereira-jr/rinha-go-fx/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransacaoRepository struct {
	dbClient *mongo.Client
	users    map[string]app.User
}

func calcularDebito(user *app.User, transacao app.Transacao) (int, error) {
	saldo := user.Saldo - transacao.Valor
	if saldo < -(user.Limite) {
		return 0, &app.ErrorCredit{Message: "Credito Incosistente"}
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

	coll := t.dbClient.Database("user").Collection("transacao")
	_, err := coll.InsertOne(
		context.TODO(),
		transacao,
	)
	if err != nil {
		fmt.Print("Erro", err)
	}
	return &result, nil
}

func (t *TransacaoRepository) Insert(id string, user app.User) error {
	t.users[id] = user
	return nil
}

func (t *TransacaoRepository) FindUser(id string) (*app.User, error) {
	if user, exists := t.users[id]; exists {
		return &user, nil
	}
	return nil, &app.ErrorApp{Message: "Not Found"}
}

func (t *TransacaoRepository) FindLatestTransacao(id string) []*app.Transacao {
	var results []*app.Transacao
	coll := t.dbClient.Database("user").Collection("transacao")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"RealizadoEm", -1}})
	findOptions.SetLimit(10)

	cur, _ := coll.Find(context.TODO(), bson.D{{}}, findOptions)
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var elem app.Transacao
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	return results
}

func NewTransacaoRepository(
	dbClient *mongo.Client,
) *TransacaoRepository {
	return &TransacaoRepository{
		dbClient: dbClient,
		users:    make(map[string]app.User),
	}
}
