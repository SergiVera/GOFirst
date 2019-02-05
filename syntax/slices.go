package main

import (
	"fmt"
)

func main() {
	matriz := []int{1, 2}
	if matriz == nil {
		fmt.Println("Está vacío")
		return
	}
	fmt.Printf("Longitud: %d\n", len(matriz))
	fmt.Println(matriz)
	arreglo := [3]int{1, 2, 3}
	slice := arreglo[:2]
	slice = arreglo[:1]
	fmt.Println(slice)
}
