package main

import (
	"fmt" //para realizar conversiones
	"strconv"
)

var num int
var texto string
var status bool = true /*este status funciona para todo el doc,
excepto cuando se define un nuevo status dentro de una función,
pues ese status tiene un valor local*/
//scope es el alcance

func main() {
	num2, num3, num4, texto, status := 2, 5, 67, "este es mi texto", false
	var moneda int64
	num2 = int(moneda) //se parsea para que queden del mismo tipo, SIEMPRE
	//los tipos de datos también son funciones en go
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(texto)
	fmt.Println("este es el status local de func main", status)
	fmt.Println("es valor de moneda es: ", moneda)

	//para convertir de un int a un string:
	texto = fmt.Sprintf("%d", moneda) /*%d es un verbo de conversión que dice que %d es número de base 10
	que se convertirá a string por la función Sprintf*/
	fmt.Println("el nuevo valor de texto que recibe el string de moneda es: ", texto)

	//convertir usando "strconv":
	texto = strconv.Itoa(int(moneda)) //convierte entero a texto
	fmt.Println("texto con strconv: ", texto)
	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println("Este es el status global para este main: ", status)
}
