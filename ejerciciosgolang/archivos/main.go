package main

import (
	"bufio"
	"fmt" //para pedir por teclado
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt") //archivo es variable que contiene el contenido del archivo.txt
	//ReadFile se encarga de abrir el archivo, leer su contenido y cerrarlo
	if err != nil {
		fmt.Println("Hubo un error")

	} else {
		fmt.Println(string(archivo)) //convierte al contenido de archivo en string
	}
}

//variante 2: abre el archivo y lo deja listo para lectura
//con esta variante le puedo dar formatos o tratamientos a lo que está leyendo
func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt") //archivo es variable que contiene el contenido del archivo.txt

	if err != nil {
		fmt.Println("Hubo un error")

	} else {
		scanner := bufio.NewScanner(archivo) //scanner lee el archivo
		for scanner.Scan() {                 //el for lee una línea, si es válida ejecuta lo que está dentro del bucle
			//Lee hasta que sea false o End Of File
			registro := scanner.Text() //convierte la línea en memoria en una cadena de texto
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	archivo.Close() //acá sí debo cerrar el archivo
}

//en esta variante, se escribe el archivo sobre el archivo anterior
func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt") // Crea archivo
	//si lo ejecuto varias veces, sobreescribe el archivo anterior
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}

	fmt.Fprintln(archivo, "Soy una nueva líneaa") //F de println es que se graba en un archivo
	archivo.Close()
	fmt.Println("Se ha creado nuevoArchivo.txt")
}

//en esta variante puedo adicionar nuevas líneas al archivo anterior
func graboArchivo2() {
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\n Esta es una nueva línea desde graboArchivo2()") == false { //append es el que incorpora nuevas líneas
		fmt.Println("Error para integrar la nueva línea")
	}
	fmt.Println("Se ha ingresado nueva línea a nuevoArchivo.txt")
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644) //WRONLY es Write and Read Only, la | espara ingresar más de un parámetro
	//O_Append es para no blanquear el archivo, sino que lo agrega al final
	//0644 son permisos de usuario, para dueño, usuario, grupo, otros
	if err != nil {
		fmt.Println("Hubo un error")
		return false //si retorna false, en la func graboArchivo2 escribe el "Error para integrar la nueva línea"
	}

	_, err = arch.WriteString(texto) //no nos interesa lo que devuelve en el primer parámetro, por eso se pone _
	if err != nil {
		fmt.Println("Hubo un error 2")
		return false
	}
	return true
}
