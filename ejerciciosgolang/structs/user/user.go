package user

//la mejor práctica es poner el nombre del package igual al nombre del archivo
//no se pueden repetir los nombres de los packages

import "time"

//definición de cómo crear una estructura:
type Usuario struct { //con mayúscula para que se vea de forma pública y sea de acceso de otros packages
	Id        int
	Nombre    string
	FechaAlta time.Time //siempre va primero el nombre del paquete y luego del punto la función.
	//el segundo Time va en mayúscula por convensiones en Go
	Status bool
}

//función de asignación de valores que no devuelve valores (es un método en otros lenguajes)

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	//esta función recibe parámetros cuando se haga referencia a la estructura
	//*Usuario es un puntero
	//cada vez que se use this acá, apuntará como referencia al objeto en que estoy parado, en este caso Usuario
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
