package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
"strconv"
	_ "github.com/go-sql-driver/mysql"
)


var Plantillas = template.Must(template.ParseGlob("Plantillas/*"))


type Phones struct {
	ID     int64
	Marca  string
	Modelo string
	Precio float64
}

func conectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/proyecto")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Conexión exitosa a MySQL")
	return db, nil
}

//crud 
func selectPhones(db *sql.DB) ([]Phones, error) {
	rows, err := db.Query("SELECT * FROM telefonos")
	if err != nil {
		return nil, fmt.Errorf("error en select: %v", err)
	}
	defer rows.Close()

	var telefonos []Phones
	for rows.Next() {
		var t Phones
		err := rows.Scan(&t.ID, &t.Marca, &t.Modelo, &t.Precio)
		if err != nil {
			return nil, err
		}
		telefonos = append(telefonos, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for _, u := range telefonos {
		fmt.Println(u.ID, u.Modelo, u.Marca)
	}
	return telefonos, nil
}

// selectPhoneByID consulta la base de datos y retorna un teléfono por ID
func selectPhoneByID(db *sql.DB, id int64) (Phones, error) {
	var t Phones
	row := db.QueryRow("SELECT id, marca, modelo, precio FROM telefonos WHERE id = ?", id)
	err := row.Scan(&t.ID, &t.Marca, &t.Modelo, &t.Precio)
	if err != nil {
		return t, fmt.Errorf("error al seleccionar teléfono: %v", err)
	}
	return t, nil
}
func insertPhone(db *sql.DB, marca, modelo string, precio float64) error {
	query := "INSERT INTO telefonos (marca, modelo, precio) VALUES (?, ?, ?)"
	_, err := db.Exec(query, marca, modelo, precio)
	if err != nil {
		return fmt.Errorf("error al insertar teléfono: %v", err)
	}
	return nil
}

func deletePhone(db *sql.DB, id int64) error {
	query := "DELETE FROM telefonos WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar teléfono: %v", err)
	}
	return nil
}
// updatePhone actualiza un teléfono en la base de datos por ID
func updatePhone(db *sql.DB, id int64, marca, modelo string, precio float64) error {
	query := "UPDATE telefonos SET marca = ?, modelo = ?, precio = ? WHERE id = ?"
	_, err := db.Exec(query, marca, modelo, precio, id)
	if err != nil {
		return fmt.Errorf("error al actualizar teléfono: %v", err)
	}
	return nil
}


// main configura las rutas y lanza el servidor web
func main() {
	http.HandleFunc("/", inicio)
	http.HandleFunc("/agregar", agregar)
	http.HandleFunc("/eliminar", eliminar)
	http.HandleFunc("/editar", editar)
	fmt.Println("Servidor corriendo en localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}


func inicio(rw http.ResponseWriter, r *http.Request) {
	db, err := conectDb()
	
	if err != nil {
		http.Error(rw, "Error conectando a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err != nil {
		http.Error(rw, "Error obteniendo teléfonos", http.StatusInternalServerError)
		return
	}
	telefonos, err := selectPhones(db)
	err = Plantillas.ExecuteTemplate(rw, "inicio", telefonos)
	if err != nil {
		http.Error(rw, "Error renderizando plantilla", http.StatusInternalServerError)
		return
	}
}

func agregar(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := Plantillas.ExecuteTemplate(rw, "agregar", nil)
		if err != nil {
			http.Error(rw, "Error renderizando plantilla", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		db, err := conectDb()
		if err != nil {
			http.Error(rw, "Error conectando a la base de datos", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Obtener los datos del formulario
		marca := r.FormValue("marca")
		modelo := r.FormValue("modelo")
		precio := r.FormValue("precio")

		// Convertir precio a float64
		var precioFloat float64
		fmt.Sscanf(precio, "%f", &precioFloat)

		// Insertar el nuevo teléfono en la base de datos
		err = insertPhone(db, marca, modelo, precioFloat)
		if err != nil {
			http.Error(rw, "Error agregando teléfono", http.StatusInternalServerError)
			return
		}

		// Redirigir a la página principal después de agregar
		http.Redirect(rw, r, "/", http.StatusSeeOther)
	}
}


func eliminar(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := conectDb()
		if err != nil {
			http.Error(rw, "Error conectando a la base de datos", http.StatusInternalServerError)
			return
		}
		defer db.Close()


		id := r.FormValue("id")
		var idInt int64
		fmt.Sscanf(id, "%d", &idInt)


		err = deletePhone(db, idInt)
		if err != nil {
			http.Error(rw, "Error eliminando teléfono", http.StatusInternalServerError)
			return
		}


		http.Redirect(rw, r, "/", http.StatusSeeOther)
	}
}


func editar(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			http.Error(rw, "ID inválido", http.StatusBadRequest)
			return
		}

		db, err := conectDb()
		if err != nil {
			http.Error(rw, "Error conectando a la base de datos", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		telefono, err := selectPhoneByID(db, idInt)
		if err != nil {
			http.Error(rw, "Error obteniendo teléfono", http.StatusInternalServerError)
			return
		}

		err = Plantillas.ExecuteTemplate(rw, "editar.html", telefono)
		if err != nil {
			http.Error(rw, "Error renderizando plantilla " + err.Error(), http.StatusInternalServerError)
		}
	}

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		marca := r.FormValue("marca")
		modelo := r.FormValue("modelo")
		precio := r.FormValue("precio")

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			http.Error(rw, "ID inválido", http.StatusBadRequest)
			return
		}

		precioFloat, err := strconv.ParseFloat(precio, 64)
		if err != nil {
			http.Error(rw, "Precio inválido", http.StatusBadRequest)
			return
		}

		db, err := conectDb()
		if err != nil {
			http.Error(rw, "Error conectando a la base de datos", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		err = updatePhone(db, idInt, marca, modelo, precioFloat)
		if err != nil {
			http.Error(rw, "Error actualizando teléfono", http.StatusInternalServerError)
			return
		}

		http.Redirect(rw, r, "/", http.StatusSeeOther)
	}
}