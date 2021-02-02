package server

import (
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

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	//"github.com/joho/godotenv"
	"os"
	"time"
)


var ValRepo modelsval.ValoracionRepositorio
var ValDB  modelsval.ValoracionDB
var ResRepo modelsres.ReseniasRepositorio
var ResDB modelsres.ReseniasDB
var PreRepo modelspre.PreguntasRepositorio
var PreDB modelspre.PreguntasDB

func StartData(){
	ValDB = *modelsval.NewValoracionDB()
	ValRepo = *modelsval.NewValoracionsRepositorio(&ValDB)

	ResDB = *modelsres.NewReseniasDB()
	ResRepo = *modelsres.NewReseniasRepositorio(&ResDB)
	
	PreDB = *modelspre.NewPreguntasDB()
	PreRepo = *modelspre.NewPreguntasRepositorio(&PreDB)
}

type applicationGin struct {
	Router *gin.Engine
}


// NO usamos esta función por ahora
// La definimos con la configuración distribuida etcd
// pero no llegamos a arrancar el microservicio
const PortVarName = "PUERTO"
const AddVarName = "ADDRESS"
const PortDefault = "8080"
const AddDefault = "localhost"

func (a *applicationGin) Start() {
	
	port := ""
	add := ""

	// Conectamos con el cliente de etcd
	etcdHost := os.Getenv("ETCD_HOST") 
	etcdPort := os.Getenv("ETCD_PORT") 
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdHost + ":" + etcdPort},
		DialTimeout: 20 * time.Second,
	})

	// Vemos que hayamos podido hacerlo
	if err != nil{
		log.Println("No se ha podido conectar con el cliente", err)
	}

	/* PUERTO */

	// Si hay una entrada en el .env 
	// Almacenamos en etcd
	port = os.Getenv(PortVarName); 
	log.WithFields(log.Fields{
		"Puerto": port,
	}).Info("Puerto leído desde .env ")

	if port != ""{
		log.Info("Usamos la configuración del .env")
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, PortVarName, port)

		if err != nil{
			log.Println("No se ha podido escribir la clave", err)
		}
	// Si falla con el .env, usamos un puerto por defecto
	//Almacenamos en etcd
	} else{
		log.Info("Usamos la configuración por defecto")
		port = PortDefault
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, PortVarName, port)

		if err != nil{
			log.Println("No se ha podido escribir la clave", err)
		}
	}

	/* DIRECCIÓN */
	// Si hay una entrada en el .env 
	// Almacenamos en etcd
	add = os.Getenv(AddVarName); 
	log.WithFields(log.Fields{
		"Dirección": add,
	}).Info("Dirección leída desde .env ")

	if add != ""{
		log.Info("Usamos la dirección del .env")
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, AddVarName, add)

		if err != nil{
			log.Println("No se ha podido escribir la clave", err)
		}
	// Si falla con el .env, usamos un puerto por defecto
	//Almacenamos en etcd
	} else{
		log.Info("Usamos la configuración para la dirección por defecto")
		add = AddDefault
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, AddVarName, add)

		if err != nil{
			log.Println("No se ha podido escribir la clave", err)
		}
	}


	log.WithFields(log.Fields{
		"Puerto": port,
		"Dirección": add,
	}).Info("GIN se ha conectado a un puerto y una dirección: ")
	log.Fatal(http.ListenAndServe(add + ":" + port, a.Router))
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

