package main

//How to declare multiple imports
import (
	"fmt"
	"strconv"
)

func main() {
	//Convert int to string
	edad := 22

	edadStr := strconv.Itoa(edad)

	fmt.Println("Mi edad es: " + edadStr)

	//Convert string to int
	altura := "170"

	alturaInt, _ := strconv.Atoi(altura)

	fmt.Println(alturaInt + 10)
}
