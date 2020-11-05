package modelsres

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	//"fmt"
	//"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"
)

//Contiene las reseñas realizadas
type ReseniasRepositorio struct {
	Resenias map[string][]Resenia //Con la tabla Hash hace la consultas más rápido
}

func NewReseniasRepositorio() *ReseniasRepositorio {
	return &ReseniasRepositorio{Resenias: make(map[string][]Resenia)}
}

//Opinar añade una reseña/opinión al repositorio
func (resRepo *ReseniasRepositorio) Opinar(asignatura string, opinion string) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	res := new(Resenia)
	err = res.SetOpinion(opinion)
	if err != nil {
		return err
	}

	if resRepo.Resenias[asignatura] != nil { // Si ya hay reeseñas antes se añaden a las existentes
		resRepo.Resenias[asignatura] = append(resRepo.Resenias[asignatura], *res)
	} else { //Si no, tenemos que crear una nueva
		resRepo.Resenias[asignatura] = []Resenia{*res}
	}

	return nil
}
