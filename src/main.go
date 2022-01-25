package main

import "fmt"

func main() {
	greeting := "Hello There"
	fmt.Println(greeting)

	var (
		nombre string
		edad   int
		pi     float64
	)
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edad)

	fmt.Print("Ingrese el valor de pi: ")
	fmt.Scanln(&pi)

	fmt.Printf("Nombre: %s Edad: %d \nvalor de Pi: %f", nombre, edad, pi)
	fmt.Println()
}
