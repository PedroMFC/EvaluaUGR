package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
)


// Comprobar que las preguntas tienen el formato adecuado
func TestPreguntaCorrecta(t *testing.T) {
	pregunta := new(modelspre.Pregunta)
	err := pregunta.SetPregunta("¿Esta es una pregunta válida?")
	assert.Nil(t, err)
	err = pregunta.SetPregunta("")
	assert.NotNil(t, err)
	err = pregunta.SetPregunta("¿Fallará si incluimos una pregunta demasiado larga a la que los usuarios deben contestar?" +
							   "¿Fallará si incluimos una pregunta demasiado larga a la que los usuarios deben contestar?" +
							   "¿Fallará si incluimos una pregunta demasiado larga a la que los usuarios deben contestar?" +
							   "¿Fallará si incluimos una pregunta demasiado larga a la que los usuarios deben contestar?" )
	assert.NotNil(t, err)
}
