package handlersres

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
)

func CrearAsignatura(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		err := repo.CrearAsignatura(asig)

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
		} else {
			c.JSON(http.StatusCreated, gin.H{"Location": "resenias/asignatura/"+asig})
		} 
	}
}

func GetResenias(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		resenias, err := repo.GetResenias(asig)


		if err != nil{
			if err.Error() == "Algo salió mal con la reseña:  la asignatura no está registrada"{
				c.JSON(http.StatusNotFound, gin.H{"error": err })
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err })

		} else {
			c.JSON(http.StatusOK, gin.H{"resenias": resenias})
		} 
	}
}

func GetResenia(repo modelsres.ReseniasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		id, err1 := strconv.Atoi( c.Param("id") )

		if err1 != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Identificador no es un entero"})
			return
		}
		resenia, err := repo.GetResenia(asig, id)


		if err != nil{
			if msg := err.Error(); (msg == "Algo salió mal con la reseña:  la asignatura no está registrada" ||
				msg == "Algo salió mal con la reseña:  la reseña no contiene un identificador válido"){

				c.JSON(http.StatusNotFound, gin.H{"error": err })
			} 
			c.JSON(http.StatusBadRequest, gin.H{"error": err })

		} else {

			c.JSON(http.StatusOK, gin.H{"resenia": resenia})
		} 
	}
}