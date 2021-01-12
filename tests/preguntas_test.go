package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/PedroMFC/EvaluaUGR/mocks"
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
	//Definimos el comportamiento que queremos
	PreMapMock = mocks.IPreSaver{} 

	PreMapMock.On("GuardarPregunta", mock.Anything, mock.Anything).Return(nil)
	PreMapMock.On("AsignaturaRegistrada", mock.Anything).Return(true)

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
	//Definimos el comportamiento que queremos
	PreMapMock = mocks.IPreSaver{} 

	PreMapMock.On("ObtenerPregunta", "CCC").Return([]modelspre.Pregunta{
		modelspre.Pregunta{Pregunta: "¿Esta es la primera pregunta?", Identificador: 0, Respuestas: []modelspre.Respuesta{}},
		modelspre.Pregunta{Pregunta: "¿Se ha hecho una segunda pregunta?", Identificador: 1, Respuestas: []modelspre.Respuesta{}},
	})

	PreMapMock.On("ObtenerPregunta", "AAA").Return([]modelspre.Pregunta{})
	PreMapMock.On("AsignaturaRegistrada", mock.Anything).Return(true)

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

//Comprobamos que las respuestas se almacenan correctamente
func TestResponder(t *testing.T){
	//Definimos el comportamiento que queremos
	PreMapMock = mocks.IPreSaver{} 

	PreMapMock.On("Responder", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	PreMapMock.On("AsignaturaRegistrada", mock.Anything).Return(true)

	err := PreRepo.Responder("ABCDEF", 0, "Es una respuesta")  //Nombre de la asignatura no válido
	assert.NotNil(t, err)

	err = PreRepo.Responder("CCC", 0, "No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
										"No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
										"No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado" +
										"No podemos almacenar esta respuesta porque es demasiado larga y no cumple con lo esperado")  //Respuesta muy larga

	assert.NotNil(t, err)

	err = PreRepo.Responder("CCC", 0, "Es la segunda respuesta")  //Identificador válido
	assert.Nil(t, err)
}

// Comprobar que crea una asignatura correctamente
func TestCrearAsignaturaPregunta(t *testing.T) {
	PreMapMock = mocks.IPreSaver{} 

	PreMapMock.On("CrearAsignatura", mock.Anything)

	err := PreRepo.CrearAsignatura("AAA")
	assert.Nil(t, err)
	err = PreRepo.CrearAsignatura("AAAAAA")
	assert.NotNil(t, err)

}