package modelsres

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	//"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"
)

//Contiene las reseñas realizadas
type ReseniasRepositorio struct {
	Resenias IResSaver
}

//NewReseniasRepositorio devuelve un repositorio de reseñas
func NewReseniasRepositorio(resSaver IResSaver) *ReseniasRepositorio {
	return &ReseniasRepositorio{Resenias: resSaver}
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

	return resRepo.Resenias[asignatura], nil
}

//GustaResenia aumenta las valoraciones positivas de una resenia
func (resRepo *ReseniasRepositorio) GustaResenia(asignatura string, id int) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	if id > len(resRepo.Resenias[asignatura]) -1 {
		return &errorsres.ErrorResenia{" la reseña no contiene in identificador válido"}
	}

	resRepo.Resenias[asignatura][id].MeGusta++ 
	return nil
}

//NoGustaResenia aumenta las valoraciones negativas de una resenia
func (resRepo *ReseniasRepositorio) NoGustaResenia(asignatura string, id int) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	if id > len(resRepo.Resenias[asignatura]) -1 {
		return &errorsres.ErrorResenia{" la reseña no contiene in identificador válido"}
	}

	resRepo.Resenias[asignatura][id].NoMeGusta++ 
	return nil
}