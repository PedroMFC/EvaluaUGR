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

	server.StartDataVal()

	handler := server.NewAppGin().Router 
	apitest.New().
		Handler(handler).
		Put("/valoraciones/asignatura/AAA").
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(handler).
		Put("/valoraciones/asignatura/AAAAAA").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiGetValoraciones(t *testing.T) {
	server.StartDataVal()
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Get("/valoraciones/asignatura/AAA").
		Expect(t).
		Status(http.StatusOK).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/valoraciones/asignatura/AA").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Es incorrecta
	apitest.New().
		Handler(handler).
		Get("/valoraciones/asignatura/AAAAAAA").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiPostValoraciones(t *testing.T) {
	server.StartDataVal()
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Post("/valoraciones/asignatura/AAA/5").
		Expect(t).
		Status(http.StatusCreated).
		End()

	// No existe 
	apitest.New().
		Handler(handler).
		Post("/valoraciones/asignatura/BBB/5").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// La valoración no es correcta
	apitest.New().
		Handler(handler).
		Post("/valoraciones/asignatura/AAA/55").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// La asignatura no es correcta
	apitest.New().
		Handler(handler).
		Post("/valoraciones/asignatura/AAAAAAA/55").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestApiPeorValoraciones(t *testing.T) {
	server.StartDataVal()
	handler := server.NewAppGin().Router 

	// Accede a la peor
	apitest.New().
		Handler(handler).
		Get("/valoraciones/peor").
		Expect(t).
		Body(`{"Peores valoradas":["AAA"]}`).
		Status(http.StatusOK).
		End()

}

func TestApiMejorValoraciones(t *testing.T) {
	server.StartDataVal()
	handler := server.NewAppGin().Router 

	// Accede a la mejor
	apitest.New().
		Handler(handler).
		Get("/valoraciones/mejor").
		Expect(t).
		Body(`{"Mejores valoradas":["AAA"]}`).
		Status(http.StatusOK).
		End()

}

func TestApiGetMediaValoraciones(t *testing.T) {
	server.StartDataVal()
	handler := server.NewAppGin().Router 

	// Es correcta
	apitest.New().
		Handler(handler).
		Get("/valoraciones/asignatura/AAA/media").
		Expect(t).
		Status(http.StatusOK).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/valoraciones/asignatura/AA/media").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Es incorrecta
	apitest.New().
		Handler(handler).
		Get("/valoraciones/asignatura/AAAAAAA/media").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}