package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	precio := 14.3
	fmt.Printf("El precio es %f\n", precio)

	name := "Uriel"
	fmt.Printf("El nombre es %v\n", name)

	var edad int
	var nombre string
	fmt.Println("Coloca tu nombre: ")
	fmt.Scanf("%s\n", &nombre)
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi nombre es %s y tengo %d\n", nombre, edad)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Coloca tu nombre: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Hola " + text)
}
