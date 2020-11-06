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


// Comprobar los errores que devuelve al preguntar algo sobre una asignatura
func TestPreguntar(t *testing.T) {
	err := PreRepo.Preguntar("ABC", "¿Esta es una pregunta válida?")
	assert.Nil(t, err)
	err = PreRepo.Preguntar("ABCDEF", "¿Esta es una pregunta válida?")
	assert.NotNil(t, err)
	err = PreRepo.Preguntar("ABC",  "¿Fallará si incluimos una pregunta demasiado larga a la que los usuarios deben contestar?" +
									"¿Fallará si incluimos una pregunta demasiado larga a la que los usuarios deben contestar?" +
									"¿Fallará si incluimos una pregunta demasiado larga a la que los usuarios deben contestar?" +
									"¿Fallará si incluimos una pregunta demasiado larga a la que los usuarios deben contestar?" )
	assert.NotNil(t, err)
}