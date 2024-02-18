package app

import "time"

type Saldo struct {
	Total       int       `json:"saldo"`
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

func Extrair(user *User) Extrato {
	var transacoes []TransacaoRecente
	var transacaoRecente *TransacaoRecente
	var trasacoeOriginal []Transacao
	if len(trasacoeOriginal) > 10 {
		trasacoeOriginal = user.Transactions[:10]
	} else {
		trasacoeOriginal = user.Transactions
	}
	for _, transacao := range trasacoeOriginal {
		transacaoRecente = &TransacaoRecente{
			Valor:       transacao.Valor,
			Tipo:        transacao.Tipo,
			Descricao:   transacao.Descricao,
			RealizadoEm: transacao.RealizadoEm,
		}
		transacoes = append(transacoes, *transacaoRecente)
	}

	saldo := &Saldo{
		Total:       user.Saldo,
		DataExtrato: time.Now(),
		Limite:      user.Limite,
	}

	extrato := &Extrato{
		Saldo:      *saldo,
		Transacoes: transacoes,
	}

	return *extrato
}
