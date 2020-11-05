package modelsval

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	//"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/errorsval"
)

//Contiene las valoraciones realizadas
type ValoracionRepositorio struct {
	Valoraciones map[string][]Valoracion
}

func NewValoracionsRepositorio() *ValoracionRepositorio {
	return &ValoracionRepositorio{Valoraciones: make(map[string][]Valoracion)}
}

//Valorar añade una valoración al repositorio
func (valRepo *ValoracionRepositorio) Valorar(asignatura string, numero int) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	val := new(Valoracion)
	err = val.SetValoracion(numero)
	if err != nil {
		return err
	}

	if valRepo.Valoraciones[asignatura] != nil { // Si ya hay valoraciones antes se añaden a las existentes
		valRepo.Valoraciones[asignatura] = append(valRepo.Valoraciones[asignatura], *val)
	} else { //Si no, tenemos que crear una nueva
		valRepo.Valoraciones[asignatura] = []Valoracion{*val}
	}

	return nil
}

//GetValoraciones nos aporta las valoraciones realizadas en una asignatura
func (valRepo *ValoracionRepositorio) GetValoraciones(asignatura string) ([]Valoracion, error) {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return nil, err
	}

	return valRepo.Valoraciones[asignatura], nil
}

//GetMedia nos aporta la valoración media de una asignatura
func (valRepo *ValoracionRepositorio) GetMedia(asignatura string) (float64, error) {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return 0, err
	}

	valoraciones := valRepo.Valoraciones[asignatura]
	if valoraciones == nil { //Si está vacío
		return 0, &errorsval.ErrorValoracion{" no hay valoraciones disponibles"}
	}

	//Si no está vacío, calculamos la media
	media := 0.0
	for _, val := range valoraciones {
		media = media + float64(val.Valoracion)
	}
	media = media / float64(len(valoraciones))

	return media, nil
}
