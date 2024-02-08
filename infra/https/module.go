package https
import "go.uber.org/fx"

var Module = fx.Provide(
  NewTransacaoRouter,
  MakeRouter,
  NewServer,
  )
