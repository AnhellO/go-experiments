package main

import (
	"fmt"
	"log"

	"github.com/AnhellO/go-experiments/samples/parsing/json/config"
)

func main() {
	conf := config.NewConfig("sample.json")
	val, err := conf.Load("X")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", val)
}
