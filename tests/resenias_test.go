package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
)

// Comprobar que indicar una opinión se hace de manera correcta
func TestOpinionCorrecta(t *testing.T) {
	resenia := new(modelsres.Resenia)
	err := resenia.SetOpinion("Esta es una reseña válida")
	assert.Nil(t, err)
	err = resenia.SetOpinion("")
	assert.NotNil(t, err)
	err = resenia.SetOpinion("Esta no es una reseña válida porque tiene demasiados caracteres." +
		"Esta no es una reseña válida porque tiene demasiados caracteres." +
		"Esta no es una reseña válida porque tiene demasiados caracteres." +
		"Esta no es una reseña válida porque tiene demasiados caracteres." +
		"Esta no es una reseña válida porque tiene demasiados caracteres.")
	assert.NotNil(t, err)
}

// Comprobar los errores que devuelve al opinar sobre una asignatura
func TestOpinar(t *testing.T) {
	err := ResRepo.Opinar("ABC", "Esta reseña es adecuada")
	assert.Nil(t, err)
	err = ResRepo.Opinar("ABCDEF", "Esta reseña es buena también")
	assert.NotNil(t, err)
	err = ResRepo.Opinar("ABC", "Esta no es una reseña válida porque tiene demasiados caracteres." +
								 "Esta no es una reseña válida porque tiene demasiados caracteres." +
								 "Esta no es una reseña válida porque tiene demasiados caracteres." +
								 "Esta no es una reseña válida porque tiene demasiados caracteres." +
								 "Esta no es una reseña válida porque tiene demasiados caracteres.")
	assert.NotNil(t, err)
}

// Comprobar que devuelve de manera correcta las reseñas
func TestGetResenias(t *testing.T) {
	res, err := ResRepo.GetResenias("AAAAAA")
	assert.NotNil(t, err)
	res, err = ResRepo.GetResenias("BBB")
	assert.Nil(t, err)
	//Comprobamos que son iguales a las que hemos inicializado
	assert.Equal(t, "Me ha parecido interesante", res[0].Opinion, "Debe coincidir la primera reseña")
	assert.Equal(t, "No me ha gustado", res[1].Opinion, "Y también debe coincidir la segunda")

	//Comprobamos que si no está la asinatura está vacío
	res, err = ResRepo.GetResenias("AAA")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(res), "El array de reseñas tiene que estar vacío")
}