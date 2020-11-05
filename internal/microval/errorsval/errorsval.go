package errorsval

import "fmt"

type ErrorValoracion struct {
	Mensaje string
}

func (errorVal *ErrorValoracion) Error() string {
	return fmt.Sprintf("Algo salió mal en la valoración: %v", errorVal.Mensaje)
}
