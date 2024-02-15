package app

type Transacao struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type User struct {
	Id           string
	Transactions []Transacao
}
