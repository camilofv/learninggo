package main

import (
	"fmt"
)

func main() {
	exponencia(2)
}

func exponencia(numero int) { //l primero es crear un cierre de la función
	if numero > 100000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2) //se autollama a la misma función dentro de la función
}
