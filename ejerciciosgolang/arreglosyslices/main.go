package main

import (
	"fmt"
)

var tabla [10]int

func main() {
	tabla[0] = 1
	tabla[5] = 15

	fmt.Println(tabla)

	//otra forma de inicializar un arreglo (también llamado vector)
	tablados := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 20}
	fmt.Println(tablados)

	for i := 0; i < len(tablados); i++ {
		fmt.Println(tablados[i])
	}

	//las matrices son arrays de dos dimensiones
	var matriz [3][3]int //se lee [filas][columnas]
	matriz[0][0] = 2
	matriz[1][1] = 5
	matriz[2][2] = 8
	fmt.Println(matriz)

	//slices son vectores dinámicos, se pueden ampliar o disminuir las dimensiones durante la ejecución
	slice := []int{2, 3, 6}
	fmt.Println(slice)

	variante2()
	variante3()
	variante4() //append
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]
	fmt.Println(porcion)
	porcion = elementos[2:4]
	fmt.Println(porcion)
	porcion = elementos[1:4]
	fmt.Println(porcion)
}

func variante3() {
	//lo que se necesita para crear un slice con un comando especial
	elementos := make([]int, 5, 20) /*5 es el largo, pero 20 es la capacidad,
	es decir, inicia con 5 elementos pero soporta hasta 20.
	Go reserva 20 posiciones en memoria, haciéndolo más dinámico
	Si le hago un len() me devuelve 5*/
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

//append recibe un slice y devuelve un slice
func variante4() {
	nums := make([]int, 0, 0) //acá tendremos que agregar posiciones obligaritamente
	for i := 0; i < 100; i++ {
		nums = append(nums, i) //a nums le agrega el número i y automáticamente se va redimensionando el slice
	}
	fmt.Printf("\n Con Append: Largo %d, Capacidad %d", len(nums), cap(nums))
	//go siempre pone/reserva una capacidad mayor, con potencias de 2 (binario)
}
