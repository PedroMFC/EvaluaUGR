package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"sort"
	//"log"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
)

func TestGuardaYObtenerResDBMap(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createResDBTest()

	asignaturas := ResDB.ObtenerAsignaturas()

	assert.Equal(t, 1, len(asignaturas), "Solo tenemos una asignatura")
	// Vemos que son las que hemos introducido
	assert.Equal(t, "AAA", asignaturas[0], "Primera asignatura es AAA")

	ResDB.CrearAsignatura("BBB")
	asignaturas = ResDB.ObtenerAsignaturas()
	assert.Equal(t, 2, len(asignaturas), "Ahora tenemos dos")
}

func TestAsigRegistradaResDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createResDBTest()

	resultado := ResDB.AsignaturaRegistrada("AAA")
	assert.Equal(t, true, resultado, "Sí está registrada")

	// Vemos una que NO está registrada
	resultado = ResDB.AsignaturaRegistrada("DDD")
	assert.Equal(t, false, resultado, "No está registrada")

}

func TestObtenerYCrearAsigResDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createValDBTest()

	asignaturas := ResDB.ObtenerAsignaturas()

	assert.Equal(t, 1, len(asignaturas), "Solo tenemos una asignatura")
	// Vemos que son las que hemos introducido
	assert.Equal(t, "AAA", asignaturas[0], "Primera asignatura es AAA")

	ResDB.CrearAsignatura("BBB")
	asignaturas = ResDB.ObtenerAsignaturas()
	assert.Equal(t, 2, len(asignaturas), "Ahora tenemos dos")
}

func TestOpinarYGustaResDB(t *testing.T) {
	if (testing.Short() == true){
		t.Skip()
	}
	createValDBTest()

	var res modelsres.Resenia
	res.Opinion = "Una opinion"
	ResDB.GuardarResenia("AAA", &res)

	err := ResDB.MeGustaResenia("AAA", 0)
	assert.Nil(t, err)

	err = ResDB.NoMeGustaResenia("AAA", 1)
	assert.Nil(t, err)

	err = ResDB.NoMeGustaResenia("AAA", 2)
	assert.NotNil(t, err)

	resenias := ResDB.ObtenerResenias("AAA")


	assert.Equal(t, 1, resenias[0].MeGusta, "Marcada como gusta")
	assert.Equal(t, 0, resenias[0].NoMeGusta, "Marcada como no gusta")
	assert.Equal(t, 0, resenias[1].MeGusta, "Marcada como gusta")
	assert.Equal(t, 1, resenias[1].NoMeGusta, "Marcada como no gusta")

}