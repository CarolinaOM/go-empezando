package main

import (
	"fmt"
)

func main() {
	carrera1 := new(Carrera)
	carrera1.nombreCarrera = "Sistemas"
	carrera1.duracion = 5
	carrera1.nombre = "LP1"
	carrera1.url = "SistemasLP1"
	carrera1.habilidad = []string{"Programacion", "Redes"}

	carrera1.Inscribirse("Caro")
	fmt.Println(carrera1)
}

// Curso es una estructura
type Curso struct {
	nombre    string
	url       string
	habilidad []string
}

// Inscribirse es un metodo de la estrutura curso
func (c Curso) Inscribirse(nombre string) {
	fmt.Printf("La persona %s se ha inscito al curso %s \n", nombre, c.nombre)
}

// Carrera nueva estructura
type Carrera struct {
	nombreCarrera string
	duracion      int
	Curso
}
