package main

import (
	"bufio" //lectura de archivos, lo que se ingresa por teclado, grabaciones y otros
	"fmt"
	"os"
)

var numero1, numero2, resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese número 1: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese número 2: ")
	fmt.Scanln(&numero2)

	fmt.Println("Descripción :")

	scanner := bufio.NewScanner(os.Stdin) //input por standar input, o sea, el teclado
	if scanner.Scan() {                   //se evalúa con Scan si se obtiene algo
		leyenda = scanner.Text()
		fmt.Println()
	}

	fmt.Println(leyenda)
	fmt.Println("la suma es: ", numero1+numero2)
}
