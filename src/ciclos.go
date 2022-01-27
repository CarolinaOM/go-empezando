package main

import "fmt"

func main() {
	i := 0

	for {
		if i == 2 {
			i++
			continue
		}

		fmt.Println(i)
		i++

		if i > 10 {
			break
		}
	}
}
