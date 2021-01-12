package modelspre

import (
	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"
	//"fmt"
)

//Contiene las reguntas realizadas
type PreguntasRepositorio struct {
	Preguntas IPreSaver
}

func NewPreguntasRepositorio(preSaver IPreSaver) *PreguntasRepositorio {
	return &PreguntasRepositorio{Preguntas: preSaver}
}

//CrearAsignaura añade una asignatura para poder valorarla
func (preRepo *PreguntasRepositorio) CrearAsignatura(asignatura string) error{
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	preRepo.Preguntas.CrearAsignatura(asignatura)

	return nil
}

//Preguntar añade una pregunta al repositorio
func (preRepo *PreguntasRepositorio) Preguntar(asignatura string, pregunta string) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	pre := new(Pregunta)
	err = pre.SetPregunta(pregunta)
	if err != nil {
		return err
	}

	// Vemos que no está vacío
	if  !preRepo.Preguntas.AsignaturaRegistrada(asignatura) {
		return &errorspre.ErrorPregunta{" la asignatura no está registrada"}
	}


	preRepo.Preguntas.GuardarPregunta(asignatura, pre)

	return nil
}

//GetPreguntas nos aporta las preguntas realizadas en una asignatura
func (preRepo *PreguntasRepositorio) GetPreguntas(asignatura string) ([]Pregunta, error) {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return nil, err
	}

	// Vemos que no está vacío
	if  !preRepo.Preguntas.AsignaturaRegistrada(asignatura) {
		return nil, &errorspre.ErrorPregunta{" la asignatura no está registrada"}
	}

	return preRepo.Preguntas.ObtenerPregunta(asignatura), nil
}

//Responder añade una respuesta al repositorio
func (preRepo *PreguntasRepositorio) Responder(asignatura string, id int, respuesta string) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	res := new(Respuesta)
	err = res.SetRespuesta(respuesta)
	if err != nil {
		return err
	}

	// Vemos que no está vacío
	if  !preRepo.Preguntas.AsignaturaRegistrada(asignatura) {
		return &errorspre.ErrorPregunta{" la asignatura no está registrada"}
	}

	return preRepo.Preguntas.Responder(asignatura, id, res)
}