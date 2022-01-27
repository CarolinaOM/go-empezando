package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func main() {
	// fmt.Println(User{nombre: "Uriel", apellido: "Lopez"})
	uriel := new(User)
	uriel.nombre = "Uriel"
	fmt.Println(uriel.nombre)
}
