package errorspre

import "fmt"

type ErrorPregunta struct {
	Mensaje string
}

func (errorPre *ErrorPregunta) Error() string {
	return fmt.Sprintf("Algo salió mal con la pregunta: %v", errorPre.Mensaje)
}
