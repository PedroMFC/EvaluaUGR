package modelsres

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"
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

	// Vemos que no está vacío
	if  !resRepo.Resenias.AsignaturaRegistrada(asignatura) {
		return &errorsres.ErrorResenia{" la asignatura no está registrada"}
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

	// Vemos que no está vacío
	if  !resRepo.Resenias.AsignaturaRegistrada(asignatura) {
		return nil, &errorsres.ErrorResenia{" la asignatura no está registrada"}
	}
	


	return resRepo.Resenias.ObtenerResenias(asignatura), nil
}


//GetResenia nos aporta una reseñas realizadas en una asignatura
func (resRepo *ReseniasRepositorio) GetResenia(asignatura string, id int) (Resenia, error) {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return Resenia{}, err
	}

	// Vemos que no está vacío
	if  !resRepo.Resenias.AsignaturaRegistrada(asignatura) {
		return Resenia{}, &errorsres.ErrorResenia{" la asignatura no está registrada"}
	}

	resenias := resRepo.Resenias.ObtenerResenias(asignatura)

	if id > len(resenias)-1{
		return Resenia{}, &errorsres.ErrorResenia{" la reseña no contiene un identificador válido"}
	}
	

	return resenias[id], nil
}

//GustaResenia aumenta las valoraciones positivas de una resenia
func (resRepo *ReseniasRepositorio) GustaResenia(asignatura string, id int) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	// Vemos que no está vacío
	if  !resRepo.Resenias.AsignaturaRegistrada(asignatura) {
		return &errorsres.ErrorResenia{" la asignatura no está registrada"}
	}


	return resRepo.Resenias.MeGustaResenia(asignatura, id)
}

//NoGustaResenia aumenta las valoraciones negativas de una resenia
func (resRepo *ReseniasRepositorio) NoGustaResenia(asignatura string, id int) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	// Vemos que no está vacío
	if  !resRepo.Resenias.AsignaturaRegistrada(asignatura) {
		return &errorsres.ErrorResenia{" la asignatura no está registrada"}
	}


	return resRepo.Resenias.NoMeGustaResenia(asignatura, id)
}
