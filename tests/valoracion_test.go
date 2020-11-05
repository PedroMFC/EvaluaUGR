package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
)

// Comprobar que NO devuelve error cuando la valoración es correcta
func TestValoracionCorrecta(t *testing.T) {
	val := new(modelsval.Valoracion)
	err := val.SetValoracion(3)
	assert.Nil(t, err)
}

// Comprobar que Sí devuelve error cuando la valoración es superior
func TestValoracionSuperior(t *testing.T) {
	val := new(modelsval.Valoracion)
	err := val.SetValoracion(7)
	assert.NotNil(t, err)
}

// Comprobar que Sí devuelve error cuando la valoración es inferior
func TestValoracionInferior(t *testing.T) {
	val := new(modelsval.Valoracion)
	err := val.SetValoracion(0)
	assert.NotNil(t, err)
}

// Comprobar los errores que devuelve al valorar una asignatura
func TestValorar(t *testing.T) {
	valRepo := modelsval.NewValoracionsRepositorio()
	err := valRepo.Valorar("ABC", 3)
	assert.Nil(t, err)
	err = valRepo.Valorar("ABC", 1)
	assert.Nil(t, err)
	err = valRepo.Valorar("ABCDEF", 1)
	assert.NotNil(t, err)
	err = valRepo.Valorar("ABC", 8)
	assert.NotNil(t, err)
}
