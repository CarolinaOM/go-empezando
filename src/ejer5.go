package main

import (
	"fmt"
)

type User struct {
	nombre     string
	edad       int
	ndocumento int64
}

func (this *User) set_name(n string) {
	this.nombre = n
}
func (this *User) set_edad(e int) {
	this.edad = e
}
func (this *User) set_documento(d int64) {
	this.ndocumento = d
}

func main() {
	uriel := new(User)
	uriel.set_name("Marcos")
	fmt.Println(uriel.nombre)
	uriel.set_edad(30)
	fmt.Println(uriel.edad)
	uriel.set_documento(02154555)
	fmt.Println(uriel.ndocumento)
}
