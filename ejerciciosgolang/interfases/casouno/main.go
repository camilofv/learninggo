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
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

type mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

func (h *hombre) respirar()    { h.respirando = true } //la propiedad no puede tener el mismo nombre que la función
func (h *hombre) comer()       { h.comiendo = true }
func (h *hombre) pensar()      { h.pensando = true }
func (h *hombre) sexo() string { return "Hombre" }

func (m *mujer) respirar()    { m.respirando = true } //la propiedad no puede tener el mismo nombre que la función
func (m *mujer) comer()       { m.comiendo = true }
func (m *mujer) pensar()      { m.pensando = true }
func (m *mujer) sexo() string { return "Mujer" }

func HumanosRespirando(hu humano) { //recibe interfaz humano, con parámetro de estructura de cualquier humano
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

func main() {
	Pedro := new(hombre)
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)
}
