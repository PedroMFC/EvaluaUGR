package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sort"
	//"log"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
)

func TestGuardaPreguntaMap(t *testing.T) {
	pre := new(modelspre.Pregunta)
	pre.SetPregunta("Una pregunta")

	//Guardamos donde ya hay registro
	PreMap.GuardarPregunta("CCC", pre)
	//Vemos que solo tiene una asignatura
	assert.Equal(t, 1, len(PreMap.Preguntas), "Solo hay una asignatura")
	//Vemos que tiene 3 preguntas
	assert.Equal(t, 3, len(PreMap.Preguntas["CCC"]), "Tiene tres preguntas")


	//Guardamos en una nueva
	PreMap.CrearAsignatura("BBB")
	PreMap.GuardarPregunta("BBB", pre)
	//Vemos que ahora hay dos asignaturas
	assert.Equal(t, 2, len(PreMap.Preguntas), "Hay 2 una asignatura")
	//Vemos que tiene 1 pregunta1
	assert.Equal(t, 1, len(PreMap.Preguntas["BBB"]), "Tiene 1 pregunta")

}

func TestAsignarIdentificadores(t *testing.T) {
	assert.Equal(t, 0, PreMap.Preguntas["CCC"][0].Identificador, "Primera pregunta enviada")
	assert.Equal(t, 1, PreMap.Preguntas["CCC"][1].Identificador, "Segunda pregunta enviada")
	assert.Equal(t, 2, PreMap.Preguntas["CCC"][2].Identificador, "Tercera pregunta enviada")

}

func TestObtenerPreguntasMap(t *testing.T) {
	//Obtenemos las preguntas de la primera
	//preguntas := PreMap.ObtenerPregunta("CCC")
	arr := PreMap.ObtenerPregunta("CCC")
	preguntas :=make([]modelspre.Pregunta, len(arr))
	copy(preguntas, PreMap.ObtenerPregunta("CCC"))
	// Las ordenamos para ver que las devuelve bien
	sort.Slice(preguntas, func(i,j int) bool {
		return preguntas[i].Pregunta < preguntas[j].Pregunta
	})

	//Vemos que tiene 3
	assert.Equal(t, 3, len(preguntas), "CCC tiene tres preguntas")
	//Vemos que son las que hemos introducido
	assert.Equal(t, "Una pregunta", preguntas[0].Pregunta, "Primera pregunta")
	assert.Equal(t, "¿Esta es la primera pregunta?", preguntas[1].Pregunta, "Segunda pregunta")
	assert.Equal(t, "¿Se ha hecho una segunda pregunta?", preguntas[2].Pregunta, "Tercera pregunta")

	//Repetimos el proceso con BBB
	preguntas = PreMap.ObtenerPregunta("BBB")
	// Las ordenamos para ver que las devuelve bien
	sort.Slice(preguntas, func(i,j int) bool {
		return preguntas[i].Pregunta < preguntas[j].Pregunta
	})

	assert.Equal(t, 1, len(preguntas), "BBB tiene una preguntas")
	assert.Equal(t, "Una pregunta", preguntas[0].Pregunta, "Primera pregunta")

	//Vemos que si no está devuelve uno vacío
	preguntas = PreMap.ObtenerPregunta("DFGF")
	assert.Equal(t, 0, len(preguntas), "DFGF no tiene una pregunta")
}

func TestResponderMap(t *testing.T) {
	//Probamos con identificadores válidos
	res := new(modelspre.Respuesta)
	res.SetRespuesta("Respuesta de prueba")

	err := PreMap.Responder("CCC", 0, res)
	assert.Nil(t, err)
	respuestas := PreMap.Preguntas["CCC"][0].Respuestas
	sort.Slice(respuestas, func(i,j int) bool {
		return respuestas[i].Respuesta < respuestas[j].Respuesta
	})

	assert.Equal(t, "Respuesta de prueba", respuestas[0].Respuesta, "Es la segunda")
	assert.Equal(t, "Sí, esta es la primera pregunta.", respuestas[1].Respuesta, "Ya hay una")

	err = PreMap.Responder("BBB", 0, res)
	assert.Nil(t, err)
	respuestas = PreMap.Preguntas["BBB"][0].Respuestas
	assert.Equal(t, "Respuesta de prueba", respuestas[0].Respuesta, "Es la segunda")

	//Probamos con identificador incorrecto
	err = PreMap.Responder("BBB", 1, res)
	assert.NotNil(t, err)

	//Probamos con asignatura incorrecta
	err = PreMap.Responder("FFF", 1, res)
	assert.NotNil(t, err)
	
}

func TestCrearAsigPreMap(t *testing.T) {
	PreMap.CrearAsignatura("DDD")
	asignaturas := len(PreMap.Preguntas)

	assert.Equal(t, 3, asignaturas, "Debe haber tres asignaturas")
}
