package app

type ErrorApp struct {
	Message string
}

func (e ErrorApp) Error() string {
	return e.Message
}
