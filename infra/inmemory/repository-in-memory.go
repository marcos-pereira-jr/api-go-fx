package inmemory

import . "github.com/marcos-pereira-jr/rinha-go-fx/app"

// TODO: Lista dinamica
type TransacaoRepositoryInMemory struct {
	users map[string]SafeTransacao
}

func (r *TransacaoRepositoryInMemory) FindSafeTransacao(id string) (*SafeTransacao, error) {
	if user, exists := r.users[id]; exists {
		return &user, nil
	}
	return nil, &ErrorApp{Message: "Not Found"}
}

func (r *TransacaoRepositoryInMemory) Extrato(id string) ([]Transacao, error) {
	safeTransacao, error := r.FindSafeTransacao(id)
	if error != nil {
		return nil, error
	}
	return safeTransacao.List(id), nil
}

func (r *TransacaoRepositoryInMemory) FindLaster(id string) *Transacao {
	safeTransacao, _ := r.FindSafeTransacao(id)
	return safeTransacao.FindOne(id)
}

func (r *TransacaoRepositoryInMemory) AddTransacao(id string, trasacao Transacao) (*Result, error) {
	safeTransacao, error := r.FindSafeTransacao(id)
	if error != nil {
		return nil, error
	}
	laster := r.FindLaster(id)
	if laster != nil {
		trasacao.Saldo = laster.Saldo + trasacao.Valor
	} else {
		trasacao.Saldo = trasacao.Valor
	}
	go safeTransacao.Add(id, trasacao)
	return &Result{Limite: 1223, Saldo: trasacao.Saldo}, nil
}

func (r *TransacaoRepositoryInMemory) InitUser(id string) {
	r.users[id] = *NewSafeTransacao()
}

func NewTransacaoRepository() TransacaoRepository {
	return &TransacaoRepositoryInMemory{
		users: make(map[string]SafeTransacao),
	}
}
