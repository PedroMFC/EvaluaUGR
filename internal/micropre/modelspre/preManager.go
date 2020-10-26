package modelspre

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

//PreguntasBD es la encargada de realizar las operaciones sobre la base de datos de las preguntas/respuestas
type PreguntasBD struct {
	preguntasDB  *gorm.DB
	respuestasDB *gorm.DB
}

var baseDatos PreguntasBD

//ConnectDatabase sirve para conectarse a la base de datos de las preguntas
func ConnectDatabase(pDB *gorm.DB, rDB *gorm.DB) {

}

//HacerPregunta sirve para almacenar una pregunta de una asignatura
func (pDB *PreguntasBD) HacerPregunta(asignatura string, opinion string) {

}

//VerPreguntas devuelve todas las preguntas de una asignatura
func (pDB *PreguntasBD) VerPreguntas(asignatura string) {

}

//ResponderPregunta almacena una respuesta de una pregunta formulada
func (pDB *PreguntasBD) ResponderPregunta(identificadorPregunta uint, respuesta string) {

}
