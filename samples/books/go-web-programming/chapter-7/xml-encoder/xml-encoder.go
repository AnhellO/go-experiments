package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Person represents a <person> node in the XML
type Person struct {
	XMLName    xml.Name `xml:"person"`
	FirstName  string   `xml:"firstName"`
	MiddleName string   `xml:"middleName"`
	LastName   string   `xml:"lastName"`
	Age        int64    `xml:"age"`
	Skills     []Skill  `xml:"skills"`
}

// Skill represents a <skill> node in the XML
type Skill struct {
	XMLName        xml.Name `xml:"skill"`
	Name           string   `xml:"skillName"`
	YearsPracticed int64    `xml:"practice"`
}

func main() {
	person := Person{
		FirstName:  "Bob",
		MiddleName: "",
		LastName:   "Jones",
		Age:        23,
		Skills: []Skill{
			{
				Name:           "Cooking",
				YearsPracticed: 3,
			},
			{
				Name:           "Basketball",
				YearsPracticed: 4,
			},
		},
	}

	xmlFile, err := os.Create("person.xml")
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return
	}
	xmlFile.WriteString(xml.Header)
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&person)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return
	}
}
