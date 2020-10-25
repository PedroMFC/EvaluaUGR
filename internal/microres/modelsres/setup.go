package models

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)


//ReseniasBD es la encargada de realizar las operaciones sobre la base de datos de reseñas
type ReseniasBD struct{ 
	DB *gorm.DB
} 

var baseDatosResenias ReseniasBD

//ConnectDatabase sirve para conectarse a la base de datos de valoraciones
func (vDB *ReseniasBD) ConnectDatabase() {
	
}