package main

import "fmt"

func main() {
	/*
		Copia el minimo de elementos en ambos elementos
	*/
	slice := []int{1, 2, 3, 4, 5, 6}
	copia := make([]int, len(slice), cap(slice)*2)

	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
