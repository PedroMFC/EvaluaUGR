package asig

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/errorsasig"
	"github.com/rivo/uniseg"
)

func AsignaturaCorrecta(nom string) error {
	numCaracteres := uniseg.GraphemeClusterCount(nom)
	if numCaracteres > 0 && numCaracteres < 6 {
		return nil
	} else {
		return &errorsasig.ErrorAsigntaura{"la asignatura no contiene un formato válido"}
	}

}
