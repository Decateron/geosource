package field

type Checkboxes []Checkbox

type Checkbox struct {
	Label string `json:"label"`
	Value bool   `json:"value"`
}

func (checkboxes *Checkboxes) Validate() error {
	return nil
}

func (checkboxes *Checkboxes) IsEmpty() bool {
	return false
}
