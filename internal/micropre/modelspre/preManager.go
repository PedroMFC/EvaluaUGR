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
func (pDB *PreguntasBD) ConnectDatabase() {

}
