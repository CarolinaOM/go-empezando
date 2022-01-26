package main

import "fmt"

func main() {
	// matriz := []int{1, 2}
	// if matriz == nil {
	// 	fmt.Println("Esta vacio")
	// } else {
	// 	fmt.Println(len(matriz))
	// }

	// // Puntero al arreglo
	// // Longitud del arreglo al que apunta
	// // Capacidad

	arreglo := [3]int{1, 2, 3}
	slice := arreglo[0:3]
	fmt.Println(slice)

}
