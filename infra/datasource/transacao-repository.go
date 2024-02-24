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

func calcularDebito(lastSaldo int, limite int, transacao app.Transacao) (int, error) {
	saldo := lastSaldo - transacao.Valor
	if saldo < -(limite) {
		return 0, &app.ErrorCredit{Message: "Credito Incosistente"}
	}
	return saldo, nil
}

func calcularCredito(lastSaldo int, transacao app.Transacao) int {
	return lastSaldo + transacao.Valor
}

func (t *TransacaoRepository) InsertTransaction(id string, transacao app.Transacao) (*app.Result, error) {
	var result app.Result
	if transacao.Tipo != "d" && transacao.Tipo != "c" {
		return nil, &app.ErrorCredit{Message: "Credito Incosistente"}

	}
	transacao.RealizadoEm = time.Now()
	lastTrasacao := t.FindLatestTransacao(id, 1)

	user, errorFindUser := t.FindUser(id)
	if errorFindUser != nil {
		return nil, errorFindUser
	}
	var lastSaldo int
	if lastTrasacao != nil {
		lastSaldo = lastTrasacao[0].Saldo
	} else {
		lastSaldo = 0
	}

	var saldo int
	if transacao.Tipo == "c" {
		saldo = calcularCredito(lastSaldo, transacao)
	}

	if transacao.Tipo == "d" {
		var errorCredito error
		saldo, errorCredito = calcularDebito(lastSaldo, user.Limite, transacao)
		if errorCredito != nil {
			return nil, errorCredito
		}
	}

	result.Limite = user.Limite
	result.Saldo = saldo
	transacao.Saldo = saldo

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

func (t *TransacaoRepository) FindLatestTransacao(id string, qtd int64) []*app.Transacao {
	var results []*app.Transacao
	coll := t.dbClient.Database("user").Collection("transacao")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"RealizadoEm", -1}})
	findOptions.SetLimit(qtd)

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
