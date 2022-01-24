package main

import "fmt"

func main() {
	nombre := "Coco"
	fmt.Println(nombre)

	var a1 int
	var b1, c1 int
	a1 = 2
	b1 = -15
	fmt.Println(a1, b1, c1)

	var a2 int = 2
	var b2, c2 int = -15, 0
	var d2 bool = true
	var s2 string = "hola"
	fmt.Println(a2, b2, c2, d2, s2)

	var a3 = 2
	var b3, c3 = -15, 0
	var d3 = true
	var s3 = "hola"
	fmt.Println(a3, b3, c3, d3, s3)

	a4 := 2
	b4, c4 := 15, 0
	d4 := true
	s4 := "hola"
	fmt.Println(a4, b4, c4, d4, s4)
}
