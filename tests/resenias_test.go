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
	//Definimos el comportamiento que queremos
	ResMapMock = mocks.IResSaver{} 

	ResMapMock.On("GuardarResenia", mock.Anything, mock.Anything).Return(nil)

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
	//Definimos el comportamiento que queremos
	ResMapMock = mocks.IResSaver{} 

	ResMapMock.On("ObtenerResenias", "BBB").Return([]modelsres.Resenia{ 
		modelsres.Resenia{Opinion: "Me ha parecido interesante"}, 
		modelsres.Resenia{Opinion: "No me ha gustado"} })

	ResMapMock.On("ObtenerResenias", "AAA").Return([]modelsres.Resenia{})

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

//Comprobamos que las valoraciones positivas funcionan correctamente
func TestMeGusta(t *testing.T) {
	err1 := ResRepo.GustaResenia("ABCDEF", 0) //Identificador incorrecto
	assert.NotNil(t, err1)

	res, err := ResRepo.GetResenias("BBB")
	assert.Nil(t, err)
	assert.Equal(t, 0, res[0].MeGusta, "La primera opinión no se ha valorado")
	assert.Equal(t, 0, res[1].MeGusta, "La segunda opinión tampoco se ha valorado")

	err = ResRepo.GustaResenia("BBB", 2) //No tiene tantas reseñas
	assert.NotNil(t, err)

	err = ResRepo.GustaResenia("BBB", 0) //Esta debe de funcionar bien
	assert.Nil(t, err)
	res, _ = ResRepo.GetResenias("BBB")
	assert.Equal(t, 1, res[0].MeGusta, "La primera opinión sí se ha valorado una vez")
	assert.Equal(t, 0, res[1].MeGusta, "La segunda opinión NO se ha valorado")

}

//Comprobamos que las valoraciones negativas funcionan correctamente
func TestNoMeGusta(t *testing.T) {
	err1 := ResRepo.NoGustaResenia("ABCDEF", 0) //Identificador incorrecto
	assert.NotNil(t, err1)

	res, err := ResRepo.GetResenias("BBB")
	assert.Nil(t, err)
	assert.Equal(t, 0, res[0].NoMeGusta, "La primera opinión no se ha valorado")
	assert.Equal(t, 0, res[1].NoMeGusta, "La segunda opinión tampoco se ha valorado")

	err = ResRepo.NoGustaResenia("BBB", 2) //No tiene tantas reseñas
	assert.NotNil(t, err)

	err = ResRepo.NoGustaResenia("BBB", 1) //Esta debe de funcionar bien
	assert.Nil(t, err)
	res, _ = ResRepo.GetResenias("BBB")
	assert.Equal(t, 0, res[0].NoMeGusta, "La primera opinión NO se ha valorado")
	assert.Equal(t, 1, res[1].NoMeGusta, "La segunda opinión SÍ se ha valorado una vez")

}