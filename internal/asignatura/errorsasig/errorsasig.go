package errorsasig

import "fmt"

type ErrorAsigntaura struct {
	Mensaje string
}

func (errorVal *ErrorAsigntaura) Error() string {
	return fmt.Sprintf("Algo salió mal: %v", errorVal.Mensaje)
}
