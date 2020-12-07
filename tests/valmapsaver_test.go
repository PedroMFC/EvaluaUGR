package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sort"
	//"log"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
)

func TestGuardaValoValMap(t *testing.T) {
	val := new(modelsval.Valoracion)
	val.SetValoracion(4)

	// Guardamos en una donde ya tenemos registro
	ValMap.GuardarValoracion("AAA", val)
	// Vemos que solo tiene una asignatura y que contiene tres valoraciones
	assert.Equal(t, 1, len(ValMap.Valoraciones), "Solo hay una asignatura")
	// Vemos que tiene tres valoraciones
	assert.Equal(t, 3, len(ValMap.Valoraciones["AAA"]), "La asignatura AAA tiene 3 valoraciones")

	//Creamos una valoración nueva
	ValMap.GuardarValoracion("BBB", val)
	//Hacemos las mismas comprobaciones que antes
	assert.Equal(t, 2, len(ValMap.Valoraciones), "Ahora hay dos asignaturas")
	assert.Equal(t, 1, len(ValMap.Valoraciones["BBB"]), "La asignatura BBB tiene 1 valoración")
}

func TestObtenerValoValMap(t *testing.T) {
	// Obtenemos las valoraciones de la primera asignatura
	valoraciones := ValMap.ObtenerValoraciones("AAA")
	// Las ordenamos para ver que las devuelve bien
	sort.Slice(valoraciones, func(i,j int) bool {
		return valoraciones[i].Valoracion < valoraciones[j].Valoracion
	})

	// Vemos que tiene 3 valoraciones
	assert.Equal(t, 3, len(valoraciones), "AAA tiene tres valoraciones")
	// Vemos que son las que hemos introducido
	assert.Equal(t, 2, valoraciones[0].Valoracion, "Primera valoración AAA: es 2")
	assert.Equal(t, 3, valoraciones[1].Valoracion, "Segunda valoración AAA: es 3")
	assert.Equal(t, 4, valoraciones[2].Valoracion, "Tercera valoración AAA: es 4")

	//Repetimos el proceso con BBB
	valoraciones = ValMap.ObtenerValoraciones("BBB")
	// Las ordenamos para ver que las devuelve bien
	sort.Slice(valoraciones, func(i,j int) bool {
		return valoraciones[i].Valoracion < valoraciones[j].Valoracion
	})

	// Vemos que tiene 1 valoración
	assert.Equal(t, 1, len(valoraciones), "BBB tiene una valoración")
	// Vemos que son las que hemos introducido
	assert.Equal(t, 4, valoraciones[0].Valoracion, "Primera valoración BBB: es 4")

	//Si no hay ninguna debe devolver un array vacío
	valoraciones = ValMap.ObtenerValoraciones("DFGF")
	assert.Equal(t, 0, len(valoraciones), "DFGF no tiene valoraciones")

}

func TestObtenerAsigValMap(t *testing.T) {
	// Obtenemos las valoraciones de la primera asignatura
	asignaturas := ValMap.ObtenerAsignaturas()
	// Las ordenamos para ver que las devuelve bien
	sort.Slice(asignaturas, func(i,j int) bool {
		return asignaturas[i] < asignaturas[j]
	})

	// Vemos que tiene 2 y que son las que hemos indicado
	assert.Equal(t, 2, len(asignaturas), "Solo tenemos dos asignaturas")
	// Vemos que son las que hemos introducido
	assert.Equal(t, "AAA", asignaturas[0], "Primera asignatura es AAA")
	assert.Equal(t, "BBB", asignaturas[1], "Segunda asignatura es BBB")

}