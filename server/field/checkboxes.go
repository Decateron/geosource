package field

type Checkboxes []Checkbox

type Checkbox struct {
	Label string `json:"label"`
	Value bool   `json:"value"`
}

func (checkboxes *Checkboxes) Validate() error {
	if checkboxes == nil {
		return ErrMissingValue
	}
	return nil
}

func (checkboxes *Checkboxes) IsEmpty() bool {
	if checkboxes == nil {
		return true
	}
	return false
}
