package modelsval

//ValoracionMap Tabla Hash que sigue la interfaz IValSaver
type ValoracionMap struct {
	Valoraciones map[string][]Valoracion
}

//NewValoracionMap devuevle un ValoracionMap
func NewValoracionMap() *ValoracionMap {
	return &ValoracionMap{Valoraciones: make(map[string][]Valoracion)}
}

//GuardarValoracion alamcena una valoración
func (valMap *ValoracionMap) GuardarValoracion(asignatura string, val *Valoracion) {
	if valMap.Valoraciones[asignatura] != nil { // Si ya hay valoraciones antes se añaden a las existentes
		valMap.Valoraciones[asignatura] = append(valMap.Valoraciones[asignatura], *val)
	} else { //Si no, tenemos que crear una nueva
		valMap.Valoraciones[asignatura] = []Valoracion{*val}
	}
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