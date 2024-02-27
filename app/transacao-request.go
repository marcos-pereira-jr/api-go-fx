package app

import "time"

type Transacao struct {
	Valor       int    `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	IdUser      string
	Saldo       int
	RealizadoEm time.Time
}

type User struct {
	Id     string
	Saldo  int
	Limite int
}
