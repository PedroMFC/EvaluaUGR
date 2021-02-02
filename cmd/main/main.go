package main


import (
	"github.com/PedroMFC/EvaluaUGR/cmd/server"

)

func main(){
	parar := make (chan int)

	server.StartData()
	
	go server.NewAppGin().Start()

	<-parar
}