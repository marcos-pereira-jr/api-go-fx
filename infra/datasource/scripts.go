package datasource

import (
	"github.com/marcos-pereira-jr/rinha-go-fx/app"
)

type Scripts struct {
	repository *TransacaoRepository
}

func (s *Scripts) Run() {
	s.repository.Insert("1", app.User{
		Id:     "1",
		Saldo:  0,
		Limite: 100000,
	})
	s.repository.Insert("2", app.User{
		Id:     "2",
		Saldo:  0,
		Limite: 80000,
	})
	s.repository.Insert("3", app.User{
		Id:     "3",
		Saldo:  0,
		Limite: 1000000,
	})
	s.repository.Insert("4", app.User{
		Id:     "4",
		Saldo:  0,
		Limite: 10000000,
	})
	s.repository.Insert("5", app.User{
		Id:     "5",
		Saldo:  0,
		Limite: 500000,
	})
}

func NewScripts(
	repository *TransacaoRepository,
) *Scripts {
	return &Scripts{
		repository: repository,
	}
}
