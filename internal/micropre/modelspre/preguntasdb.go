package modelspre

import (
	"os"
	"fmt"
	//"log"
	"strconv"
	"sort"

	"github.com/PedroMFC/EvaluaUGR/internal/micropre/errorspre"

	"database/sql"
)

func ConnectPreguntas( DbUser, DbPassword, DbPort, DbHost, DbName string) *sql.DB{
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := sql.Open("postgres", DBURL)

	if err != nil {
		panic("Failed to connect to database!")
	}

	_, err = db.Exec("DROP TABLE IF EXISTS asig_pre_table;")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS pre_table;")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS res_pre_table;")

	if err != nil {
		panic(err)
	}



	_, err = db.Exec("CREATE TABLE asig_pre_table ( asignatura text PRIMARY KEY);")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE pre_table ( asignatura text, pregunta text, identificador integer);")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE res_pre_table ( asignatura text, pregunta_id integer, respuesta text);")
	if err != nil {
		panic(err)
	}


	return db
}

type PreguntasDB struct {
	DB *sql.DB
}

func NewPreguntasDB() *PreguntasDB {
	db := ConnectPreguntas(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	return &PreguntasDB{DB:db}
}

//CrearAsignatura crea una entrada para una asignatura
func (preDB *PreguntasDB) CrearAsignatura(asignatura string) {
	sqlStatement := `DELETE FROM pre_table WHERE asignatura = $1`
	_,err := preDB.DB.Exec(sqlStatement, asignatura)

	if err != nil {
		panic(err)
	}

	sqlStatement = `DELETE FROM res_pre_table WHERE asignatura = $1`
	_,err = preDB.DB.Exec(sqlStatement, asignatura)

	if err != nil {
		panic(err)
	}


	sqlStatement = `DELETE FROM asig_pre_table WHERE asignatura = $1`
	_,err = preDB.DB.Exec(sqlStatement, asignatura)

	if err != nil {
		panic(err)
	}

	sqlStatement = `INSERT INTO asig_pre_table VALUES ($1);`

	_,err = preDB.DB.Exec(sqlStatement, asignatura )

	if err != nil {
		panic(err)
	}
}

//AsignaturaRegistrada comprueba si una asignatura está registrada
func (preDB *PreguntasDB) AsignaturaRegistrada(asignatura string) bool {
	asignaturas := preDB.ObtenerAsignaturas()

	for _, val := range asignaturas{
		if val == asignatura{
			return true
		}
	}

	return false
}

func (preDB *PreguntasDB) GuardarPregunta(asignatura string, pre *Pregunta){
	sqlStatement := `INSERT INTO pre_table VALUES ($1,$2,$3);`

	id := preDB.contarPreguntas(asignatura)

	_,err := preDB.DB.Exec(sqlStatement, asignatura, pre.Pregunta, id)

	if err != nil {
		panic(err)
	}
}

func (preDB *PreguntasDB) ObtenerPregunta(asignatura string) []Pregunta {
	sqlStatement := `SELECT pregunta, identificador FROM pre_table WHERE asignatura = $1`

	rows, err := preDB.DB.Query(sqlStatement, asignatura)

	if err != nil {
		// handle this error better than this
		panic(err)
	}

	var preguntasDB []Pregunta
	for rows.Next() {
		var pregunta string
		var identificador string
		err = rows.Scan(&pregunta, &identificador)
		if err != nil {
		  // handle this error
		  panic(err)
		}
		var pre Pregunta
		pre.Pregunta = pregunta
		pre.Identificador,_ = strconv.Atoi(identificador)
		preguntasDB = append(preguntasDB, pre)
	}

	//Ahora le incluimmos las respuestas guardadas
	preguntasRespuestasDB := preguntasDB
	for i,pre :=range preguntasDB{
		sqlStatement = `SELECT respuesta FROM res_pre_table WHERE (asignatura = $1 AND pregunta_id=$2)`

		rows, err := preDB.DB.Query(sqlStatement, asignatura, pre.Identificador)

		if err != nil {
			// handle this error better than this
			panic(err)
		}

		var respuestas []Respuesta
		for rows.Next() {
			var respuesta string
			err = rows.Scan(&respuesta)
			if err != nil {
			// handle this error
			panic(err)
			}
			var res Respuesta
			res.Respuesta = respuesta
			respuestas = append(respuestas, res)
		}		

		preguntasRespuestasDB[i].Respuestas = respuestas
	}


	//Ordenamos según identificador
	if (len(preguntasRespuestasDB) > 1){
		sort.Slice(preguntasRespuestasDB, func(i,j int) bool{
			return preguntasRespuestasDB[i].Identificador < preguntasRespuestasDB[j].Identificador
		})
	}

	return preguntasRespuestasDB
}

func (preDB *PreguntasDB) Responder(asignatura string, id int, res *Respuesta) error {
	actuales := preDB.contarPreguntas(asignatura)

	if id > (actuales -1) {
		return &errorspre.ErrorPregunta{" no hay una pregunta con ese identificador para la asignatura señalada"}
	}

	sqlStatement := `INSERT INTO res_pre_table VALUES ($1,$2,$3);`

	_,err := preDB.DB.Exec(sqlStatement, asignatura, id, res.Respuesta)

	if err != nil {
		panic(err)
	}

	return nil
}



//ObtenerAsignaturas devuelve las asignaturas almacenadas
func (preDB *PreguntasDB) ObtenerAsignaturas() []string {

	sqlStatement := `SELECT asignatura FROM asig_pre_table;`

	var asignaturasDB []string


	rows, err := preDB.DB.Query(sqlStatement)

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

func (preDB *PreguntasDB) contarPreguntas(asignatura string) int {
	contador := 0
	
	sqlStatement := `SELECT * FROM pre_table WHERE asignatura = $1;`

	rows,_ := preDB.DB.Query(sqlStatement, asignatura)

	for rows.Next() {
		contador++
	}

	return contador
}
