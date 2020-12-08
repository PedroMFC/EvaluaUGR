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

func (preMap *PreguntasMap) GuardarPregunta(asignatura string, pre *Pregunta){
	if preMap.Preguntas[asignatura] != nil { // Si ya hay preguntas antes se añaden a las existentes
		pre.Identificador = len(preMap.Preguntas[asignatura])
		preMap.Preguntas[asignatura] = append(preMap.Preguntas[asignatura], *pre)
	} else { //Si no, tenemos que crear una nueva
		pre.Identificador = 0
		preMap.Preguntas[asignatura] = []Pregunta{*pre}
	}
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



