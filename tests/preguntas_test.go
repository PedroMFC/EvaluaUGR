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
	err = PreRepo.Preguntar("AB", "¿Esta es una pregunta válida también?")
	assert.Nil(t, err)

	assert.Equal(t, 0, PreRepo.Preguntas["AB"][0].Identificador, "Es la primera pregunta enviada")
	assert.Equal(t, 1, PreRepo.Preguntas["AB"][1].Identificador, "Es la segunda pregunta enviada")

}

//Comprobamos que las respuestas se almacenan correctamente
func TestResponder(t *testing.T){
	err := PreRepo.Responder("ABCDEF", 0, "Es una respuesta")  //Nombre de la asignatura no válido
	assert.NotNil(t, err)

	err = PreRepo.Responder("CCC", 3, "Es una respuesta")  //Identificador no válido
	assert.NotNil(t, err)

	err = PreRepo.Responder("CCC", 0, "No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
										"No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
										"No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
										"No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado")  //Respuesta muy larga

	assert.NotNil(t, err)

	err = PreRepo.Responder("CCC", 0, "Es la segunda respuesta")  //Identificador válido
	assert.Nil(t, err)
	assert.Equal(t, 2, len(PreRepo.Preguntas["CCC"][0].Respuestas), "Tiene que tener dos respuestas con la que hemos incluido antes")
	assert.Equal(t, "Es la segunda respuesta", PreRepo.Preguntas["CCC"][0].Respuestas[1].Respuesta, "Vemos que se ha almacenado correctamente")
	assert.Equal(t, 0, len(PreRepo.Preguntas["CCC"][1].Respuestas), "La segunda pregunta no tiene respuestas")
}
