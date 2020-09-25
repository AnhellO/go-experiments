package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Data struct {
	XMLName    xml.Name `xml:"data" json:"-"`
	PersonList []Person `xml:"person" json:"people"`
}

type Person struct {
	FirstName string   `xml:"firstname" json:"firstname"`
	LastName  string   `xml:"lastname" json:"lastname"`
	Address   *Address `xml:"address" json:"address,omitempty"`
}

type Address struct {
	City  string `xml:"city" json:"city,omitempty"`
	State string `xml:"state" json:"state,omitempty"`
}

func main() {
	var data Data

	// Convert from XML to JSON
	rawXMLData := getRawXML()
	xml.Unmarshal([]byte(rawXMLData), &data)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))

	// Convert from JSON to XML
	rawJSONData := getRawJSON()
	json.Unmarshal([]byte(rawJSONData), &data)
	xmlData, _ := xml.Marshal(data)
	fmt.Println(string(xmlData))
}

func getRawXML() string {
	return `
		<data>
			<person>
				<firstname>Angel</firstname>
				<lastname>Jaime</lastname>
				<address>
					<city>Saltillo</city>
					<state>COAH</state>
				</address>
			</person>
			<person>
				<firstname>Ana</firstname>
				<lastname>De Armas</lastname>
			</person>
		</data>
	`
}

func getRawJSON() string {
	return `
	{
		"people": [
			{
				"firstname": "Angel",
				"lastname": "Jaime",
				"address": {
					"city": "Saltillo",
					"state": "COAH"
				}
			},
			{
				"firstname": "Ana",
				"lastname": "De Armas"
			}
		]
	}
`
}
