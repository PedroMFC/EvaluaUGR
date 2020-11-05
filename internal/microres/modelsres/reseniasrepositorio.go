package modelsres


//Contiene las reseñas realizadas
type ReseniasRepositorio struct {
	Reseniass map[string][]Resenia  //Con la tabla Hash hace la consultas más rápido
}
