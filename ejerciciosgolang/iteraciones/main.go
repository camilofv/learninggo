package main

import (
	"fmt"
)

func main() {
	//el for más simple
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//el for más normal
	fmt.Println("for normal: ")
	for i := 1; i < 8; i++ {
		fmt.Println(i)
	}
	//for en un loop indefinido, para romper el loop
	numero := 0
	k := 1
	for {
		fmt.Println("continúo")
		if k == 3 {
			fmt.Println("El número secreto es 100")
		}
		k++
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	//instrucción continue
	var j = 0
	for j < 10 {
		fmt.Printf("\n valor de j: %d", j) // \n es new line
		if j == 5 {
			fmt.Printf("   Multiplicamos por 2 \n") //Printf + \ es lo mismo que Println
			j = j * 2
			continue /*acá ingresó y continue manda la ejecución de nuevo al inicio del ciclo for,
			torciendo la iteración para volver a chequear la condición*/
		}
		fmt.Printf("   Pasó por aquí \n")
		j++
	}

	//go to permite enviar la rutina a un comienzo de rutina determinado
	var m int
RUTINA: /*esto demarca una sección en el código,
	es como un marcador o etiqueta y es inocuo a la ejecución, no imprime nada*/
	for m < 10 {
		if m == 4 {
			m = m + 2
			fmt.Println("voy a RUTINA sumando 2 a variable m")
			goto RUTINA
		}
		fmt.Printf("Valor de m : %d\n", m)
		m++
	}

}
