package structs

import "fmt"

type Actions interface {
	Walk(distance float64) string
	Run(distance float64) string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Talk(phrase string) string {
	return fmt.Sprintf("%s is %d and says: %s", p.Name, p.Age, phrase)
}

func (p *Person) Walk(distance float64) string {
	return fmt.Sprintf("%s is %d and walks %.2fkm", p.Name, p.Age, distance)
}

func (p *Person) Run(distance float64) string {
	return fmt.Sprintf("%s is %d and runs %.2fkm", p.Name, p.Age, distance)
}

type Animal struct {
	Name    string
	Species string
}

func (a *Animal) Walk(distance float64) string {
	return fmt.Sprintf("The %s is a %s and walks %.2fkm", a.Name, a.Species, distance)
}

func (a *Animal) Run(distance float64) string {
	return fmt.Sprintf("The %s is a %s and runs %.2fkm", a.Name, a.Species, distance)
}
