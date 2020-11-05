package modelsres

import (
	"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"
	"github.com/rivo/uniseg"
)

//Valoracion contiene información acerca de los datos de una asignatura
type Resenia struct {
	Opinion string
	Identificador int
	MeGusta int
	NoMeGusta int
}

//SetOpinion almacena la respuesta de la reseña
func (res *Resenia) SetOpinion(opi string) error {
	numCaracteres := uniseg.GraphemeClusterCount(opi)
	if numCaracteres < 1 || numCaracteres > 300 {
		return &errorsres.ErrorResenia{" la reseña no tiene un número de caracteres adecuados"}
	}

	res.Opinion = opi

	return nil
}
