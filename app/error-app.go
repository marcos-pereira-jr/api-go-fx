package app

type ErrorApp struct {
	Message string
}

func (e ErrorApp) Error() string {
	return e.Message
}

type ErrorCredit struct {
	Message string
}

func (e ErrorCredit) Error() string {
	return e.Message
}
