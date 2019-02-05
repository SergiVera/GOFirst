package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	copia := make([]int, len(slice), cap(slice)*2)

	//Copia el mínimo de elementos en ambos arreglos
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
