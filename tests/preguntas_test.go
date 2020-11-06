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


// Comprobar que devuelve de manera correcta las preguntas
func TestGetPreguntas(t *testing.T) {
	pre, err := PreRepo.GetPreguntas("AAAAAA")
	assert.NotNil(t, err)
	pre, err = PreRepo.GetPreguntas("CCC")
	assert.Nil(t, err)
	//Comprobamos que son iguales a las que hemos inicializado
	assert.Equal(t, "¿Esta es la primera pregunta?", pre[0].Pregunta, "Debe coincidir la primera pregunta")
	assert.Equal(t, "¿Se ha hecho una segunda pregunta?", pre[1].Pregunta, "Y también debe coincidir la segunda")

	//Comprobamos que si no está la asinatura está vacío
	pre, err = PreRepo.GetPreguntas("AAA")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(pre), "El array de preguntas tiene que estar vacío")
}

// Comprobar que las respuestas tienen el formato adecuado
func TestRespuestaCorrecta(t *testing.T) {
	respuesta := new(modelspre.Respuesta)
	err := respuesta.SetRespuesta("Esta es una respuesta adecuada")
	assert.Nil(t, err)
	err = respuesta.SetRespuesta("")
	assert.NotNil(t, err)
	err = respuesta.SetRespuesta("No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
							   "No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
							   "No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
							   "No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" )
	assert.NotNil(t, err)
}


//Comprobamos que se asignan bien los identificadores al enviar la pregunta
func TestIdentificadoresPregunta(t *testing.T) {
	err := PreRepo.Preguntar("AB", "¿Esta es una pregunta válida?")
	assert.Nil(t, err)
	err = PreRepo.Preguntar("AB", "¿Esta es una pregunta es válida también?")
	assert.Nil(t, err)

	assert.Equal(t, 0, PreRepo.Preguntas["AB"][0].Identificador, "Es la primera pregunta enviada")
	assert.Equal(t, 1, PreRepo.Preguntas["AB"][1].Identificador, "Es la segunda pregunta enviada")

}
