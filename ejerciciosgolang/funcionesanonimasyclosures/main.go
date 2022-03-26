package main

import (
	"fmt"
)

/*las funciones anónimas no tienen nombres
sirven para crear operaciones donde no voy a necesitar una función
se puede isolar su código (se oculta dentro de una variable)
y se puede modificar en tiempo de ejecución*/

var calculo func(int, int) int = func(num1 int, num2 int) int { //función también es un tipo de dato
	return num1 + num2
}

func main() {

	fmt.Printf("Suma 5+7 = %d \n", calculo(5, 7))

	//restamos redefiniendo la función cáculo, modificándola y cambiando su resultado
	calculo = func(num1 int, num2 int) int { /*esto es lo más parecido a una sobrecarga de funciones,
		pero no permite cambiar la cantidad y tipo de los parámetros de entrada*/
		return num1 - num2
	}
	fmt.Printf("Resta 5-7 = %d \n", calculo(5, 7))

	operaciones()

	//closure

	tablaDel := 2
	fmt.Printf("La tabla del %d es: \n", tablaDel)
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

//otro ejemplo de funciones anónimas
func operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Printf("La suma de 23 + 13 es %d \n", resultado())

}

/*closure: concepto de funciones anónimas, aumenta la seguridad por la isolación de código, haciendo que
no todos puedan ingresar a nuestras variables o funciones
Los closures pueden acceder a variables creadas por fuera de la función, variables creadas en la función original*/

func Tabla(valor int) func() int { //función que devuelve otra función que a su vez devuelve un valor int
	numero := valor
	secuencia := 0
	return func() int { //la función principal retorna el valor entregado por la función interna
		//Tabla solo toma esta parte de la operación, no lo anterior, y es la gracia de los closures
		//MI tabla tomó solo la función que le devolvió Tabla
		secuencia += 1
		return numero * secuencia
	}
}
