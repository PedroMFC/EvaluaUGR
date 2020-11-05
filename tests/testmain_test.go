package tests

import (
	"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
	"testing"
)

var ValRepo modelsval.ValoracionRepositorio
var ResRepo modelsres.ReseniasRepositorio

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}

func setup() {
	// Do something here.

	//Creamos el conjunto de valoraciones a testear
	ValRepo = *modelsval.NewValoracionsRepositorio()
	val := new(modelsval.Valoracion)
	val2 := new(modelsval.Valoracion)
	val.Valoracion = 2
	val2.Valoracion = 3
	ValRepo.Valoraciones["AAA"] = []modelsval.Valoracion{*val, *val2}

	//Creamos el conjunto de reseñas a testear
	ResRepo = *modelsres.NewReseniasRepositorio()
	res := new(modelsres.Resenia)
	res2 := new(modelsres.Resenia)
	res.Opinion = "Me ha parecido interesante"
	res2.Opinion = "No me ha gustado"
	ResRepo.Resenias["BBB"] = []modelsres.Resenia{*res, *res2}

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	ValRepo = modelsval.ValoracionRepositorio{} //Eliniminamos la estructura
	ResRepo = modelsres.ReseniasRepositorio{}
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}
