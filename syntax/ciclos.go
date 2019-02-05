package main

import (
	"fmt"
)

func main() {
	//Simulación de ciclo while
	i := 0
	for {
		fmt.Printf("Hola Mundo while %d!\n", i)
		i++
		if i > 10 {
			break
		}
	}

	//Con ciclo for
	for i := 0; i < 10; i++ {
		fmt.Printf("Hola Mundo con for %d!\n", i)
	}
}
