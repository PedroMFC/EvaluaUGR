package modelspre

import (
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"
	"github.com/rivo/uniseg"
)


//Respuesta contiene un string que responde a una pregunta formulada
type Respuesta struct {
	Respuesta   string 
}


//SetPregunta almacena la pregunta
func (res *Respuesta) SetRespuesta(respu string) error {
	numCaracteres := uniseg.GraphemeClusterCount(respu)
	if numCaracteres < 1 || numCaracteres > 300 {
		return &errorspre.ErrorPregunta{" la respuesta no tiene un número de caracteres adecuados"}
	}

	res.Respuesta = respu

	return nil
}

