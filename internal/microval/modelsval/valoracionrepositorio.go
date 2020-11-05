package modelsval

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
)

//Contiene las valoraciones realizadas
type ValoracionRepositorio struct {
	Valoraciones map[string][]Valoracion
}

func NewValoracionsRepositorio() *ValoracionRepositorio{
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
	} else{ //Si no, tenemos que crear una nueva
		valRepo.Valoraciones[asignatura] = []Valoracion{*val}
	}
	return nil
}
