package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"sort"
	//"log"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
)

func TestGuardaYObtenerPreDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createPreDBTest()

	asignaturas := PreDB.ObtenerAsignaturas()

	assert.Equal(t, 1, len(asignaturas), "Solo tenemos una asignatura")
	// Vemos que son las que hemos introducido
	assert.Equal(t, "AAA", asignaturas[0], "Primera asignatura es AAA")

	PreDB.CrearAsignatura("BBB")
	asignaturas = PreDB.ObtenerAsignaturas()
	assert.Equal(t, 2, len(asignaturas), "Ahora tenemos dos")
}

func TestAsigRegistradaPreDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createPreDBTest()

	resultado := PreDB.AsignaturaRegistrada("AAA")
	assert.Equal(t, true, resultado, "Sí está registrada")

	// Vemos una que NO está registrada
	resultado = PreDB.AsignaturaRegistrada("DDD")
	assert.Equal(t, false, resultado, "No está registrada")
}


func TestObtenerYCrearAsigPreDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createPreDBTest()

	asignaturas := PreDB.ObtenerAsignaturas()

	assert.Equal(t, 1, len(asignaturas), "Solo tenemos una asignatura")
	// Vemos que son las que hemos introducido
	assert.Equal(t, "AAA", asignaturas[0], "Primera asignatura es AAA")

	PreDB.CrearAsignatura("BBB")
	asignaturas = PreDB.ObtenerAsignaturas()
	assert.Equal(t, 2, len(asignaturas), "Ahora tenemos dos")
}

func TestPreguntaryResponderPreDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createPreDBTest()

	var pre modelspre.Pregunta
	pre.Pregunta = "Añado una pregunta"
	PreDB.GuardarPregunta("AAA", &pre)

	var res modelspre.Respuesta
	res.Respuesta = "Esto la respuesta"
	err := PreDB.Responder("AAA", 0, &res)
	assert.Nil(t, err)
	err = PreDB.Responder("AAA", 8, &res)
	assert.NotNil(t, err)

	preguntas := PreDB.ObtenerPregunta("AAA")
	assert.Equal(t, 2, len(preguntas[0].Respuestas), "tenemos dos respuestas")
}
