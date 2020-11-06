package modelspre

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	//"fmt"
	//"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"
)

//Contiene las reguntas realizadas
type PreguntasRepositorio struct {
	Preguntas map[string][]Pregunta 
}

