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

	fmt.Printf("mi nombre es: %s", nombre)
	fmt.Println()

	// Metodo if

	if false {
		fmt.Println("Se cumple la condicion")
	} else {
		fmt.Println("No se cumple la condicion")
	}

	var n int
	fmt.Print("Ingrese el numero: ")
	fmt.Scanln(&n)
	if n == 0 {
		fmt.Println(n, "Es neutro")
	} else if n%2 == 0 {
		fmt.Println(n, "Es par")
	} else {
		fmt.Println(n, "Es impar")
	}
}
