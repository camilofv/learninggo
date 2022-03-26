package main

//ejercicios de manejos de errores y excepciones: defer, panic and recover
//en go no existen los try catch ni exceptions
import (
	"fmt"
	"log"
	"os"
	//funciones permite grabar en el log una bitácora, un paso a paso en el sistema, sirve para depurar
)

//Defer es una función que se ejecuta sí o sí cuando se detecta que una función se va por un return, opor un error o por un fin de función
//Panic forza un error,el sistema va a abortar y muestra un mensaje
//Recover es para tener el control sobre un panic, se ejecuta cuando se detecta un panic

func main() {

	//ejemploDefer()
	ejemploPanic()

}

func ejemploDefer() {
	archivo := "prueba.txt" //este archivo no existe, así que debería dar error
	f, err := os.Open(archivo)

	defer f.Close() //si una instrucción está antecedida por defer, esta instrucción no se ejecuta secuencialmente y sigue a la siguiente instrucción
	//se va a ejecutar cuando salga de la función

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1) //acá el exit cierra todo, pero antes se ejecuta el defer
	}
}

func ejemploPanic() {
	defer func() {
		//esta función anónima es el recover, se ejecuta antes del aborto por panic
		// es anónima porque el defer ejecuta una sola cosa, no le puedo dar un listado de instrucciones, pero con una función sí puedo hacerlo
		reco := recover() //si no hubo panic, reco será nil
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó Panic \n %v", reco) // %v es un valor que es variable, porque reco no es un texto, es un objeto
			//Fatalf es lo mismo que un Printf, solo que graba directamente en el archivo de log
			//Fatalf es mejor que Printf porque muestra fecha, hora y ejecuta un os.Exit(1)
		}
	}() //ojo, es una función dentro de una función

	a := 1
	if a == 1 {
		panic("Se encontró el valor de 1") //al entrar en panic, se aborta el programa
		//sirve cuandopor validación  falta algún dato crucial y que no puede continuar el sistema sin ese dato,
		//o cuando se detecta un fallo en el hardware, entre otras cosas.
	}
}
