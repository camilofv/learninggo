package main

import "fmt"

//interfases:definen conductas, operaciones, comportamientos de objetos en funciones

//en una intefase se definen los métodos que se usan para implementar esa interfase

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

//Género humano
type hombre struct { //ahora hombre es genérico, hombre y mujer
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

type mujer struct {
	hombre //con esto mujer hereda todas las definiciones y propiedades del hombre
}

func (h *hombre) respirar() { h.respirando = true } //la propiedad no puede tener el mismo nombre que la función
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

func HumanosRespirando(hu humano) { //recibe interfaz humano, con parámetro de estructura de cualquier humano
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

/*-----------------------------------------------------------------*/
//Género Animal
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
}

func (p *perro) respirar()         { p.respirando = true } //la propiedad no puede tener el mismo nombre que la función
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

func main() {
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)

	Maria := new(mujer) //el sexo de María inicia en false, así que no es necesario ponerlo
	HumanosRespirando(Maria)

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)
	fmt.Printf("Total carnívoros %d", totalCarnivoros)
}
