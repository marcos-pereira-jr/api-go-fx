package datasource

import (
	"github.com/marcos-pereira-jr/rinha-go-fx/app"
)

type Scripts struct {
	repository *TransacaoRepository
}

func (s *Scripts) Run() {
	s.repository.Insert("1", app.User{
		Id: "1",
	})
}

func NewScripts(
	repository *TransacaoRepository,
) *Scripts {
	return &Scripts{
		repository: repository,
	}
}
