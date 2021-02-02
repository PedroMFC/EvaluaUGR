package tests

import (
	"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
	"github.com/PedroMFC/EvaluaUGR/mocks"
	"testing"
)

var ValRepo modelsval.ValoracionRepositorio
var ValMap  modelsval.ValoracionMap
var ValDB modelsval.ValoracionDB
var ValMapMock mocks.IValSaver

var ResRepo modelsres.ReseniasRepositorio
var ResMap  modelsres.ReseniasMap
var ResMapMock mocks.IResSaver

var PreRepo modelspre.PreguntasRepositorio
var PreMap  modelspre.PreguntasMap
var PreMapMock mocks.IPreSaver

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}

func setup() {
	// Do something here.

	//Creamos el conjunto de valoraciones a testear
	//Creamos el conjunto de valoraciones a testear
	ValMap = *modelsval.NewValoracionMap()
	val := new(modelsval.Valoracion)
	val2 := new(modelsval.Valoracion)
	val.Valoracion = 2
	val2.Valoracion = 3
	ValMap.Valoraciones["AAA"] = []modelsval.Valoracion{*val, *val2}

	ValRepo = *modelsval.NewValoracionsRepositorio(&ValMapMock)

	//Creamos el conjunto de reseñas a testear
	ResMap = *modelsres.NewReseniasMap()
	res := new(modelsres.Resenia)
	res2 := new(modelsres.Resenia)
	res.Opinion = "Me ha parecido interesante"
	res.MeGusta = 0
	res.NoMeGusta = 0
	res.Identificador = 0
	res2.Opinion = "No me ha gustado"
	res2.MeGusta = 0
	res2.NoMeGusta = 0
	res2.Identificador = 1
	ResMap.Resenias["BBB"] = []modelsres.Resenia{*res, *res2}

	ResRepo = *modelsres.NewReseniasRepositorio(&ResMapMock)

	//Creamos el conjunto de preguntas a testear
	PreMap = *modelspre.NewPreguntasMap()
	pre := new(modelspre.Pregunta)
	pre2 := new(modelspre.Pregunta)
	pre.Pregunta = "¿Esta es la primera pregunta?"
	pre.Identificador = 0
	respu := new(modelspre.Respuesta)
	respu.Respuesta = "Sí, esta es la primera pregunta."
	pre.Respuestas = []modelspre.Respuesta{*respu}
	pre2.Pregunta = "¿Se ha hecho una segunda pregunta?"
	pre2.Identificador = 1

	PreMap.Preguntas["CCC"] = []modelspre.Pregunta{*pre, *pre2}
	PreRepo = *modelspre.NewPreguntasRepositorio(&PreMapMock)
	
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	ValRepo = modelsval.ValoracionRepositorio{} //Eliniminamos la estructura
	ResRepo = modelsres.ReseniasRepositorio{}
	PreRepo = modelspre.PreguntasRepositorio{}

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func createValDBTest(){
	ValDB = *modelsval.NewValoracionDB()

	sqlStatement := `INSERT INTO val_table VALUES ('AAA','{1,2}');`
	_,err := ValDB.DB.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}
}