package modelsval

import (
	"github.com/PedroMFC/EvaluaUGR/internal/microval/errorsval"
)

//Contiene las valoraciones realizadas
type ValoracionRepositorio struct {
	Valoraones map[string][]Valoracion
}
