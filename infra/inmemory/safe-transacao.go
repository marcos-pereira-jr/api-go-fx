package inmemory

import (
	. "github.com/marcos-pereira-jr/rinha-go-fx/app"
)

type Command struct {
	action string
	value  Transacao
	result chan<- interface{}
}

type SafeTransacao struct {
	commandCh  chan Command
	transacoes []Transacao
}

func (s *SafeTransacao) start() {
	for cmd := range s.commandCh {
		switch cmd.action {
		case "add":
			s.transacoes = append([]Transacao{cmd.value}, s.transacoes...)
			if len(s.transacoes) > 10 {
				s.transacoes = s.transacoes[:10]
			}
		case "findOne":
			if 0 != len(s.transacoes) {
				find := s.transacoes[0]
				cmd.result <- find
			} else {
				cmd.result <- nil
			}
		case "list":
			listCopy := make([]Transacao, len(s.transacoes))
			copy(listCopy, s.transacoes)
			cmd.result <- listCopy
		}
	}
}

func (s *SafeTransacao) Add(id string, transacao Transacao) {
	s.commandCh <- Command{action: "add", value: transacao}
}
func (s *SafeTransacao) List(id string) []Transacao {
	resultChan := make(chan interface{})
	s.commandCh <- Command{action: "list", result: resultChan}
	result := <-resultChan
	return result.([]Transacao)
}

func (s *SafeTransacao) FindOne(id string) *Transacao {
	resultChan := make(chan interface{})
	s.commandCh <- Command{action: "findOne", result: resultChan}
	result := <-resultChan
	if result == nil {
		return nil
	}
	transacao := result.(Transacao)
	return &transacao
}

func NewSafeTransacao() *SafeTransacao {
	st := &SafeTransacao{
		commandCh: make(chan Command),
	}
	go st.start()
	return st
}
