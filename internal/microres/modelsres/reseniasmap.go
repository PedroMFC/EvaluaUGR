package modelsres

import (
	"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"
)

type ReseniasMap struct {
	Resenias map[string][]Resenia
}

func NewReseniasMap() *ReseniasMap {
	return &ReseniasMap{Resenias: make(map[string][]Resenia)}
}

func (resMap *ReseniasMap) GuardarResenia(asignatura string, opinion *Resenia){
	if resMap.Resenias[asignatura] != nil { // Si ya hay reeseñas antes se añaden a las existentes
		opinion.Identificador = len(resMap.Resenias[asignatura])
		resMap.Resenias[asignatura] = append(resMap.Resenias[asignatura], *opinion)
	} else { //Si no, tenemos que crear una nueva
		opinion.Identificador = 0
		resMap.Resenias[asignatura] = []Resenia{*opinion}
	}
}

func (resMap *ReseniasMap) ObtenerResenias(asignatura string) []Resenia{
	
	return resMap.Resenias[asignatura]
}

func (resMap *ReseniasMap) MeGustaResenia(asignatura string, id int) error {
	if id > len(resMap.Resenias[asignatura]) -1 {
		return &errorsres.ErrorResenia{" la reseña no contiene in identificador válido"}
	}

	resMap.Resenias[asignatura][id].MeGusta++ 

	return nil
}

func (resMap *ReseniasMap) NoMeGustaResenia(asignatura string, id int) error {
	if id > len(resMap.Resenias[asignatura]) -1 {
		return &errorsres.ErrorResenia{" la reseña no contiene in identificador válido"}
	}

	resMap.Resenias[asignatura][id].NoMeGusta++ 

	return nil
}

