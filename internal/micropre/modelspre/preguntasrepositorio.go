package modelspre

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	//"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"
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

//Responder añade una respuesta al repositorio
func (preRepo *PreguntasRepositorio) Responder(asignatura string, id int, respuesta string) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	if id > len(preRepo.Preguntas[asignatura]) -1 {
		return &errorspre.ErrorPregunta{" no hay una pregunta con ese identificador para la asignatura señalada"}
	}

	res := new(Respuesta)
	err = res.SetRespuesta(respuesta)
	if err != nil {
		return err
	}

	if len(preRepo.Preguntas[asignatura][id].Respuestas) > 0{
		preRepo.Preguntas[asignatura][id].Respuestas = append(preRepo.Preguntas[asignatura][id].Respuestas, *res)
	} else{
		preRepo.Preguntas[asignatura][id].Respuestas = []Respuesta{*res}
	}
	
	return nil
}