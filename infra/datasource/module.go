package datasource

import "go.uber.org/fx"

var Module = fx.Provide(
	NewMongoClient,
	NewMongoConfig,
	NewTransacaoRepository,
	NewScripts,
)
