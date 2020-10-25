package modelsval

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ValoracionBD es la encargada de realizar las operaciones sobre la base de datos
type ValoracionBD struct {
	DB *gorm.DB
}

var baseDatos ValoracionBD

//ConnectDatabase sirve para conectarse a la base de datos de valoraciones
func (vDB *ValoracionBD) ConnectDatabase() {

}

//ValorarAsignatura añade a la base de datos una valoración
func (vDB *ValoracionBD) ValorarAsignatura(asignatura string, opinion int) {

}

//VerValoraciones obtine todas las valoraciones de una asignatura
func (vDB *ValoracionBD) VerValoraciones(asignatura string) {

}

//VerValoracionMedia devuelve la valoración media de una asignatura
func (vDB *ValoracionBD) VerValoracionMedia(asignaura string) {

}
