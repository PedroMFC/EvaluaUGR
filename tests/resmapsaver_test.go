package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sort"
	//"log"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
)

func TestGuardaReseResMap(t *testing.T) {
	res := new(modelsres.Resenia)
	res.SetOpinion("Una opinión")
	res.MeGusta = 0
	res.NoMeGusta = 0

	//Guardamos en una donde ya tenemos registro
	ResMap.GuardarResenia("BBB", res)
	// Vemos que solo tiene una asignatura y que contiene tres valoraciones
	assert.Equal(t, 1, len(ResMap.Resenias), "Solo hay una asignatura")
	// Vemos que tiene tres valoraciones
	assert.Equal(t, 3, len(ResMap.Resenias["BBB"]), "La asignatura BBB tiene 3 reseñas")

	//Creamos una reseña nueva
	ResMap.GuardarResenia("CCC", res)
	//Hacemos las mismas comprobaciones que antes
	assert.Equal(t, 2, len(ResMap.Resenias), "Ahora hay dos asignaturas")
	assert.Equal(t, 1, len(ResMap.Resenias["CCC"]), "La asignatura CCC tiene 1 valoración")
}

//Comprobamos que se asignan bien los identificadores al enviar la reseña
func TestIdentificadoresCorrectos(t *testing.T) {
	assert.Equal(t, 0, ResMap.Resenias["BBB"][0].Identificador, "Es la primera reseña enviada")
	assert.Equal(t, 1, ResMap.Resenias["BBB"][1].Identificador, "Es la segunda reseña enviada")
	assert.Equal(t, 2, ResMap.Resenias["BBB"][2].Identificador, "Es la tercera reseña enviada")

}

func TestObtenerResenias(t *testing.T) {
	// Obtenemos las valoraciones de la primera asignatura
	resenias := ResMap.ObtenerResenias("BBB")
	// Las ordenamos para ver que las devuelve bien
	sort.Slice(resenias, func(i,j int) bool {
		return resenias[i].Opinion < resenias[j].Opinion
	})

	
	// Vemos que tiene 3 resenias
	assert.Equal(t, 3, len(resenias), "BBB tiene tres reseñas")
	// Vemos que son las que hemos introducido
	assert.Equal(t, "Me ha parecido interesante", resenias[0].Opinion, "Primera reseña de BB:")
	assert.Equal(t, "No me ha gustado", resenias[1].Opinion, "Segunda reseña de BB")
	assert.Equal(t, "Una opinión", resenias[2].Opinion, "Tercera reseña de BBB")

	//Repetimos el proceso con CCC
	resenias = ResMap.ObtenerResenias("CCC")
	// Las ordenamos para ver que las devuelve bien
	sort.Slice(resenias, func(i,j int) bool {
		return resenias[i].Opinion < resenias[j].Opinion
	})

	// Vemos que tiene 1 valoración
	assert.Equal(t, 1, len(resenias), "CCC tiene una reseña")
	// Vemos que son las que hemos introducido
	assert.Equal(t, "Una opinión", resenias[0].Opinion, "Primera reseña de CCC")

	//Vemos que si no está devuelve uno vacío
	resenias = ResMap.ObtenerResenias("DFGF")
	assert.Equal(t, 0, len(resenias), "DFGF no tiene una reseña")
}

func TestGustaResMap(t *testing.T) {
	//Vemos con identificadores válidos
	err := ResMap.MeGustaResenia("BBB", 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, ResMap.Resenias["BBB"][1].MeGusta , "Tiene un me gusta")
	err = ResMap.MeGustaResenia("BBB", 2)
	assert.Nil(t, err)
	err = ResMap.MeGustaResenia("BBB", 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, ResMap.Resenias["BBB"][2].MeGusta , "Tiene dos me gusta")
	assert.Equal(t, 0, ResMap.Resenias["BBB"][0].MeGusta , "Tiene cero me gusta")

	//Probamos con identificador incorrecto
	err = ResMap.MeGustaResenia("BBB", 5)
	assert.NotNil(t, err)

	//Probamos con asignatura incorrecta
	err = ResMap.MeGustaResenia("FGG", 0)
	assert.NotNil(t, err)
}

func TestNoGustaResMap(t *testing.T) {
	//Vemos con identificadores válidos
	err := ResMap.NoMeGustaResenia("BBB", 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, ResMap.Resenias["BBB"][1].NoMeGusta , "Tiene un no me gusta")
	err = ResMap.NoMeGustaResenia("BBB", 2)
	assert.Nil(t, err)
	err = ResMap.NoMeGustaResenia("BBB", 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, ResMap.Resenias["BBB"][2].NoMeGusta , "Tiene dos no me gusta")
	assert.Equal(t, 0, ResMap.Resenias["BBB"][0].NoMeGusta , "Tiene cero no me gusta")

	//Probamos con identificador incorrecto
	err = ResMap.NoMeGustaResenia("BBB", 5)
	assert.NotNil(t, err)

	//Probamos con asignatura incorrecta
	err = ResMap.NoMeGustaResenia("FGG", 0)
	assert.NotNil(t, err)
}