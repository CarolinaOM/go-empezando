package main

import "fmt"

type language int

const (
	EN language = iota
	FR
	ES
	DE
)

func main() {
	// var a [3]int
	// var b [3]string

	// fmt.Println(a)
	// fmt.Printf("%q\n", b)

	// var ages [3]int = [3]int{10, 20, 30}
	// fmt.Println(ages)

	// var languages = [3]string{"go", "php", "java"}
	// fmt.Println(languages)

	// var ages = [...]int{10, 20, 30}
	// fmt.Println(ages)

	// ages := [3]int {10, 20, 30}
	// ages = [4]int {10, 20, 30, 40}

	// var length int
	// fmt.Println("Ingresa la longitud del arreglo: ")
	// fmt.Scanf("%d", &length)
	// var ages [length]int
	// fmt.Println(ages )

	// symbol := [...]string{
	// 	EN: "English",
	// 	FR: "French",
	// 	ES: "Spanish",
	// 	DE: "German",
	// }
	// fmt.Println(ES, symbol[ES])

	// ages := [...]int{10: 12}
	// fmt.Println(ages)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
}
