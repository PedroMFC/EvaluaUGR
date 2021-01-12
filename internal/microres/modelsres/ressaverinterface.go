package modelsres

// $> mockery -dir internal/microres/modelsres -name IResSaver

//IResSaver es la interfaz para las estructuras de datos de las reseñas
type IResSaver interface {
	GuardarResenia(asignatura string, opinion *Resenia)
	ObtenerResenias(asignatura string) []Resenia
	MeGustaResenia(asignatura string, id int) error
	NoMeGustaResenia(asignatura string, id int) error
	CrearAsignatura(asignatura string)
}
