package tests

import (

	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"

	"github.com/PedroMFC/EvaluaUGR/mocks"
	"github.com/stretchr/testify/mock"

)

/*Mockeamos los datos de VALORACIONES*/
var ValMapMock2 mocks.IValSaver

func StartDataVal(repo *modelsval.ValoracionRepositorio){
	ValMapMock2.On("ObtenerValoraciones", "AAA").Return([]modelsval.Valoracion{ modelsval.Valoracion{2}, modelsval.Valoracion{3} })
	ValMapMock2.On("AsignaturaRegistrada", "AAA").Return(true)
	ValMapMock2.On("AsignaturaRegistrada", "BBB").Return(false)
	ValMapMock2.On("AsignaturaRegistrada", "AA").Return(false)
	ValMapMock2.On("CrearAsignatura", mock.Anything)
	ValMapMock2.On("GuardarValoracion", mock.Anything, mock.Anything)
	ValMapMock2.On("ObtenerAsignaturas").Return([]string{"AAA"})
	
	*repo = *modelsval.NewValoracionsRepositorio(&ValMapMock2)
}

/*Mockeamos los datos de RESEÑAS*/
var ResMapMock2  mocks.IResSaver

func StartDataRes(repo *modelsres.ReseniasRepositorio){

	ResMapMock2.On("CrearAsignatura", mock.Anything)
	ResMapMock2.On("ObtenerResenias", "BBB").Return([]modelsres.Resenia{
		modelsres.Resenia{Opinion: "Me ha parecido interesante"},
		modelsres.Resenia{Opinion: "No me ha gustado"}})
	ResMapMock2.On("AsignaturaRegistrada", "AAA").Return(false)
	ResMapMock2.On("AsignaturaRegistrada", "BBB").Return(true)
	ResMapMock2.On("AsignaturaRegistrada", "ABB").Return(false)
	ResMapMock2.On("AsignaturaRegistrada", "TDA").Return(false)
	ResMapMock2.On("GuardarResenia", mock.Anything, mock.Anything)
	ResMapMock2.On("MeGustaResenia", "BBB", 0).Return(nil)
	ResMapMock2.On("NoMeGustaResenia", "BBB", 0).Return(nil)
	ResMapMock2.On("NoMeGustaResenia", "BBB", 7).Return(&errorsres.ErrorResenia{" la reseña no contiene in identificador válido"})

	*repo = *modelsres.NewReseniasRepositorio(&ResMapMock2)
}

/*Mockeamos los datos de PREGUNTAS*/
var PreMapMock2  mocks.IPreSaver

func StartDataPre(repo *modelspre.PreguntasRepositorio){
	PreMapMock2.On("CrearAsignatura", mock.Anything)
	PreMapMock2.On("ObtenerPregunta", "CCC").Return([]modelspre.Pregunta{
		modelspre.Pregunta{Pregunta: "¿Esta es la primera pregunta?", Identificador: 0, Respuestas: []modelspre.Respuesta{}},
		modelspre.Pregunta{Pregunta: "¿Se ha hecho una segunda pregunta?", Identificador: 1, Respuestas: []modelspre.Respuesta{}},
	})
	PreMapMock2.On("AsignaturaRegistrada", "ABB").Return(false)
	PreMapMock2.On("AsignaturaRegistrada", "AAA").Return(false)
	PreMapMock2.On("AsignaturaRegistrada", "CCC").Return(true)
	PreMapMock2.On("GuardarPregunta", mock.Anything, mock.Anything)
	PreMapMock2.On("Responder", "CCC", 0, mock.Anything).Return(nil)
	PreMapMock2.On("Responder", "CCC", 5, mock.Anything).Return(&errorspre.ErrorPregunta{" no hay una pregunta con ese identificador para la asignatura señalada"})


	*repo = *modelspre.NewPreguntasRepositorio(&PreMapMock2)
}