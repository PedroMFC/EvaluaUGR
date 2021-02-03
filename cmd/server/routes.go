package server

import (
	log "github.com/sirupsen/logrus"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/PedroMFC/EvaluaUGR/internal/customlog"
	"net/http"
	"net"

	"github.com/gin-gonic/gin"

	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/handlersval"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/modelsres"
	"github.com/PedroMFC/EvaluaUGR/internal/microres/handlersres"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/modelspre"
	"github.com/PedroMFC/EvaluaUGR/internal/micropre/handlerspre"
	"github.com/PedroMFC/EvaluaUGR/internal"

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

func (a *applicationGin) Start() {
	
	config := internal.GetConfig()
	port := ""
	add := ""

	// Conectamos con el cliente de etcd
	etcdHost := os.Getenv(config.EtcdHost) 
	etcdPort := os.Getenv(config.EtcdPort) 
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
	port = os.Getenv(config.PortVarName); 
	log.WithFields(log.Fields{
		"Puerto": port,
	}).Info("Puerto leído desde entorno")

	if port != ""{
		log.Info("Usamos la configuración del entorno")
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, config.PortVarName, port)

		if err != nil{
			log.Println("No se ha podido escribir la clave", err)
		}
	// Si falla con el .env, usamos un puerto por defecto
	//Almacenamos en etcd
	} else{
		log.Info("Usamos la configuración por defecto")
		port = config.PortDefault
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, config.PortVarName, port)

		if err != nil{
			log.Println("No se ha podido escribir la clave", err)
		}
	}

	/* DIRECCIÓN */
	// Si hay una entrada en el .env 
	// Almacenamos en etcd
	add = os.Getenv(config.AddVarName); 
	log.WithFields(log.Fields{
		"Dirección": add,
	}).Info("Dirección leída desde entorno")

	if add != ""{
		log.Info("Usamos la dirección del entorno")
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, config.AddVarName, add)

		if err != nil{
			log.Println("No se ha podido escribir la clave", err)
		}
	// Si falla con el .env, usamos un puerto por defecto
	//Almacenamos en etcd
	} else{
		log.Info("Usamos la configuración para la dirección por defecto")
		add = config.AddDefault
		ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
		_, err := cli.Put(ctx, config.AddVarName, add)

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
	config := internal.GetConfig()
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

	// Incluimos logstash
	logstashHost := os.Getenv(config.LogHost) 
	logstassPort := os.Getenv(config.LogPort) 
	conn, err := net.Dial("tcp", logstashHost + ":" + logstassPort)
    if err != nil {
        log.Println(err)
	} else {
		hook := logrustash.New(conn, logrustash.DefaultFormatter(log.Fields{"type": "EvaluaUGR"}))
		logger.AddHook(hook)
		log.AddHook(hook)
	}

	router.Use(customlog.Logger(logger), gin.Recovery())

	//Aquí definimos las rutas
	//Valoraciones
	router.PUT("/asignaturas/:asig/valoraciones", handlersval.CrearAsignatura(ValRepo))
	router.GET("/asignaturas/:asig/valoraciones", handlersval.GetValoraciones(ValRepo))
	router.POST("/asignaturas/:asig/valoraciones/:val", handlersval.Valorar(ValRepo))
	router.GET("/asignaturaspeorvalorada", handlersval.GetPeor(ValRepo))
	router.GET("/asignaturasmejorvalorada", handlersval.GetMejor(ValRepo))
	router.GET("/asignaturas/:asig/valoraciones/media", handlersval.GetMedia(ValRepo))

	//Reseñas
	router.PUT("/asignaturas/:asig/resenias", handlersres.CrearAsignatura(ResRepo))
	router.GET("/asignaturas/:asig/resenias", handlersres.GetResenias(ResRepo))
	router.GET("/asignaturas/:asig/resenias/:id", handlersres.GetResenia(ResRepo))
	router.POST("/asignaturas/:asig/resenias", handlersres.Opinar(ResRepo))
	router.POST("/asignaturas/:asig/resenias/:id/gusta", handlersres.GustaResenia(ResRepo))
	router.POST("/asignaturas/:asig/resenias/:id/nogusta", handlersres.NoGustaResenia(ResRepo))

	//Preguntas
	router.PUT("/asignaturas/:asig/preguntas", handlerspre.CrearAsignatura(PreRepo))
	router.GET("/asignaturas/:asig/preguntas", handlerspre.GetPreguntas(PreRepo))
	router.GET("/asignaturas/:asig/preguntas/:id", handlerspre.GetPregunta(PreRepo))
	router.POST("/asignaturas/:asig/preguntas", handlerspre.Preguntar(PreRepo))
	router.POST("/asignaturas/:asig/preguntas/:id", handlerspre.Responder(PreRepo))

	return &applicationGin{Router: router}
}

