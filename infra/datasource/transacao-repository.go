package datasource

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/marcos-pereira-jr/rinha-go-fx/app"
	"github.com/marcos-pereira-jr/rinha-go-fx/infra/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransacaoRepository struct {
	dbClient *mongo.Client
	users    map[string]app.User
	config   *config.Config
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll := t.dbClient.Database("user").Collection("transacao")
	var result app.Result
	if transacao.Tipo != "d" && transacao.Tipo != "c" {
		return nil, &app.ErrorCredit{Message: "Credito Incosistente"}

	}
	transacao.RealizadoEm = time.Now()
	user, errorFindUser := t.FindUser(id)
	if errorFindUser != nil {
		return nil, errorFindUser
	}

	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"RealizadoEm", -1}})
	var lastTrasacao app.Transacao
	notFound := coll.FindOne(ctx, bson.D{{"iduser", id}}, findOptions).Decode(&lastTrasacao)

	var lastSaldo int
	if notFound != nil {
		lastSaldo = lastTrasacao.Saldo
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
	transacao.IdUser = id
	go t.AsyncInsert(transacao)
	return &result, nil
}

func (t *TransacaoRepository) AsyncInsert(transacao app.Transacao) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	coll := t.dbClient.Database("user").Collection("transacao")

	_, err := coll.InsertOne(
		ctx,
		transacao,
	)
	if err != nil {
		fmt.Print("Erro", err)
	}
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

	cur, _ := coll.Find(context.TODO(), bson.D{{"iduser", id}}, findOptions)
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
	config *config.Config,
) *TransacaoRepository {
	return &TransacaoRepository{
		dbClient: dbClient,
		users:    make(map[string]app.User),
		config:   config,
	}
}
