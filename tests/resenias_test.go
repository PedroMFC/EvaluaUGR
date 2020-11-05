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
