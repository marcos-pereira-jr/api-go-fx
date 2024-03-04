package inmemory

import "github.com/marcos-pereira-jr/rinha-go-fx/app"

type Scripts struct {
	repository app.TransacaoRepository
}

func (s *Scripts) Run() {
	s.repository.InitUser("1")
}

func NewScripts(
	repository app.TransacaoRepository,
) *Scripts {
	return &Scripts{
		repository: repository,
	}
}
