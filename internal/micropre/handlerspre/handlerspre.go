package handlerspre

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
)


func CrearAsignatura(repo modelspre.PreguntasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		err := repo.CrearAsignatura(asig)

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return
		} else {
			c.JSON(http.StatusCreated, gin.H{"Location": "preguntas/asignatura/"+asig})
			return
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



func GetPregunta(repo modelspre.PreguntasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		id, err1 := strconv.Atoi( c.Param("id") )

		if err1 != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Identificador no es un entero"})
			return
		}
		pregunta, err := repo.GetPregunta(asig, id)

		if err != nil{
			if msg := err.Error(); (msg == "Algo salió mal con la pregunta:  la asignatura no está registrada" ||
				msg == "Algo salió mal con la pregunta:  la pregunta no contiene un identificador válido"){

				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			} 
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		} else {

			c.JSON(http.StatusOK, gin.H{"pregunta": pregunta})
		} 
	}
}

type CreatePreguntaInput struct {
	Pregunta string `json:"pregunta" binding:"required"`
}

func Preguntar(repo modelspre.PreguntasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		var input CreatePreguntaInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No se han enviado los campos requeridos"})
			return
		}

		err := repo.Preguntar(asig, input.Pregunta)

		if err != nil{
			if msg := err.Error(); msg == "Algo salió mal con la pregunta:  la asignatura no está registrada" {

				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			} 
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		} else {
			pre, _ := repo.GetPreguntas(asig)
			dir := "preguntas/asignatura/" + asig + "/" + strconv.Itoa(len(pre)-1)
			c.JSON(http.StatusCreated, gin.H{"Location": dir})
		} 
	}
}

type CreateRespuestaInput struct {
	Respuesta string `json:"respuesta" binding:"required"`
}

func Responder(repo modelspre.PreguntasRepositorio) gin.HandlerFunc {
	return func(c *gin.Context) {
		asig := c.Param("asig")
		id, err1 := strconv.Atoi( c.Param("id") )

		if err1 != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Identificador no es un entero"})
			return
		}
		var input CreateRespuestaInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No se han enviado los campos requeridos"})
			return
		}

		err := repo.Responder(asig, id, input.Respuesta)


		if err != nil{
			if msg := err.Error(); (msg == "Algo salió mal con la pregunta:  la asignatura no está registrada" ||
				msg == "Algo salió mal con la pregunta:  no hay una pregunta con ese identificador para la asignatura señalada") {

				c.JSON(http.StatusNotFound, gin.H{"error": err })
				return
			} 
			c.JSON(http.StatusBadRequest, gin.H{"error": err })
			return

		} else {

			c.JSON(http.StatusCreated, gin.H{"Mensaje": "Creada correctamente"})
		} 
	}
}