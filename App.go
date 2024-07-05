package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var Plantillas = template.Must(template.ParseGlob("Plantillas/*"))
func main() {
	http.HandleFunc("/", inicio)
	http.HandleFunc("/agregar", Agregar)
	fmt.Println("servidor corriendo:3000")
	http.ListenAndServe("localhost:3000", nil)
}
func inicio(rw http.ResponseWriter, r *http.Request) {

	Plantillas.ExecuteTemplate(rw, "inicio", nil)

}
func Agregar(rw http.ResponseWriter, r *http.Request) {

	Plantillas.ExecuteTemplate(rw, "agregar", nil)

}
