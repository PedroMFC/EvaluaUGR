package modelspre

import (
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"
)

type PreguntasMap struct {
	Preguntas map[string][]Pregunta 
}

func NewPreguntasMap() *PreguntasMap {
	return &PreguntasMap{Preguntas: make(map[string][]Pregunta)}
}

//CrearAsignatura crea una entrada para una asignatura
func (preMap *PreguntasMap) CrearAsignatura(asignatura string) {
	preMap.Preguntas[asignatura] = []Pregunta{}
}

//AsignaturaRegistrada comprueba si una asignatura está registrada
func (preMap *PreguntasMap) AsignaturaRegistrada(asignatura string) bool {
	return preMap.Preguntas[asignatura] != nil
}

func (preMap *PreguntasMap) GuardarPregunta(asignatura string, pre *Pregunta){
	pre.Identificador = len(preMap.Preguntas[asignatura])
	preMap.Preguntas[asignatura] = append(preMap.Preguntas[asignatura], *pre)
}

func (preMap *PreguntasMap) ObtenerPregunta(asignatura string) []Pregunta {
	return preMap.Preguntas[asignatura]
}

func (preMap *PreguntasMap) Responder(asignatura string, id int, res *Respuesta) error {
	if id > len(preMap.Preguntas[asignatura]) -1 {
		return &errorspre.ErrorPregunta{" no hay una pregunta con ese identificador para la asignatura señalada"}
	}


	if len(preMap.Preguntas[asignatura][id].Respuestas) > 0{
		preMap.Preguntas[asignatura][id].Respuestas = append(preMap.Preguntas[asignatura][id].Respuestas, *res)
	} else{
		preMap.Preguntas[asignatura][id].Respuestas = []Respuesta{*res}
	}
	
	return nil
}



