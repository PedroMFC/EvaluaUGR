package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"sort"
	//"log"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
)

func TestGuardaYObtenerValoDBMap(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createValDBTest()

	val := new(modelsval.Valoracion)
	val.SetValoracion(4)

	// Guardamos en una donde ya tenemos registro
	ValDB.GuardarValoracion("AAA", val)

	valoraciones := ValDB.ObtenerValoraciones("AAA")

	// Vemos que solo tiene una asignatura y que contiene tres valoraciones
	assert.Equal(t, 3, len(valoraciones), "Tenemos tres valoraciones")
}

func TestAsigRegistradaValDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createValDBTest()

	resultado := ValDB.AsignaturaRegistrada("AAA")
	assert.Equal(t, true, resultado, "Sí está registrada")

	// Vemos una que NO está registrada
	resultado = ValDB.AsignaturaRegistrada("DDD")
	assert.Equal(t, false, resultado, "No está registrada")

}

func TestObtenerYCrearAsigValDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createValDBTest()

	asignaturas := ValDB.ObtenerAsignaturas()

	assert.Equal(t, 1, len(asignaturas), "Solo tenemos una asignatura")
	// Vemos que son las que hemos introducido
	assert.Equal(t, "AAA", asignaturas[0], "Primera asignatura es AAA")

	ValDB.CrearAsignatura("BBB")
	asignaturas = ValDB.ObtenerAsignaturas()
	assert.Equal(t, 2, len(asignaturas), "Ahora tenemos dos")
}

