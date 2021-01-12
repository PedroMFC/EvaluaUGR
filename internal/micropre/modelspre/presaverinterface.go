package modelspre

// $> mockery -dir internal/micropre/modelspre -name IPreSaver

//IPreSaver define las operaciones sobre la estructura de datos de preguntas
type IPreSaver interface {
	GuardarPregunta(asignatura string, pre *Pregunta)
	ObtenerPregunta(asignatura string) []Pregunta
	Responder(asignatura string, id int, res *Respuesta) error
	CrearAsignatura(asignatura string)
}
