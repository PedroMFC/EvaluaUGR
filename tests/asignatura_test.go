package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
)

// Comprobar que NO devuelve error cuando el nombre de la asignatura es correcto
func TestAsignaturaCorrecta(t *testing.T) {
	err := asig.AsignaturaCorrecta("ABD")
	assert.Nil(t, err)

}

// Comprobar que SÍ devuelve error cuando el nombre de la asignatura está vacío
func TestAsignaturaNula(t *testing.T) {
	err := asig.AsignaturaCorrecta("")
	assert.NotNil(t, err)
}

// Comprobar que SÍ devuelve error cuando el nombre de la asignatura es demasiado largo
func TestAsignaturaMayor(t *testing.T) {
	err := asig.AsignaturaCorrecta("ABDCDEF")
	assert.NotNil(t, err)
}
