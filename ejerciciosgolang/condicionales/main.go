package main

import (
	"fmt"
)

var estado bool

func main() {
	estado = true
	if estado = false; estado == true { //ejemplo de asignar un valor de una variable dentro del if
		fmt.Println(estado)
	} else {
		fmt.Println("es falso")
	}

	switch numero := 7; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("el número es mayor a 5")
	}
}
