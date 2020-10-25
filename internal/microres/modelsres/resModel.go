package modelsres

//Resenia contiene información acerca de los datos de una reseña de una asignatura
type Resenia struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Asignatura string `json:"asignatura"`
	Opinion    string `json:"opinion"`
	MeGusta    int    `json:"megusta"`
	NoMeGusta  int    `json:"nomegusta"`
}

//IManejoResenias se encarga de implementar todas las operaciones sobre la base de datos con las reseñas
type IManejoResenias interface {
	GuardarResenia(asignatura string, resenia int)
	VerResenias(asignatura string)
	IndicarMeGusta(resenia int)
	IndicarNoMeGusta(resenia int)
}
