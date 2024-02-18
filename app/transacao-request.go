package app

import "time"

type Transacao struct {
	Valor       int    `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadoEm time.Time
}

type User struct {
	Id           string
	Transactions []Transacao
	Saldo        int
	Limite       int
}
