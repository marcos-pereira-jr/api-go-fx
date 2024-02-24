package app

import (
	"time"
)

type Saldo struct {
	Total       int       `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      int       `json:"limite"`
}

type TransacaoRecente struct {
	Valor       int       `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadoEm time.Time `json:"realizada_em"`
}

type Extrato struct {
	Saldo      Saldo              `json:"saldo"`
	Transacoes []TransacaoRecente `json:"ultimas_transacoes"`
}

func Extrair(user *User, trasacoeOriginal []*Transacao) Extrato {
	var transacoes []TransacaoRecente
	var transacaoRecente *TransacaoRecente
	var total int
	for i, transacao := range trasacoeOriginal {
		if i == 0 {
			total = transacao.Saldo
		}
		transacaoRecente = &TransacaoRecente{
			Valor:       transacao.Valor,
			Tipo:        transacao.Tipo,
			Descricao:   transacao.Descricao,
			RealizadoEm: transacao.RealizadoEm,
		}
		transacoes = append(transacoes, *transacaoRecente)
	}

	saldo := &Saldo{
		Total:       total,
		DataExtrato: time.Now(),
		Limite:      user.Limite,
	}

	extrato := &Extrato{
		Saldo:      *saldo,
		Transacoes: transacoes,
	}

	return *extrato
}
