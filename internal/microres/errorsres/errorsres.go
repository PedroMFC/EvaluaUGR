package errorsres

import "fmt"

type ErrorResenia struct {
	Mensaje string
}

func (errorVal *ErrorResenia) Error() string {
	return fmt.Sprintf("Algo salió mal con la reseña: %v", errorVal.Mensaje)
}
