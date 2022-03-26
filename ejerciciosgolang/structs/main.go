package main

import (
	"fmt"
	"time" //para manejar todo tipo de datos con fechas y horas

	"structs/user"
	/*
		los paquetes deben estar en la carpeta en que estoy parado que debe tener un archivo con el mismo nombre.go
	*/)

// en go no existe el POO (clases, métodos , propiedades, herencias, polimorfismo)
//para eso se basa en estructuras, con formas especiales de manejarlos

//para hacer herencia:
type pepe struct {
	user.Usuario
}

/*este main era para antes de usar otros packages

type usuario struct { //con minúscula para que se vea de forma privada y sea de acceso solo en este package
	Id        int
	Nombre    string
	FechaAlta time.Time //siempre va primero el nombre del paquete y luego del punto la función.
	//el segundo Time va en mayúscula por convensiones en Go
	Status bool
}

func main() {
	User := new(usuario)
	User.Id = 10
	User.Nombre = "Pablo"
	fmt.Println(User)
	//esto escribe:
	//&{10 Pablo 0001-01-01 00:00:00 +0000 UTC false}
	//el & es un puntero y la fecha no inicializada es la mínima posible

}*/

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Pablo", time.Now(), true)
	fmt.Println(u.Usuario)
}

/*El problema con ./user en Windows lo resolví como presento a continuación:



a. Con el cmd en la carpeta ejer09 usé la lógica de módulos: go mod init ejer09

b. Esto me crea un archivo go.mod  , sin esto no corre

c. En main, sección import cambié "./user" por "ejer09/user"

d. Para la sección en el struct pepe, cambié u.Usuario por user.Usuario, para que tuviera el mismo nombre del package que quiero usar.



Con esos cambios, me corrió el código correctamente sin enviar todas las carpetas y archivos al módulo raíz.

Espero que les sirva, saludos.*/
