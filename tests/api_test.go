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