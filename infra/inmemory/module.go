package inmemory

import "go.uber.org/fx"

var Module = fx.Provide(
	NewTransacaoRepository,
	NewScripts,
)
