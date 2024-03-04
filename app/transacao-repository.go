package app

type TransacaoRepository interface {
	AddTransacao(id string, trasacao Transacao) (*Result, error)
	Extrato(id string) ([]Transacao, error)
	InitUser(id string)
}
