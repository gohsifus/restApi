// Package errs Для определения пользовательских ошибок
package errs

type errorType uint

const (
	// UnHandledErr Прочие ошибки должны возвращать статус 500
	UnHandledErr = errorType(500)
	// IncorrectDataErr Ошибки формата данных возвращают статус 400
	IncorrectDataErr = errorType(400)
	// BusinessLogicErr Ошибки бизнес логики возвращают 503
	BusinessLogicErr = errorType(503)
)

// CustomError тип реализующий интерфейс Error
type CustomError struct {
	status        errorType
	originalError error
}

func (c CustomError) Error() string {
	return c.originalError.Error()
}

//New ...
func New(err error, status errorType) CustomError {
	return CustomError{
		status:        status,
		originalError: err,
	}
}

// Status вернет статус связанный с ошибкой (который нужно вернуть)
func (c CustomError) Status() int {
	return int(c.status)
}

// Wrap оборачивает стандартную ошибку и присваивает ей статус UnHandled
func Wrap(err error) CustomError {
	if err, ok := err.(CustomError); ok {
		return err
	}

	return New(err, UnHandledErr)
}
