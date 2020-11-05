package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	//"fmt"
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
	//valRepo := modelsval.NewValoracionsRepositorio()
	err := ValRepo.Valorar("ABC", 3)
	assert.Nil(t, err)
	err = ValRepo.Valorar("ABC", 1)
	assert.Nil(t, err)
	err = ValRepo.Valorar("ABCDEF", 1)
	assert.NotNil(t, err)
	err = ValRepo.Valorar("ABC", 8)
	assert.NotNil(t, err)
}

// Comprobar que devuelve de manera correcta las valoraciones
func TestGetValoraciones(t *testing.T) {
	val, err := ValRepo.GetValoraciones("AAAAAA")
	assert.NotNil(t, err)
	val, err = ValRepo.GetValoraciones("AAA")
	assert.Nil(t, err)
	//Comprobamos que son iguales a las que hemos inicializado
	assert.Equal(t, 2, val[0].Valoracion, "La primera valoracion es 2")
	assert.Equal(t, 3, val[1].Valoracion, "La segunda valoracion es 3")

	//Comprobamos que si no está la asinatura está vacío
	val, err = ValRepo.GetValoraciones("BBB")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(val), "El array de valoraciones tiene que estar vacío")
}


func TestGetMedia(t *testing.T) {
	media, err := ValRepo.GetMedia("AAAAAA")
	assert.NotNil(t, err)
	media, err = ValRepo.GetMedia("AAA")
	assert.Nil(t, err)
	//Comprobamos que son iguales a las que hemos inicializado
	assert.Equal(t, 2.5, media, "La media es 2.5")
}