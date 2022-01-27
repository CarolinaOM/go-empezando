package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	archivo, error := os.Open("./hola.txt")

	if error != nil {
		fmt.Println("Existe un error")
	}
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}
}
