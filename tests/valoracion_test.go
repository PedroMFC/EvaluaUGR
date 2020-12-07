package tests

import (
	"github.com/PedroMFC/EvaluaUGR/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"testing"
	"sort"
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
	//Definimos el comportamiento que queremosa que tenga
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("GuardarValoracion", mock.Anything, mock.Anything).Return(nil)

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
	//Definimos el comportamiento que queremos
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("ObtenerValoraciones", "AAA").Return([]modelsval.Valoracion{ modelsval.Valoracion{2}, modelsval.Valoracion{3} })
	ValMapMock.On("ObtenerValoraciones", "BBB").Return([]modelsval.Valoracion{})

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
	//Definimos el comportamiento que queremos
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("ObtenerValoraciones", "AAA").Return([]modelsval.Valoracion{ modelsval.Valoracion{2}, modelsval.Valoracion{3} })

	media, err := ValRepo.GetMedia("AAAAAA")
	assert.NotNil(t, err)
	media, err = ValRepo.GetMedia("AAA")
	assert.Nil(t, err)
	//Comprobamos que son iguales a las que hemos inicializado
	assert.Equal(t, 2.5, media, "La media es 2.5")
}

//Comprobar que calcula la peor valorada correctamente
//Comprobar que calcula la peor valorada correctamente
func TestPeorValoradaVacio(t *testing.T){
	//Definimos el comportamiento que queremos
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("ObtenerAsignaturas").Return([]string{})

	menosValoradas := ValRepo.GetPeorValorada()
	assert.Equal(t, 0, len(menosValoradas), "El array de valoraciones tiene que estar vacío")
}

func TestPeorValoradaContenido(t *testing.T){
	//Definimos el comportamiento que queremos
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("ObtenerAsignaturas").Return([]string{"ABC", "AAA"})
	ValMapMock.On("ObtenerValoraciones", "AAA").Return([]modelsval.Valoracion{ modelsval.Valoracion{2}, modelsval.Valoracion{3} })
	ValMapMock.On("ObtenerValoraciones", "ABC").Return([]modelsval.Valoracion{ modelsval.Valoracion{3}, modelsval.Valoracion{1} })

	menosValoradas := ValRepo.GetPeorValorada()
	assert.Equal(t, 1, len(menosValoradas), "El array tiene que tener una asignatura") // Ahora mismo es hay dos asignaturas: map[AAA:[{2} {3}] ABC:[{3} {1}]]
	assert.Equal(t, "ABC", menosValoradas[0], "Esa no es la peor valorada")
}

func TestPeorValoradaContenidoDoble(t *testing.T){
	//Definimos el comportamiento que queremos
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("ObtenerAsignaturas").Return([]string{"ABC", "AAA", "DEF"})
	ValMapMock.On("ObtenerValoraciones", "AAA").Return([]modelsval.Valoracion{ modelsval.Valoracion{2}, modelsval.Valoracion{3} })
	ValMapMock.On("ObtenerValoraciones", "ABC").Return([]modelsval.Valoracion{ modelsval.Valoracion{3}, modelsval.Valoracion{1} })
	ValMapMock.On("ObtenerValoraciones", "DEF").Return([]modelsval.Valoracion{ modelsval.Valoracion{2} })


	menosValoradas := ValRepo.GetPeorValorada()
	sort.Slice(menosValoradas, func(i,j int) bool {
		return menosValoradas[i] < menosValoradas[j]
	})
	assert.Equal(t, 2, len(menosValoradas), "El array tiene que tener dos asignaturas") // Ahora mismo: map[AAA:[{2} {3}] ABC:[{3} {1}] DEF:[{2}]]
	assert.Equal(t, "ABC", menosValoradas[0], "La primera peor valorada es ABC")
	assert.Equal(t, "DEF", menosValoradas[1], "La segunda peor valorada es DEF")
}


//Comprobar que calcula la peor valorada correctamente
func TestMejorValoradaVacio(t *testing.T){
	//Definimos el comportamiento que queremos
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("ObtenerAsignaturas").Return([]string{})

	menosValoradas := ValRepo.GetMejorValorada()
	assert.Equal(t, 0, len(menosValoradas), "El array de valoraciones tiene que estar vacío")
}

func TestMejorValoradaContenido(t *testing.T){
	//Definimos el comportamiento que queremos
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("ObtenerAsignaturas").Return([]string{"ABC", "AAA"})
	ValMapMock.On("ObtenerValoraciones", "AAA").Return([]modelsval.Valoracion{ modelsval.Valoracion{2}, modelsval.Valoracion{3} })
	ValMapMock.On("ObtenerValoraciones", "ABC").Return([]modelsval.Valoracion{ modelsval.Valoracion{3}, modelsval.Valoracion{1} })

	menosValoradas := ValRepo.GetMejorValorada()
	assert.Equal(t, 1, len(menosValoradas), "El array tiene que tener una asignatura") // Ahora mismo es hay dos asignaturas: map[AAA:[{2} {3}] ABC:[{3} {1}]]
	assert.Equal(t, "AAA", menosValoradas[0], "Esa no es la mejor valorada")
}

func TestMejorValoradaContenidoDoble(t *testing.T){
	//Definimos el comportamiento que queremos
	ValMapMock = mocks.IValSaver{} 

	ValMapMock.On("ObtenerAsignaturas").Return([]string{"ABC", "AAA", "DEF"})
	ValMapMock.On("ObtenerValoraciones", "AAA").Return([]modelsval.Valoracion{ modelsval.Valoracion{4}, modelsval.Valoracion{2} })
	ValMapMock.On("ObtenerValoraciones", "ABC").Return([]modelsval.Valoracion{ modelsval.Valoracion{3}, modelsval.Valoracion{1} })
	ValMapMock.On("ObtenerValoraciones", "DEF").Return([]modelsval.Valoracion{ modelsval.Valoracion{3} })


	masValoradas := ValRepo.GetMejorValorada()
	sort.Slice(masValoradas, func(i,j int) bool {
		return masValoradas[i] < masValoradas[j]
	})
	assert.Equal(t, 2, len(masValoradas), "El array tiene que tener dos asignaturas") // Ahora mismo: map[AAA:[{2} {3}] ABC:[{3} {1}] DEF:[{2}]]
	assert.Equal(t, "AAA", masValoradas[0], "La primera peor valorada es ABC")
	assert.Equal(t, "DEF", masValoradas[1], "La segunda peor valorada es DEF")
}