package main

import (
	"fmt"

	structs "gitlab.com/AnhellO/go-experiments/samples/oop/structs"
)

func main() {
	person := &structs.Person{"Vivan", 22}
	fmt.Println(person.Talk("Go is awesome"))
	fmt.Println(person.Walk(5.5))
	fmt.Println(person.Run(2.5))

	animal := &structs.Animal{"Lobo", "Canidos"}
	fmt.Println(animal.Walk(10.5))
	fmt.Println(animal.Run(12.5))
}
