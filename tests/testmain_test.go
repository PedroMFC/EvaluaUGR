package tests

import (
	"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
	"testing"
)

var ValRepo modelsval.ValoracionRepositorio
var ValMap  modelsval.ValoracionMap

var ResRepo modelsres.ReseniasRepositorio
var PreRepo modelspre.PreguntasRepositorio

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

	//Creamos el conjunto de reseñas a testear
	ResRepo = *modelsres.NewReseniasRepositorio()
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
	ResRepo.Resenias["BBB"] = []modelsres.Resenia{*res, *res2}

	//Creamos el conjunto de preguntas a testear
	PreRepo = *modelspre.NewPreguntasRepositorio()
	pre := new(modelspre.Pregunta)
	pre2 := new(modelspre.Pregunta)
	pre.Pregunta = "¿Esta es la primera pregunta?"
	pre.Identificador = 0
	respu := new(modelspre.Respuesta)
	respu.Respuesta = "Sí, esta es la primera pregunta."
	pre.Respuestas = []modelspre.Respuesta{*respu}
	pre2.Pregunta = "¿Se ha hecho una segunda pregunta?"
	pre2.Identificador = 1
	PreRepo.Preguntas["CCC"] = []modelspre.Pregunta{*pre, *pre2}

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	ValRepo = modelsval.ValoracionRepositorio{} //Eliniminamos la estructura
	ResRepo = modelsres.ReseniasRepositorio{}
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}
