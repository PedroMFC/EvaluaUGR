package modelsres

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"sort"

	"github.com/PedroMFC/EvaluaUGR/internal/microres/errorsres"

	"database/sql"
	//"github.com/lib/pq"
)


func ConnectResenias( DbUser, DbPassword, DbPort, DbHost, DbName string) *sql.DB{
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := sql.Open("postgres", DBURL)

	if err != nil {
		panic("Failed to connect to database!")
	}

	_, err = db.Exec("DROP TABLE IF EXISTS res_table;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS asig_res_table;")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE asig_res_table ( asignatura text PRIMARY KEY);")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE res_table ( asignatura text, opinion text, identificador integer, megusta integer, nomegusta integer);")

	if err != nil {
		panic(err)
	}


	return db
}


type ReseniasDB struct {
	DB *sql.DB
}

func NewReseniasDB() *ReseniasDB {
	db := ConnectResenias(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	return &ReseniasDB{DB:db}
}

//CrearAsignatura crea una entrada para una asignatura
func (resDB *ReseniasDB) CrearAsignatura(asignatura string) {
	sqlStatement := `DELETE FROM res_table WHERE asignatura = $1`
	_,err := resDB.DB.Exec(sqlStatement, asignatura)

	if err != nil {
		panic(err)
	}

	sqlStatement = `DELETE FROM asig_res_table WHERE asignatura = $1`
	_,err = resDB.DB.Exec(sqlStatement, asignatura)

	if err != nil {
		panic(err)
	}

	sqlStatement = `INSERT INTO asig_res_table VALUES ($1);`

	_,err = resDB.DB.Exec(sqlStatement, asignatura )

	if err != nil {
		panic(err)
	}
}

//AsignaturaRegistrada comprueba si una asignatura está registrada
func (resDB *ReseniasDB) AsignaturaRegistrada(asignatura string) bool {
	asignaturas := resDB.ObtenerAsignaturas()

	for _, val := range asignaturas{
		if val == asignatura{
			return true
		}
	}

	return false
}


func (resDB *ReseniasDB) GuardarResenia(asignatura string, opinion *Resenia){
	sqlStatement := `INSERT INTO res_table VALUES ($1,$2,$3,$4,$5);`

	id := resDB.contarResenias(asignatura)

	_,err := resDB.DB.Exec(sqlStatement, asignatura, opinion.Opinion, id, opinion.MeGusta, opinion.NoMeGusta )
	

	if err != nil {
		panic(err)
	}

}

func (resDB *ReseniasDB) ObtenerResenias(asignatura string) []Resenia{
	sqlStatement := `SELECT opinion, identificador, megusta, nomegusta FROM res_table WHERE asignatura = $1`

	rows, err := resDB.DB.Query(sqlStatement, asignatura)

	if err != nil {
		// handle this error better than this
		panic(err)
	}

	var reseniasDB []Resenia
	for rows.Next() {
		var opinion string
		var identificador string
		var gusta string
		var nogusta string
		err = rows.Scan(&opinion, &identificador, &gusta, &nogusta)
		if err != nil {
		  // handle this error
		  panic(err)
		}
		var resenia Resenia
		resenia.Opinion = opinion
		resenia.Identificador,_ = strconv.Atoi(identificador)
		resenia.MeGusta,_ = strconv.Atoi(gusta)
		resenia.NoMeGusta,_ = strconv.Atoi(nogusta)
		reseniasDB = append(reseniasDB, resenia)
	}

	if (len(reseniasDB) > 1){
		sort.Slice(reseniasDB, func(i,j int) bool{
			return reseniasDB[i].Identificador < reseniasDB[j].Identificador
		})
	}
	return reseniasDB
}

func (resDB *ReseniasDB) MeGustaResenia(asignatura string, id int) error {
	actuales := resDB.contarResenias(asignatura)

	if id > (actuales -1) {
		return &errorsres.ErrorResenia{" la reseña no contiene in identificador válido"}
	}

	sqlStatement := `SELECT megusta FROM res_table WHERE (asignatura=$1 AND identificador=$2);`
	var gusta int
	err := resDB.DB.QueryRow(sqlStatement, asignatura, id).Scan(&gusta)
	if err != nil {
        log.Fatal(err)
	}

	gusta++

	sqlStatement = `UPDATE res_table SET megusta=$1 WHERE (asignatura=$2 AND identificador=$3);`

	_,err = resDB.DB.Exec(sqlStatement, gusta ,asignatura, id)

	if err != nil{
		panic(err)
	}

	return nil
}

func (resDB *ReseniasDB) NoMeGustaResenia(asignatura string, id int) error {
	actuales := resDB.contarResenias(asignatura)

	if id > (actuales -1) {
		return &errorsres.ErrorResenia{" la reseña no contiene in identificador válido"}
	}
	
	sqlStatement := `SELECT nomegusta FROM res_table WHERE (asignatura=$1 AND identificador=$2);`
	var gusta int
	err := resDB.DB.QueryRow(sqlStatement, asignatura, id).Scan(&gusta)
	if err != nil {
        log.Fatal(err)
	}

	gusta++

	sqlStatement = `UPDATE res_table SET nomegusta=$1 WHERE (asignatura=$2 AND identificador=$3);`

	_,err = resDB.DB.Exec(sqlStatement, gusta ,asignatura, id)

	if err != nil{
		panic(err)
	}

	return nil
}

//ObtenerAsignaturas devuelve las asignaturas almacenadas
func (resDB *ReseniasDB) ObtenerAsignaturas() []string {

	sqlStatement := `SELECT asignatura FROM asig_res_table;`

	var asignaturasDB []string


	rows, err := resDB.DB.Query(sqlStatement)

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

	log.Println("ASIGNATURAS ", asignaturasDB)

	return asignaturasDB
}

func (resDB *ReseniasDB) contarResenias(asignatura string) int {
	contador := 0
	
	sqlStatement := `SELECT * FROM res_table WHERE asignatura = $1;`

	rows,_ := resDB.DB.Query(sqlStatement, asignatura)

	for rows.Next() {
		contador++
	}

	return contador
}