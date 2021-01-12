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
		Put("/valoraciones/AAA").
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(handler).
		Put("/valoraciones/AAAAAA").
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
		Get("/valoraciones/AAA").
		Expect(t).
		Status(http.StatusOK).
		End()

	// Es correcta pero no tiene registro
	apitest.New().
		Handler(handler).
		Get("/valoraciones/AA").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// Es incorrecta
	apitest.New().
		Handler(handler).
		Get("/valoraciones/AAAAAAA").
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
		Post("/valoraciones/AAA/5").
		Expect(t).
		Status(http.StatusCreated).
		End()

	// No existe 
	apitest.New().
		Handler(handler).
		Post("/valoraciones/BBB/5").
		Expect(t).
		Status(http.StatusNotFound).
		End()

	// La valoración no es correcta
	apitest.New().
		Handler(handler).
		Post("/valoraciones/AAA/55").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// La asignatura no es correcta
	apitest.New().
		Handler(handler).
		Post("/valoraciones/AAAAAAA/55").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}