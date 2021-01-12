package handlerspre

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"strconv"

	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
)


func CrearAsignatura(repo modelspre.PreguntasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		err := repo.CrearAsignatura(asig)

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
		} else {
			c.JSON(http.StatusCreated, gin.H{"Location": "preguntas/asignatura/"+asig})
		} 
	}
}