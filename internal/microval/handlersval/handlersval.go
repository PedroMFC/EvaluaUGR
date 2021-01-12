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

func GetValoraciones(repo modelsval.ValoracionRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		valoraciones, err := repo.GetValoraciones(asig)

		//https://github.com/gin-gonic/gin/issues/1335
		//time.Sleep(20 * time.Second)

		if err != nil{
			if err.Error() == "Algo salió mal en la valoración:  la asignatura no está registrada"{
				c.JSON(http.StatusNotFound, gin.H{"error": err })
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err })

		} else {
			var valoracionesNum []int
			for _, val :=range valoraciones{
				valoracionesNum = append(valoracionesNum, val.Valoracion)
			}

			c.JSON(http.StatusOK, gin.H{"valoraciones": valoracionesNum})
		} 
	}
}