package main

import (
	"log"

	"gitlab.com/AnhellO/go-experiments/samples/web-applications/soap2/calculator"

	"github.com/hooklift/gowsdl/soap"
)

func main() {
	client := soap.NewClient("http://www.dneonline.com/calculator.asmx")
	service := calculator.NewCalculatorSoap(client)
	reply, err := service.Add(&calculator.Add{IntA: 4, IntB: 10})
	if err != nil {
		log.Fatalf("could't get sum from calculator: %v", err)
	}
	log.Println(reply)
}
