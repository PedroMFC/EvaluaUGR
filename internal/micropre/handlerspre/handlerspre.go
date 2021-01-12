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

func GetPreguntas(repo modelspre.PreguntasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		preguntas, err := repo.GetPreguntas(asig)

		if err != nil{
			if err.Error() == "Algo salió mal con la pregunta:  la asignatura no está registrada"{
				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		} else {

			c.JSON(http.StatusOK, gin.H{"preguntas": preguntas})
		} 
	}
}