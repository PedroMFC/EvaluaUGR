package modelsres

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ReseniasBD es la encargada de realizar las operaciones sobre la base de datos de reseñas
type ReseniasBD struct {
	DB *gorm.DB
}

var baseDatosResenias ReseniasBD

//ConnectDatabase sirve para conectarse a la base de datos de valoraciones
func (vDB *ReseniasBD) ConnectDatabase() {

}

//GuardarResenia almacena una reseña en la base de datos
func (vDB *ReseniasBD) GuardarResenia(asignatura string, resenia int) {

}

//VerResenias devuelve todas las reseñas de una asigntura
func (vDB *ReseniasBD) VerResenias(asignatura string) {

}

//IndicarMeGusta sirve para indicar que una reseña ha sido útil
func (vDB *ReseniasBD) IndicarMeGusta(resenia int) {

}

//IndicarNoMeGusta sirve para indicar que una reseña NO ha sido útil
func (vDB *ReseniasBD) IndicarNoMeGusta(resenia int) {

}
