package tests

import (
	"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"
	"github.com/stretchr/testify/mock"
	"github.com/PedroMFC/EvaluaUGR/mocks"
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
	//Definimos el comportamiento que queremos
	ResMapMock = mocks.IResSaver{} 

	ResMapMock.On("MeGustaResenia", "BBB", 0).Return(nil)
	ResMapMock.On("MeGustaResenia", "BBB", 1).Return(&errorsres.ErrorResenia{""})

	err := ResRepo.GustaResenia("ABCDEF", 0) //Asignatura incorrecta
	assert.NotNil(t, err)

	err = ResRepo.GustaResenia("BBB", 0) 
	assert.Nil(t, err)
	err = ResRepo.GustaResenia("BBB", 1) 
	assert.NotNil(t, err)

}

//Comprobamos que las valoraciones negativas funcionan correctamente
func TestNoMeGusta(t *testing.T) {
	//Definimos el comportamiento que queremos
	ResMapMock = mocks.IResSaver{} 

	ResMapMock.On("NoMeGustaResenia", "BBB", 0).Return(nil)
	ResMapMock.On("NoMeGustaResenia", "BBB", 1).Return(&errorsres.ErrorResenia{""})

	err := ResRepo.NoGustaResenia("ABCDEF", 0) //Asignatura incorrecta
	assert.NotNil(t, err)

	err = ResRepo.NoGustaResenia("BBB", 0) 
	assert.Nil(t, err)
	err = ResRepo.NoGustaResenia("BBB", 1) 
	assert.NotNil(t, err)

}

// Comprobar que crea una asignatura correctamente
func TestCrearAsignaturaResenias(t *testing.T) {
	ResMapMock = mocks.IResSaver{} 

	ResMapMock.On("CrearAsignatura", mock.Anything)

	err := ResRepo.CrearAsignatura("AAA")
	assert.Nil(t, err)
	err = ResRepo.CrearAsignatura("AAAAAA")
	assert.NotNil(t, err)

}