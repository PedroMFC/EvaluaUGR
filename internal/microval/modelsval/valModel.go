package models

//Valoracion contiene información acerca de los datos de una asignatura
type Valoracion struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Asignatura  string `json:"asignatura"`
	Valoracion int `json:"valoracion"`
}

//IManejoValoraciones se encarga de implementar todas las operaciones sobre la base de datos con las valoraciones
type IManejoValoraciones interface{
	valorarAsigantura(asignatura string, opinion int)
	verValoraciones(asignatura string)
	verValoracionMedia(asignaura string)
}
