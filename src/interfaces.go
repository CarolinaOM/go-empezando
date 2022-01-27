package main

import "fmt"

type User interface {
	Permisos() int // 1 - 5
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.nombre

}

func auth(user User) string {
	if user.Permisos() > 4 {
		return user.Nombre() + "Tiene permisos de administrador"
	}
	return ""
}

func main() {
	admin := Admin{"Uriel"}
	fmt.Println(auth(admin))
}
