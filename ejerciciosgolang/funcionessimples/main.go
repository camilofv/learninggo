package main

import (
	"fmt"
)

func main() {
	fmt.Println(uno(5))

	numero, estado := dos(1)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(5))
	fmt.Println(calculo(2, 4, 7, 9, 0, 12))
}

/*forma simple
func uno(numero int) int {
	return numero * 2
}
*/

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

//parámetros variables: son parámetros dinámicos cuando no sé cuántos parámetros entran
func calculo(numero ...int) int { //con los ... indica que no sabe la cantidad de parámetros de entrada
	var total int
	/*for _, num := range numero { //range permite iterar lo ingresado que ha sido convertido en una lista de elementos en la memoria
		//range devulve dos valores, primero el número del elemento
		total = total + num
	}
	return total
	*/
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
