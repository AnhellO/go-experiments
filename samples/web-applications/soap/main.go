package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	calculator "gitlab.com/AnhellO/go-experiments/samples/web-applications/soap/calculator"

	"github.com/fiorix/wsdl2go/soap"
)

func main() {
	client := soap.Client{
		URL:           "http://www.dneonline.com/calculator.asmx",
		Namespace:     calculator.Namespace,
		URNamespace:   calculator.Namespace,
		ThisNamespace: calculator.Namespace,
		Envelope:      "http://www.w3.org/2003/05/soap-envelope",
		Pre: func(r *http.Request) {
			buf := new(strings.Builder)
			_, err := io.Copy(buf, r.Body)

			if err != nil {
				fmt.Errorf("Error: %v", err)
			}

			updatedBody := strings.Replace(buf.String(), "<Add>", "<Add xmlns=\"http://tempuri.org/\">", -1)
			r.Body = ioutil.NopCloser(strings.NewReader(updatedBody))
			r.ContentLength = int64(len(updatedBody))
		},
	}

	soapService := calculator.NewCalculatorSoap(&client)
	response, err := soapService.Add(&calculator.Add{IntA: 5, IntB: 5})

	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	fmt.Println(*response)
}
