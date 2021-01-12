package modelsval

//ValoracionMap Tabla Hash que sigue la interfaz IValSaver
type ValoracionMap struct {
	Valoraciones map[string][]Valoracion
}

//NewValoracionMap devuevle un ValoracionMap
func NewValoracionMap() *ValoracionMap {
	return &ValoracionMap{Valoraciones: make(map[string][]Valoracion)}
}

//CrearAsignatura crea una entrada para una asignatura
func (valMap *ValoracionMap) CrearAsignatura(asignatura string) {
	valMap.Valoraciones[asignatura] = []Valoracion{}
}

//AsignaturaRegistrada comprueba si una asignatura está registrada
func (valMap *ValoracionMap) AsignaturaRegistrada(asignatura string) bool {
	return valMap.Valoraciones[asignatura] != nil
}

//GuardarValoracion alamcena una valoración
func (valMap *ValoracionMap) GuardarValoracion(asignatura string, val *Valoracion) {
	valMap.Valoraciones[asignatura] = append(valMap.Valoraciones[asignatura], *val)
}

//ObtenerValoraciones devuelve valoraciones de una asignatura
func (valMap *ValoracionMap) ObtenerValoraciones(asignatura string) []Valoracion {

	return valMap.Valoraciones[asignatura]
}

//ObtenerAsignaturas devuelve las asignaturas almacenadas
func (valMap *ValoracionMap) ObtenerAsignaturas() []string {

	keys := make([]string, 0, len(valMap.Valoraciones))
	
	for k := range valMap.Valoraciones {
		keys = append(keys, k)
	}

	return keys
}