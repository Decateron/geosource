package main

import (
	"encoding/json"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	data := `
{
	"type": "checkboxes",
	"label": "bar",
	"required": true,
	"form": [
		"foo",
		"bar"
	],
	"value": [
		true,
		false
	]
}
`
	field, err := UnmarshalField([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	err = field.Validate()
	if err != nil {
		log.Fatal(err)
	}
	blob, err := json.Marshal(field)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(blob))
}
