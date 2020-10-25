package modelsres

//Resenia contiene información acerca de los datos de una reseña de una asignatura
type Resenia struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Asignatura string `json:"asignatura"`
	Opinion    string `json:"opinion"`
	MeGusta    int    `json:"megusta"`
	NoMeGusta  int    `json:"nomegusta"`
}
