package modelsres

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	//"fmt"
)

//Contiene las reseñas realizadas
type ReseniasRepositorio struct {
	Resenias IResSaver
}

//NewReseniasRepositorio devuelve un repositorio de reseñas
func NewReseniasRepositorio(resSaver IResSaver) *ReseniasRepositorio {
	return &ReseniasRepositorio{Resenias: resSaver}
}

//CrearAsignaura añade una asignatura para poder valorarla
func (resRepo *ReseniasRepositorio) CrearAsignatura(asignatura string) error{
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	resRepo.Resenias.CrearAsignatura(asignatura)

	return nil
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

	res.MeGusta = 0
	res.NoMeGusta = 0
	resRepo.Resenias.GuardarResenia(asignatura, res) 
	return nil
}


//GetResenias nos aporta las reseñas realizadas en una asignatura
func (resRepo *ReseniasRepositorio) GetResenias(asignatura string) ([]Resenia, error) {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return nil, err
	}

	return resRepo.Resenias.ObtenerResenias(asignatura), nil
}

//GustaResenia aumenta las valoraciones positivas de una resenia
func (resRepo *ReseniasRepositorio) GustaResenia(asignatura string, id int) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	return resRepo.Resenias.MeGustaResenia(asignatura, id)
}

//NoGustaResenia aumenta las valoraciones negativas de una resenia
func (resRepo *ReseniasRepositorio) NoGustaResenia(asignatura string, id int) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	return resRepo.Resenias.NoMeGustaResenia(asignatura, id)
}
