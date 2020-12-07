package modelsval

// &> mockery -dir internal/microval/modelsval -name IValSaver

//IValSaver es la interfaz que deben seguir las estrucuturas de datos que almacenan valoraciones
type IValSaver interface {
	GuardarValoracion(asignatura string, val *Valoracion)
	ObtenerValoraciones(asignatura string) []Valoracion
	ObtenerAsignaturas() []string
}
