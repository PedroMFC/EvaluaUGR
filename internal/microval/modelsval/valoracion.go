package modelsval

import (
	"github.com/PedroMFC/EvaluaUGR/internal/microval/errorsval"
)

//Valoracion contiene información acerca de los datos de una asignatura
type Valoracion struct {
	Valoracion int
}

func (val *Valoracion) SetValoracion(num int) error {
	if num < 1 || num > 5 {
		return &errorsval.ErrorValoracion{" la valoración no está en el rango especificado"}
	} 

	val.Valoracion = num

	return nil
}
