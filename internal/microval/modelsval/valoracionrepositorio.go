package modelsval

import (
	"sort"

	"github.com/PedroMFC/EvaluaUGR/internal/asignatura/asig"
	//"fmt"
	"github.com/PedroMFC/EvaluaUGR/internal/microval/errorsval"
)

//ValoracionRepositorio Contiene las valoraciones realizadas
type ValoracionRepositorio struct {
	Valoraciones IValSaver
}

//NewValoracionsRepositorio devuelve una ValoracionRepositorio
func NewValoracionsRepositorio(val IValSaver) *ValoracionRepositorio {
	return &ValoracionRepositorio{Valoraciones: val}
}

//CrearAsignaura añade una asignatura para poder valorarla
func (valRepo *ValoracionRepositorio) CrearAsignatura(asignatura string) error{
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	valRepo.Valoraciones.CrearAsignatura(asignatura)

	return nil
}


//Valorar añade una valoración al repositorio
func (valRepo *ValoracionRepositorio) Valorar(asignatura string, numero int) error {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return err
	}

	val := new(Valoracion)
	err = val.SetValoracion(numero)
	if err != nil {
		return err
	}

	// Vemos que no está vacío
	if  !valRepo.Valoraciones.AsignaturaRegistrada(asignatura) {
		return &errorsval.ErrorValoracion{" la asignatura no está registrada"}
	}

	valRepo.Valoraciones.GuardarValoracion(asignatura, val)

	return nil
}

//GetValoraciones nos aporta las valoraciones realizadas en una asignatura
func (valRepo *ValoracionRepositorio) GetValoraciones(asignatura string) ([]Valoracion, error) {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return nil, err
	}

	// Vemos que no está vacío
	if  !valRepo.Valoraciones.AsignaturaRegistrada(asignatura) {
		return nil, &errorsval.ErrorValoracion{" la asignatura no está registrada"}
	}

	return valRepo.Valoraciones.ObtenerValoraciones(asignatura), nil
}

//GetMedia nos aporta la valoración media de una asignatura
func (valRepo *ValoracionRepositorio) GetMedia(asignatura string) (float64, error) {
	err := asig.AsignaturaCorrecta(asignatura)
	if err != nil {
		return 0, err
	}

	// Vemos que no está vacío
	if  !valRepo.Valoraciones.AsignaturaRegistrada(asignatura) {
		return -1, &errorsval.ErrorValoracion{" la asignatura no está registrada"}
	}
	
	valoraciones := valRepo.Valoraciones.ObtenerValoraciones(asignatura)

	//Si no está vacío, calculamos la media
	media := 0.0
	for _, val := range valoraciones {
		media = media + float64(val.Valoracion)
	}
	media = media / float64(len(valoraciones))

	return media, nil
}


//GetPeorValorada devuelve una lista con las asignaturas con peores valoraciones en media
func (valRepo *ValoracionRepositorio) GetPeorValorada() []string {
	menosValoradas := []string{}

	if len(valRepo.Valoraciones.ObtenerAsignaturas()) == 0 { //Si el repositorio está vacío
		return menosValoradas
	}

	mediasAsignaturas := []AsigMedia{}
	// Obtenemos el conjunto de medias
	for _,k := range valRepo.Valoraciones.ObtenerAsignaturas(){
		med, err := valRepo.GetMedia(k)
		if err == nil{
			mediasAsignaturas = append(mediasAsignaturas, AsigMedia{k, med})
		}
	}

	//Vemos las asignaturas que tienes menos valoraciones
	sort.Slice(mediasAsignaturas, func(i,j int) bool {
		return mediasAsignaturas[i].media < mediasAsignaturas[j].media
	})

	menosValoradas = append(menosValoradas, mediasAsignaturas[0].asig)

	if len(valRepo.Valoraciones.ObtenerAsignaturas()) > 1 { // Por si solo hay una asignatura
		i := 1
		for mediasAsignaturas[0].media == mediasAsignaturas[i].media {
			menosValoradas = append(menosValoradas, mediasAsignaturas[i].asig)
			i = i+1
		}
	}

	return menosValoradas
}

//GetMejorValorada devuelve una lista con las asignaturas con mejor valoración en media
func (valRepo *ValoracionRepositorio) GetMejorValorada() []string {
	menosValoradas := []string{}

	if len(valRepo.Valoraciones.ObtenerAsignaturas()) == 0 { //Si el repositorio está vacío
		return menosValoradas
	}

	mediasAsignaturas := []AsigMedia{}
	// Obtenemos el conjunto de medias
	for _,k := range valRepo.Valoraciones.ObtenerAsignaturas(){
		med, err := valRepo.GetMedia(k)
		if err == nil{
			mediasAsignaturas = append(mediasAsignaturas, AsigMedia{k, med})
		}
	}

	//Vemos las asignaturas que tienes menos valoraciones
	sort.Slice(mediasAsignaturas, func(i,j int) bool {
		return mediasAsignaturas[i].media > mediasAsignaturas[j].media
	})

	menosValoradas = append(menosValoradas, mediasAsignaturas[0].asig)

	if len(valRepo.Valoraciones.ObtenerAsignaturas()) > 1 { // Por si solo hay una asignatura
		i := 1
		for mediasAsignaturas[0].media == mediasAsignaturas[i].media {
			menosValoradas = append(menosValoradas, mediasAsignaturas[i].asig)
			i = i+1
		}
	}

	return menosValoradas
}
