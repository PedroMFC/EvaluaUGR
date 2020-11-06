package modelspre


import (
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"
	"github.com/rivo/uniseg"
)


//Pregunta contiene los datos de una pregunta formulada
type Pregunta struct {
	Pregunta   string 
	Identificador int
	Respuestas []Respuesta
}


//SetPregunta almacena la pregunta
func (pre *Pregunta) SetPregunta(pregu string) error {
	numCaracteres := uniseg.GraphemeClusterCount(pregu)
	if numCaracteres < 1 || numCaracteres > 300 {
		return &errorspre.ErrorPregunta{" la pregunta no tiene un número de caracteres adecuados"}
	}

	pre.Pregunta = pregu

	return nil
}

