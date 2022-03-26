package main

import (
	"fmt"
	"time"
)

//Channels hace que nuestras go routines paralelas puedan enviar la información pertinente
// y que no se termine el programa de manera abrupta
//Un canal es un espacio de memoria en el que dialogan distintas rutinas, que transporta algún tipo de dato

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1) //ejecuta lo que está en bucle pero se da un canal para comunicarse con otras rutinas
	fmt.Println("llegué hasta aquí")
	//sintaxis para que se está escuchando al canal1
	msg := <-canal1 //como el await de Node, espera que le llegue un valor para ejecutarse
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) { //acá se comunica el main y el bucle
	inicio := time.Now()
	for i := 0; i < 1000000000; i++ {

	}
	final := time.Now()
	canal1 <- final.Sub(inicio) //se asigna a canal1 el valor de esa función
	//final.Sub(inicio) es la diferencia entre inicio y final, es la duración que es un tipo de dato
}
