package main

import (
	"net/http" //todo lo que es redes
)

//colocar un servidor web para recibir peticiones por el navegador y poder responder

func main() {
	//HandleFunc maneja el ingreso en mi localhost y va a ejecutar la función
	//el primer parámetro es la ruta y el segundo es la función que debe ejecutar
	http.HandleFunc("/", home) //el servidor detectará cuando en localhost esté en home o en barra /
	//ahora, go debe escuchar un puerto determinado se ingrese información
	http.ListenAndServe(":3000", nil) //cuando recibe una instrucción en el puerto, la ejecuta
}

//creación de función con los dos parámetros que nos envía http
func home(w http.ResponseWriter, r *http.Request) { //Response es cuando se recibe, request es cuando envío la respuesta
	http.ServeFile(w, r, "./index.html") //cuando se ejecute función home, debe mostrar index.html
}
