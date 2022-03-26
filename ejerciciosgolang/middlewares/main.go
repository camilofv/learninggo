package main

import "fmt"

//los middlewares son interceptores, es decir, permiten encapsular la ejecución de una función preexistente con una función nueva
/*son interceptores que permiten ejecutar instrucciones comunes a varias funciones
que reciben y devuelven los mismos tipos de variables*/

var result int

func main() {
	//los middlewares sirven para crear capas de funcionalidad, seguridad, encriptación
	//hacen breves modificaciones con amplia funcionalidad
	fmt.Println("Inicio")

	result = operacionesMidd(sumar)(2, 3) //hace un llamado a la función middleware, con la función real que debe ejecutar y los parámetors que debe ejecutar
	fmt.Println(result)
	result = operacionesMidd(restar)(10, 6)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 4)
	fmt.Println(result)

}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

//middleware recibe primero una función con los tipos de los datos que tiene, no con sus nombres
//luego retorna la misma función después de ejecutar el middleware, y lo que recibe y lo que devuelve deben ser exactamente del mismo tipo
func operacionesMidd(f func(int, int) int) func(int, int) int { //esto opera sobre varias funciones
	//lo que recibe y lo que devuelve debe ser siempre del mismo tipo que las otras funciones
	return func(a, b int) int { //los tipos de acá también deben coincidir con las entradas y salidas de la función
		fmt.Println("Inicio de operación") //acá se integra una nueva capa de procesamiento, antes de ejecutar el la función sumar, restar o multiplicar
		return f(a, b)                     //acá devuelve la función antigua y sus parámetros para ejecutarla donde corresponde
	}
}
