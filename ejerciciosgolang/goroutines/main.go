package main

import (
	"fmt"
	"strings"
	"time"
)

//GORoutines: ejecución asíncrona

func main() {
	go miNombreLento("Antu") //poniendo go lo trabaja de manera asíncrona
	fmt.Println("Estoy aquí")
	var x string
	fmt.Scanln(&x) //no termina de ejecutarse miNombreLento si escribo algo mientras se ejecuta
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "") //separa las letras por una cadena vacía
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) //hago que cada letra se muestre en cada segundo
		fmt.Println(letra)
	}
}
