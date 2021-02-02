package modelsval

import(
	"os"
	"fmt"
	//"log"
	"strconv"
	"database/sql"
	"github.com/lib/pq"
)

func ConnectValoracion( DbUser, DbPassword, DbPort, DbHost, DbName string) *sql.DB{
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := sql.Open("postgres", DBURL)

	if err != nil {
		panic("Failed to connect to database!")
	}

	_, err = db.Exec("DROP TABLE IF EXISTS val_table;")

	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE val_table (asignatura	text PRIMARY KEY, valoraciones integer[]);")

	if err != nil {
		panic(err)
	}

	return db
}

//ValoracionMap Tabla Hash que sigue la interfaz IValSaver
type ValoracionDB struct {
	DB *sql.DB
}

//NewValoracionMap devuevle un ValoracionMap
func NewValoracionDB() *ValoracionDB {
	db := ConnectValoracion(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	return &ValoracionDB{DB:db}
}

//CrearAsignatura crea una entrada para una asignatura
func (valDB *ValoracionDB) CrearAsignatura(asignatura string) {
	sqlStatement := `DELETE FROM val_table WHERE asignatura = $1`
	_,err := valDB.DB.Exec(sqlStatement, asignatura)

	if err != nil {
		panic(err)
	}

	sqlStatement = `INSERT INTO val_table VALUES ($1,$2);`

	valoraciones := []int{}
	_,err = valDB.DB.Exec(sqlStatement, asignatura, pq.Array(valoraciones) )

	if err != nil {
		panic(err)
	}
}

//AsignaturaRegistrada comprueba si una asignatura está registrada
func (valDB *ValoracionDB) AsignaturaRegistrada(asignatura string) bool {
	asignaturas := valDB.ObtenerAsignaturas()

	for _, val := range asignaturas{
		if val == asignatura{
			return true
		}
	}

	return false
}

//GuardarValoracion alamcena una valoración
func (valDB *ValoracionDB) GuardarValoracion(asignatura string, val *Valoracion) {	
	sqlStatement := `UPDATE val_table SET valoraciones=array_append(valoraciones,$1) WHERE asignatura=$2 ;`
	_,err := valDB.DB.Exec(sqlStatement, val.Valoracion ,asignatura)

	if err != nil{
		panic(err)
	}

}

//ObtenerValoraciones devuelve valoraciones de una asignatura
func (valDB *ValoracionDB) ObtenerValoraciones(asignatura string) []Valoracion {
	sqlStatement := `SELECT valoraciones FROM val_table WHERE asignatura=$1;`

	var valoracionesDB []string


	if err := valDB.DB.QueryRow(sqlStatement, asignatura).Scan(pq.Array(&valoracionesDB)); err != nil {
        panic(err)
	}

	
	var valoraciones []Valoracion
	for _, val := range valoracionesDB{
		i,_ := strconv.Atoi(val)
		valoraciones = append(valoraciones, Valoracion{i})
	}

	return valoraciones
}

//ObtenerAsignaturas devuelve las asignaturas almacenadas
func (valDB *ValoracionDB) ObtenerAsignaturas() []string {

	sqlStatement := `SELECT asignatura FROM val_table;`

	var asignaturasDB []string


	rows, err := valDB.DB.Query(sqlStatement)

	if err != nil {
		// handle this error better than this
		panic(err)
	}

	for rows.Next() {
		var asig string
		err = rows.Scan(&asig)
		if err != nil {
		  // handle this error
		  panic(err)
		}
		asignaturasDB = append(asignaturasDB, asig)
	}
	  // get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return asignaturasDB
}