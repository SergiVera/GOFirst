package main

import (
	"fmt"
)

/*
	== igual a
	!= diferente de
	< menor que
	> mayor que
	>= mayor o igual que
	<= menor o igual que
	&& AND
	|| OR
*/

func main() {
	if true {
		fmt.Println("Hola Mundo!")
	}
	x := 14
	y := 14
	if x > y {
		fmt.Printf("%d es mayor que %d\n", x, y)
		return
	} else if x < y {
		fmt.Printf("%d es menor que %d\n", x, y)
		return
	}
	fmt.Printf("Son iguales")
}
