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

func NewPreguntasRepositorio() *PreguntasRepositorio {
	return &PreguntasRepositorio{Preguntas: make(map[string][]Pregunta)}
}

//Preguntar añade una pregunta al repositorio
func (preRepo *PreguntasRepositorio) Preguntar(asignatura string, pregunta string) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	pre := new(Pregunta)
	err = pre.SetPregunta(pregunta)
	if err != nil {
		return err
	}

	if preRepo.Preguntas[asignatura] != nil { // Si ya hay reeseñas antes se añaden a las existentes
		pre.Identificador = len(preRepo.Preguntas[asignatura])
		preRepo.Preguntas[asignatura] = append(preRepo.Preguntas[asignatura], *pre)
	} else { //Si no, tenemos que crear una nueva
		pre.Identificador = 0
		preRepo.Preguntas[asignatura] = []Pregunta{*pre}
	}

	return nil
}

//GetPreguntas nos aporta las preguntas realizadas en una asignatura
func (preRepo *PreguntasRepositorio) GetPreguntas(asignatura string) ([]Pregunta, error) {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return nil, err
	}

	return preRepo.Preguntas[asignatura], nil
}