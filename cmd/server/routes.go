package server

import (
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"
	log "github.com/sirupsen/logrus"
	"github.com/PedroMFC/EvaluaUGR/internal/customlog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/handlersval"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/handlersres"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/handlerspre"

	"github.com/PedroMFC/EvaluaUGR/mocks"
	"github.com/stretchr/testify/mock"


	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"github.com/joho/godotenv"
	"os"
	"time"
)

/*Mockeamos los datos de VALORACIONES*/

var ValRepo modelsval.ValoracionRepositorio
var ValMapMock2 mocks.IValSaver

func StartDataVal(){
	ValMapMock2.On("ObtenerValoraciones", "AAA").Return([]modelsval.Valoracion{ modelsval.Valoracion{2}, modelsval.Valoracion{3} })
	ValMapMock2.On("AsignaturaRegistrada", "AAA").Return(true)
	ValMapMock2.On("AsignaturaRegistrada", "BBB").Return(false)
	ValMapMock2.On("AsignaturaRegistrada", "AA").Return(false)
	ValMapMock2.On("CrearAsignatura", mock.Anything)
	ValMapMock2.On("GuardarValoracion", mock.Anything, mock.Anything)
	ValMapMock2.On("ObtenerAsignaturas").Return([]string{"AAA"})
	
	ValRepo = *modelsval.NewValoracionsRepositorio(&ValMapMock2)
}

/*Mockeamos los datos de RESEÑAS*/
var ResRepo modelsres.ReseniasRepositorio
var ResMapMock2  mocks.IResSaver

func StartDataRes(){

	ResMapMock2.On("CrearAsignatura", mock.Anything)
	ResMapMock2.On("ObtenerResenias", "BBB").Return([]modelsres.Resenia{
		modelsres.Resenia{Opinion: "Me ha parecido interesante"},
		modelsres.Resenia{Opinion: "No me ha gustado"}})
	ResMapMock2.On("AsignaturaRegistrada", "AAA").Return(false)
	ResMapMock2.On("AsignaturaRegistrada", "BBB").Return(true)
	ResMapMock2.On("AsignaturaRegistrada", "ABB").Return(false)
	ResMapMock2.On("AsignaturaRegistrada", "TDA").Return(false)
	ResMapMock2.On("GuardarResenia", mock.Anything, mock.Anything)
	ResMapMock2.On("MeGustaResenia", "BBB", 0).Return(nil)
	ResMapMock2.On("NoMeGustaResenia", "BBB", 0).Return(nil)
	ResMapMock2.On("NoMeGustaResenia", "BBB", 7).Return(&errorsres.ErrorResenia{" la reseña no contiene in identificador válido"})

	ResRepo = *modelsres.NewReseniasRepositorio(&ResMapMock2)
}

/*Mockeamos los datos de PREGUNTAS*/
var PreRepo modelspre.PreguntasRepositorio
var PreMapMock2  mocks.IPreSaver

func StartDataPre(){
	PreMapMock2.On("CrearAsignatura", mock.Anything)
	PreMapMock2.On("ObtenerPregunta", "CCC").Return([]modelspre.Pregunta{
		modelspre.Pregunta{Pregunta: "¿Esta es la primera pregunta?", Identificador: 0, Respuestas: []modelspre.Respuesta{}},
		modelspre.Pregunta{Pregunta: "¿Se ha hecho una segunda pregunta?", Identificador: 1, Respuestas: []modelspre.Respuesta{}},
	})
	PreMapMock2.On("AsignaturaRegistrada", "ABB").Return(false)
	PreMapMock2.On("AsignaturaRegistrada", "AAA").Return(false)
	PreMapMock2.On("AsignaturaRegistrada", "CCC").Return(true)
	PreMapMock2.On("GuardarPregunta", mock.Anything, mock.Anything)
	PreMapMock2.On("Responder", "CCC", 0, mock.Anything).Return(nil)
	PreMapMock2.On("Responder", "CCC", 5, mock.Anything).Return(&errorspre.ErrorPregunta{" no hay una pregunta con ese identificador para la asignatura señalada"})


	PreRepo = *modelspre.NewPreguntasRepositorio(&PreMapMock2)
}

type applicationGin struct {
	Router *gin.Engine
}


// NO usamos esta función por ahora
// La definimos con la configuración distribuida etcd
// pero no llegamos a arrancar el microservicio
func (a *applicationGin) Start(port string) {
	const PortVarName = "PUERTO"

	// Conectamos con el cliente de etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	// Vemos que hayamos podido hacerlo
	if err != nil{
		log.Fatal("No se ha podido conectar con el cliente", err)
	}

	// Si hemos introducido un puerto directamente empezamos a trabajar en ese puerto
	// Almacenamos en etcd
	if port!= ""{
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, PortVarName, port) 

		if err != nil{
			log.Fatal("No se ha podido escribir la clave", err)
		}
	// Si no, leemos desde el archivo .env si existe
	} else{ 
		err := godotenv.Load()

		if err != nil {
			log.Info("No hay archivo .env")
		}

		// Si hay una entrada en el .env
		//Almacenamos en etcd
		port = os.Getenv(PortVarName); 
		log.WithFields(log.Fields{
			"Puerto": port,
		}).Info("Puerto leído desde .env")

		if port != ""{
			log.Info("Usamos la configuración del .env")
			ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
			_, err := cli.Put(ctx, PortVarName, port)

			if err != nil{
				log.Fatal("No se ha podido escribir la clave", err)
			}
		// Si las anteriores han fallado, usamos un puerto por defecto
		//Almacenamos en etcd
		} else{
			log.Info("Usamos la configuración por defecto")
			port = "8080"
			ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
			_, err := cli.Put(ctx, PortVarName, port)

			if err != nil{
				log.Fatal("No se ha podido escribir la clave", err)
			}
		}

	}

	log.WithFields(log.Fields{
		"Puerto": port,
	}).Info("GIN se ha conectado a un puerto")
	log.Fatal(http.ListenAndServe("localhost:" + port, a.Router))
}


func NewAppGin() *applicationGin {
	//router := gin.Default()
	router := gin.New()

	//Log
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		EnvironmentOverrideColors: true,
		TimestampFormat : "2006-01-02 15:04:05",
	})
	logger.Formatter = &log.TextFormatter{}

	router.Use(customlog.Logger(logger), gin.Recovery())

	//Aquí definimos las rutas
	//Valoraciones
	router.PUT("/valoraciones/asignatura/:asig", handlersval.CrearAsignatura(ValRepo))
	router.GET("/valoraciones/asignatura/:asig", handlersval.GetValoraciones(ValRepo))
	router.POST("/valoraciones/asignatura/:asig/:val", handlersval.Valorar(ValRepo))
	router.GET("/valoraciones/peor", handlersval.GetPeor(ValRepo))
	router.GET("/valoraciones/mejor", handlersval.GetMejor(ValRepo))
	router.GET("/valoraciones/asignatura/:asig/media", handlersval.GetMedia(ValRepo))

	//Reseñas
	router.PUT("/resenias/asignatura/:asig", handlersres.CrearAsignatura(ResRepo))
	router.GET("/resenias/asignatura/:asig", handlersres.GetResenias(ResRepo))
	router.GET("/resenias/asignatura/:asig/:id", handlersres.GetResenia(ResRepo))
	router.POST("/resenias/asignatura/:asig", handlersres.Opinar(ResRepo))
	router.POST("/resenias/asignatura/:asig/:id/gusta", handlersres.GustaResenia(ResRepo))
	router.POST("/resenias/asignatura/:asig/:id/nogusta", handlersres.NoGustaResenia(ResRepo))

	//Preguntas
	router.PUT("/preguntas/asignatura/:asig", handlerspre.CrearAsignatura(PreRepo))
	router.GET("/preguntas/asignatura/:asig", handlerspre.GetPreguntas(PreRepo))
	router.GET("/preguntas/asignatura/:asig/:id", handlerspre.GetPregunta(PreRepo))
	router.POST("/preguntas/asignatura/:asig", handlerspre.Preguntar(PreRepo))
	router.POST("/preguntas/asignatura/:asig/:id", handlerspre.Responder(PreRepo))

	return &applicationGin{Router: router}
}

