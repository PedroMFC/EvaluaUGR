package tests

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/PedroMFC/EvaluaUGR/cmd/server"
)

/*
VALORACIONES
*/
func TestApiCrearAsigValoraciones(t *testing.T) {

	StartDataVal(&server.ValRepo)

	handler := server.NewAppGin().Router 
	apitest.New().
		Handler(handler).
		Put("/asignaturas/AAA/valoraciones").
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(handler).
		Put("/asignaturas/AAAAAA/valoraciones").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiGetValoraciones(t *testing.T) {
	StartDataVal(&server.ValRepo)
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/AAA/valoraciones").
		Expect(t).
		Status(http.StatusOK).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/asignaturas/AA/valoraciones").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Es incorrecta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/AAAAAAA/valoraciones").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiPostValoraciones(t *testing.T) {
	StartDataVal(&server.ValRepo)
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/AAA/valoraciones/5").
		Expect(t).
		Status(http.StatusCreated).
		End()

	// No existe 
	apitest.New().
		Handler(handler).
		Post("/asignaturas/BBB/valoraciones/5").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// La valoración no es correcta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/AAA/valoraciones/55").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// La asignatura no es correcta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/AAAAAAA/valoraciones/55").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiPeorValoraciones(t *testing.T) {
	StartDataVal(&server.ValRepo)
	handler := server.NewAppGin().Router 

	// Accede a la peor
	apitest.New().
		Handler(handler).
		Get("/asignaturaspeorvalorada").
		Expect(t).
		Body(`{"Peores valoradas":["AAA"]}`).
		Status(http.StatusOK).
		End()

}

func TestApiMejorValoraciones(t *testing.T) {
	StartDataVal(&server.ValRepo)
	handler := server.NewAppGin().Router 

	// Accede a la mejor
	apitest.New().
		Handler(handler).
		Get("/asignaturasmejorvalorada").
		Expect(t).
		Body(`{"Mejores valoradas":["AAA"]}`).
		Status(http.StatusOK).
		End()

}

func TestApiGetMediaValoraciones(t *testing.T) {
	StartDataVal(&server.ValRepo)
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/AAA/valoraciones/media").
		Expect(t).
		Status(http.StatusOK).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/asignaturas/AA/valoraciones/media").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Es incorrecta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/AAAAAAA/valoraciones/media").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

/*
RESEÑAS
*/
func TestApiCrearAsigResenias(t *testing.T) {
	StartDataRes(&server.ResRepo)
	handler := server.NewAppGin().Router 

	apitest.New().
		Handler(handler).
		Put("/asignaturas/AAA/resenias").
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(handler).
		Put("/asignaturas/AAAAAA/resenias").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiGetResenias(t *testing.T) {
	StartDataRes(&server.ResRepo)
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBB/resenias").
		Expect(t).
		Status(http.StatusOK).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/asignaturas/ABB/resenias").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// No es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBBBBBB/resenias").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiGetResenia(t *testing.T) {
	StartDataRes(&server.ResRepo)
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBB/resenias/0").
		Expect(t).
		Status(http.StatusOK).
		End()

	// No tiene registro
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBB/resenias/5").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/asignaturas/ABB/resenias/0").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// No es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBBBBBB/resenias/0").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// No es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBB/resenias/sdfsfd").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	
}

func TestApiOpinarResenias(t *testing.T) {
	StartDataRes(&server.ResRepo)
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/BBB/resenias").
		Body("{ \"opinion\": \"Muy bonita\" }").
		Expect(t).
		Status(http.StatusCreated).
		End()
	
	// Asignatura no registrada
	apitest.New().
		Handler(handler).
		Post("/asignaturas/TDA/resenias").
		Body("{ \"opinion\": \"Muy bonita\" }").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Petición incorrecta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/BBB/resenias").
		Body("{\"opion\": \"Muy bonita\" }").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Asignatura no válida
	apitest.New().
		Handler(handler).
		Post("/asignaturas/BBBBBB/resenias").
		Body("{ \"opinion\": \"Muy bonita\" }").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiGustaResenias(t *testing.T) {
	StartDataRes(&server.ResRepo)
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/BBB/resenias/0/gusta").
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(handler).
		Post("/asignaturas/BBB/resenias/0/nogusta").
		Expect(t).
		Status(http.StatusCreated).
		End()

	// Petición incorrecta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/BBB/resenias/7/nogusta").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	apitest.New().
		Handler(handler).
		Post("/asignaturas/BBBBBBBB/resenias/0/nogusta").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

/**
PREGUNTAS
**/
func TestApiCrearAsigPreguntas(t *testing.T) {
	StartDataPre(&server.PreRepo)
	handler := server.NewAppGin().Router 

	apitest.New().
		Handler(handler).
		Put("/asignaturas/AAA/preguntas").
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(handler).
		Put("/asignaturas/AAAAAA/preguntas").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiGetPreguntas(t *testing.T) {
	StartDataPre(&server.PreRepo)
	handler := server.NewAppGin().Router

	// Es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/CCC/preguntas").
		Expect(t).
		Status(http.StatusOK).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/asignaturas/ABB/preguntas").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// No es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBBBBBB/preguntas").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiGetPregunta(t *testing.T) {
	StartDataPre(&server.PreRepo)
	handler := server.NewAppGin().Router

	// Es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/CCC/preguntas/0").
		Expect(t).
		Status(http.StatusOK).
		End()

	// No tiene registro
	apitest.New().
		Handler(handler).
		Get("/asignaturas/CCC/preguntas/5").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/asignaturas/ABB/preguntas/0").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// No es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBBBBBB/preguntas/0").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// No es correcta
	apitest.New().
		Handler(handler).
		Get("/asignaturas/BBB/preguntas/sdfsfd").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	
}

func TestApiPreguntarPreguntas(t *testing.T) {
	StartDataPre(&server.PreRepo)
	handler := server.NewAppGin().Router

	// Es correcta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/CCC/preguntas").
		Body("{ \"pregunta\": \"Me recomiendan la asignatura\" }").
		Expect(t).
		Status(http.StatusCreated).
		End()

	// Petición asignatura no registrada
	apitest.New().
		Handler(handler).
		Post("/asignaturas/AAA/preguntas").
		Body("{\"pregunta\": \"Me recomiendan la asignatura\" }").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Asignatura no válida
	apitest.New().
		Handler(handler).
		Post("/asignaturas/AAAAAAAA/preguntas").
		Body("{\"pregunta\": \"Me recomiendan la asignatura\" }").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// Petición incorrecta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/AAAAAAAA/preguntas").
		Body("{\"pnta\": \"Me recomiendan la asignatura\" }").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiResponderPreguntas(t *testing.T) {
	StartDataPre(&server.PreRepo)
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/CCC/preguntas/0").
		Body("{ \"respuesta\": \"Es una respuesta\"}").
		Expect(t).
		Status(http.StatusCreated).
		End()

	// Pregunta no registrada
	apitest.New().
		Handler(handler).
		Post("/asignaturas/CCC/preguntas/5").
		Body("{ \"respuesta\": \"Es una respuesta\"}").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Asignatura no registrada
	apitest.New().
		Handler(handler).
		Post("/asignaturas/AAA/preguntas/0").
		Body("{ \"respuesta\": \"Es una respuesta\"}").
		Expect(t).
		Status(http.StatusNotFound).
		End()


	// Asignatura no válida
	apitest.New().
		Handler(handler).
		Post("/asignaturas/CCCCCCC/preguntas/0").
		Body("{\"respuesta\": \"Es una respuesta\"}").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// petición incorrecta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/CCC/preguntas/0").
		Body("{\"resua\": \"Es una respuesta\"}").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// petición incorrecta
	apitest.New().
		Handler(handler).
		Post("/asignaturas/CCC/preguntas/sdfsdf").
		Body("{\"respuesta\": \"Es una respuesta\"}").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

}