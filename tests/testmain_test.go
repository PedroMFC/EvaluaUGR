package tests

import (
	"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
	"testing"
)

var ValRepo modelsval.ValoracionRepositorio

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

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here.

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}
