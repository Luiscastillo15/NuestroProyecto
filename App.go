package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var Plantillas = template.Must(template.ParseGlob("Plantillas/*"))

func conectDb() (*sql.DB, error) {
	db, error := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/proyecto")
	if error != nil {
		panic(error.Error())
	}

	error = db.Ping()

	if error != nil {
		return nil, error
	}

	fmt.Print("conexion exitosa Mysql")
	return db, nil
}

func selectPhones(db *sql.DB) {
	rows, erro := db.Query("select * from telefonos")
	if erro != nil {
		fmt.Print("eror en select")
		panic(erro.Error())
	}
	defer rows.Close()

	type Phones struct {
		id     int64
		marca  string
		modelo string
		precio float64
	}

	var telefonos []Phones
	for rows.Next() {
		var t Phones
		erro := rows.Scan(&t.id, &t.marca, &t.modelo)
		if erro != nil {
			panic(erro.Error())
		}
		telefonos = append(telefonos, t)
	}
	err := rows.Err()
	if err != nil {
		panic(err.Error())
	}

	for _, u := range telefonos {
		fmt.Println(u.id, u.modelo, u.marca)
	}
}

func main() {

	http.HandleFunc("/", inicio)
	http.HandleFunc("/agregar", Agregar)

	fmt.Println("servidor corriendo:3000")
	db, error := conectDb()
	if error != nil {
		panic(error.Error())
	}
	selectPhones(db)
	http.ListenAndServe("localhost:3000", nil)

}
func inicio(rw http.ResponseWriter, r *http.Request) {

	Plantillas.ExecuteTemplate(rw, "inicio", nil)

}
func Agregar(rw http.ResponseWriter, r *http.Request) {

	Plantillas.ExecuteTemplate(rw, "agregar", nil)

}
