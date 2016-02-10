package main

import (
	"encoding/json"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	channelData := `
{
	"name": "test",
	"fields": [
		{
			"type": "checkboxes",
			"label": "bar",
			"required": true,
			"form": [
				"foo",
				"bar"
			]
		},
		{
			"type": "text",
			"label": "baz",
			"required": true
		}
	]
}`

	channel, err := UnmarshalChannel([]byte(channelData))
	if err != nil {
		log.Fatal(err)
	}

	submissionData := `
{
	"title": "hello, world!",
	"channel": "test",
	"values": [[true, false], "hello"]
}`

	post, err := channel.UnmarshalSubmission([]byte(submissionData))
	if err != nil {
		log.Fatal(err)
	}

	blob, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(blob))
}
