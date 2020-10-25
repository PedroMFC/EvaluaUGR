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

//ValoracionDB implementa la interfaz IManjeValoraciones
func (vDB *ValoracionBD) valorarAsigantura(asignatura string, opinion int) {

}

func (vDB *ValoracionBD) verValoraciones(asignatura string) {

}

func (vDB *ValoracionBD) verValoracionMedia(asignaura string) {

}
