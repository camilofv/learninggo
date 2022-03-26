package main

import (
	"fmt"
)

//mapas es una estructura que se puede serializar, como un vector, por ejemplo
//el índice de un mapa no es numérico, se puede crear con una clave y valor

func main() {
	paises := make(map[string]string) //la clave por la que ingresamos los ementos es string
	//el valor que se asigna a ese elemento es también string
	//guarda info de forma más amigable que arrays (vectores) y permite hacer cosas más complejas

	fmt.Println(paises)
	paises["Mexico"] = "D.F." //la clave es méxixo y el valor es su capital
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	/*paises := make(map[string]string, 5), reserva 5 espacios pero si me paso se guarda igual*/

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	fmt.Println(campeonato)
	//para adicionar neuvos elementos, se incorporan directamente, sin append:
	campeonato["River Plate"] = 25
	//esta asignación tb sirve para modificar un valor:
	campeonato["Chivas"] = 55
	//para eliminar:
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	//recorrer un mapa
	for equipo, puntaje := range campeonato { //se crean dos variables porque range devuelve ambos valores

		fmt.Printf("El equipo %s, tiene un puntaje de : %d \n", equipo, puntaje)

	}
	//si quiero saber si existe:
	puntaje, existe := campeonato["Mineiros"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe) // %t es apra boleanos
	puntaje, existe = campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe) // %t es apra boleanos
}
