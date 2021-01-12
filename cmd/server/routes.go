package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/handlersval"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"github.com/joho/godotenv"
	"os"
	"time"
)

/*Singleton of Truth VALORACIONES*/

var ValRepo modelsval.ValoracionRepositorio
var ValMap  modelsval.ValoracionMap
func StartDataVal(){
	ValMap = *modelsval.NewValoracionMap()
	val := new(modelsval.Valoracion)
	val2 := new(modelsval.Valoracion)
	val.Valoracion = 2
	val2.Valoracion = 3
	ValMap.Valoraciones["AAA"] = []modelsval.Valoracion{*val, *val2}
	ValRepo = *modelsval.NewValoracionsRepositorio(&ValMap)
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
			log.Println("Error loading .env file")
		}

		// Si hay una entrada en el .env
		//Almacenamos en etcd
		port = os.Getenv(PortVarName); 
		log.Println("Puerto del .env" + port)
		if port != ""{
			log.Println("Usamos la configuración del .env")
			ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
			_, err := cli.Put(ctx, PortVarName, port)

			if err != nil{
				log.Fatal("No se ha podido escribir la clave", err)
			}
		// Si las anteriores han fallado, usamos un puerto por defecto
		//Almacenamos en etcd
		} else{
			log.Println("Usamos la configuración por defecto")
			port = "8080"
			ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
			_, err := cli.Put(ctx, PortVarName, port)

			if err != nil{
				log.Fatal("No se ha podido escribir la clave", err)
			}
		}

	}

	log.Println("GIN valoraciones conectado en " + port)
	log.Fatal(http.ListenAndServe("localhost:" + port, a.Router))
}


func NewAppGin() *applicationGin {
	router := gin.Default()

	//Aquí definimos las rutas
	//Valoraciones
	router.PUT("valoraciones/:asig", handlersval.CrearAsignatura(ValRepo))
	router.GET("/valoraciones/:asig", handlersval.GetValoraciones(ValRepo))

	return &applicationGin{Router: router}
}

