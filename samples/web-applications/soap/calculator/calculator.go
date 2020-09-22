// Code generated by wsdl2go. DO NOT EDIT.

package calculatorsoap12

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://tempuri.org/"

// NewCalculatorSoap creates an initializes a CalculatorSoap.
func NewCalculatorSoap(cli *soap.Client) CalculatorSoap {
	return &calculatorSoap{cli}
}

// CalculatorSoap was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type CalculatorSoap interface {
	// Adds two integers. This is a test WebService. ©DNE Online
	Add(Add *Add) (*AddResponse, error)

	// Divide was auto-generated from WSDL.
	Divide(Divide *Divide) (*DivideResponse, error)

	// Multiply was auto-generated from WSDL.
	Multiply(Multiply *Multiply) (*MultiplyResponse, error)

	// Subtract was auto-generated from WSDL.
	Subtract(Subtract *Subtract) (*SubtractResponse, error)
}

// Add was auto-generated from WSDL.
type Add struct {
	IntA int `xml:"intA" json:"intA" yaml:"intA"`
	IntB int `xml:"intB" json:"intB" yaml:"intB"`
}

// AddResponse was auto-generated from WSDL.
type AddResponse struct {
	AddResult int `xml:"AddResult" json:"AddResult" yaml:"AddResult"`
}

// Divide was auto-generated from WSDL.
type Divide struct {
	IntA int `xml:"intA" json:"intA" yaml:"intA"`
	IntB int `xml:"intB" json:"intB" yaml:"intB"`
}

// DivideResponse was auto-generated from WSDL.
type DivideResponse struct {
	DivideResult int `xml:"DivideResult" json:"DivideResult" yaml:"DivideResult"`
}

// Multiply was auto-generated from WSDL.
type Multiply struct {
	IntA int `xml:"intA" json:"intA" yaml:"intA"`
	IntB int `xml:"intB" json:"intB" yaml:"intB"`
}

// MultiplyResponse was auto-generated from WSDL.
type MultiplyResponse struct {
	MultiplyResult int `xml:"MultiplyResult" json:"MultiplyResult" yaml:"MultiplyResult"`
}

// Subtract was auto-generated from WSDL.
type Subtract struct {
	IntA int `xml:"intA" json:"intA" yaml:"intA"`
	IntB int `xml:"intB" json:"intB" yaml:"intB"`
}

// SubtractResponse was auto-generated from WSDL.
type SubtractResponse struct {
	SubtractResult int `xml:"SubtractResult" json:"SubtractResult" yaml:"SubtractResult"`
}

// Operation wrapper for Add.
// OperationAddSoapIn was auto-generated from WSDL.
type OperationAddSoapIn struct {
	Add *Add `xml:"Add,omitempty" json:"Add,omitempty" yaml:"Add,omitempty"`
}

// Operation wrapper for Add.
// OperationAddSoapOut was auto-generated from WSDL.
type OperationAddSoapOut struct {
	AddResponse *AddResponse `xml:"AddResponse,omitempty" json:"AddResponse,omitempty" yaml:"AddResponse,omitempty"`
}

// Operation wrapper for Divide.
// OperationDivideSoapIn was auto-generated from WSDL.
type OperationDivideSoapIn struct {
	Divide *Divide `xml:"Divide,omitempty" json:"Divide,omitempty" yaml:"Divide,omitempty"`
}

// Operation wrapper for Divide.
// OperationDivideSoapOut was auto-generated from WSDL.
type OperationDivideSoapOut struct {
	DivideResponse *DivideResponse `xml:"DivideResponse,omitempty" json:"DivideResponse,omitempty" yaml:"DivideResponse,omitempty"`
}

// Operation wrapper for Multiply.
// OperationMultiplySoapIn was auto-generated from WSDL.
type OperationMultiplySoapIn struct {
	Multiply *Multiply `xml:"Multiply,omitempty" json:"Multiply,omitempty" yaml:"Multiply,omitempty"`
}

// Operation wrapper for Multiply.
// OperationMultiplySoapOut was auto-generated from WSDL.
type OperationMultiplySoapOut struct {
	MultiplyResponse *MultiplyResponse `xml:"MultiplyResponse,omitempty" json:"MultiplyResponse,omitempty" yaml:"MultiplyResponse,omitempty"`
}

// Operation wrapper for Subtract.
// OperationSubtractSoapIn was auto-generated from WSDL.
type OperationSubtractSoapIn struct {
	Subtract *Subtract `xml:"Subtract,omitempty" json:"Subtract,omitempty" yaml:"Subtract,omitempty"`
}

// Operation wrapper for Subtract.
// OperationSubtractSoapOut was auto-generated from WSDL.
type OperationSubtractSoapOut struct {
	SubtractResponse *SubtractResponse `xml:"SubtractResponse,omitempty" json:"SubtractResponse,omitempty" yaml:"SubtractResponse,omitempty"`
}

// calculatorSoap implements the CalculatorSoap interface.
type calculatorSoap struct {
	cli *soap.Client
}

// Adds two integers. This is a test WebService. ©DNE Online
func (p *calculatorSoap) Add(Add *Add) (*AddResponse, error) {
	α := struct {
		OperationAddSoapIn `xml:"tns:Add"`
	}{
		OperationAddSoapIn{
			Add,
		},
	}

	γ := struct {
		OperationAddSoapOut `xml:"AddResponse"`
	}{}
	if err := p.cli.RoundTripSoap12("http://tempuri.org/Add", α, &γ); err != nil {
		return nil, err
	}
	return γ.AddResponse, nil
}

// Divide was auto-generated from WSDL.
func (p *calculatorSoap) Divide(Divide *Divide) (*DivideResponse, error) {
	α := struct {
		OperationDivideSoapIn `xml:"tns:Divide"`
	}{
		OperationDivideSoapIn{
			Divide,
		},
	}

	γ := struct {
		OperationDivideSoapOut `xml:"DivideResponse"`
	}{}
	if err := p.cli.RoundTripSoap12("http://tempuri.org/Divide", α, &γ); err != nil {
		return nil, err
	}
	return γ.DivideResponse, nil
}

// Multiply was auto-generated from WSDL.
func (p *calculatorSoap) Multiply(Multiply *Multiply) (*MultiplyResponse, error) {
	α := struct {
		OperationMultiplySoapIn `xml:"tns:Multiply"`
	}{
		OperationMultiplySoapIn{
			Multiply,
		},
	}

	γ := struct {
		OperationMultiplySoapOut `xml:"MultiplyResponse"`
	}{}
	if err := p.cli.RoundTripSoap12("http://tempuri.org/Multiply", α, &γ); err != nil {
		return nil, err
	}
	return γ.MultiplyResponse, nil
}

// Subtract was auto-generated from WSDL.
func (p *calculatorSoap) Subtract(Subtract *Subtract) (*SubtractResponse, error) {
	α := struct {
		OperationSubtractSoapIn `xml:"tns:Subtract"`
	}{
		OperationSubtractSoapIn{
			Subtract,
		},
	}

	γ := struct {
		OperationSubtractSoapOut `xml:"SubtractResponse"`
	}{}
	if err := p.cli.RoundTripSoap12("http://tempuri.org/Subtract", α, &γ); err != nil {
		return nil, err
	}
	return γ.SubtractResponse, nil
}
