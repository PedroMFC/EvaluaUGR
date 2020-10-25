package modelspre

//Pregunta contiene los datos de una pregunta formulada
type Pregunta struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Asignatura string `json:"asignatura"`
	Pregunta   string `json:"pregunta"`
}

//Respuesta contine los datos de las respuestas
type Respuesta struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	PreguntaID int    `json:"preguntaid"`
	Respuesta  string `json:"respuesta"`
}
