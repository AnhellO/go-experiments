package main

import (
	"fmt"

	calculator "gitlab.com/AnhellO/go-experiments/samples/web-applications/soap/calculator"

	"github.com/fiorix/wsdl2go/soap"
)

func main() {
	client := soap.Client{
		URL:         "http://www.dneonline.com/calculator.asmx",
		Namespace:   calculator.Namespace,
		Header:      "SOAPAction: \"http://tempuri.org/Add\"",
		ContentType: "application/soap+xml; charset=utf-8;",
	}

	soapService := calculator.NewCalculatorSoap(&client)
	response, err := soapService.Add(&calculator.Add{IntA: 5, IntB: 5})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client)
	fmt.Printf("%v\n", soapService)
	fmt.Println(response)
}
