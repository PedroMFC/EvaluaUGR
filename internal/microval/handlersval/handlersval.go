package handlersval

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
)

func CrearAsignatura(repo modelsval.ValoracionRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		err := repo.CrearAsignatura(asig)

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
		} else {
			c.JSON(http.StatusCreated, gin.H{"Location": "valoraciones/asignatura/"+asig})
		} 
	}
}