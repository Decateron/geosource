package main

const (
	TEXT = iota
	IMAGE = iota
	AUDIO = iota
	VIDEO = iota
	CHECKBOXES = iota
	RADIO_BUTTONS = iota
)

var types = []string {
	"TEXT",
	"IMAGE",
	"AUDIO",
	"VIDEO",
	"CHECK_BOXES",
	"RADIO_BUTTONS",
}

type Field struct {
	Type  string      `json:"type"`
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}