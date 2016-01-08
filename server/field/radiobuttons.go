package field

type Radiobuttons []Radiobutton

type Radiobutton struct {
	Label string `json:"label"`
	Value bool   `json:"value"`
}

// Checks that at most one radiobutton is selected.
func (radiobuttons *Radiobuttons) Validate() error {
	if radiobuttons == nil {
		return ErrMissingValue
	}
	oneSelected := false
	for _, radiobutton := range *radiobuttons {
		if radiobutton.Value {
			if oneSelected {
				return ErrInvalidValue
			}
			oneSelected = true
		}
	}
	return nil
}

// Checks that at least one radiobutton has been selected.
func (radiobuttons *Radiobuttons) IsEmpty() bool {
	if radiobuttons == nil {
		return true
	}
	for _, radiobutton := range *radiobuttons {
		if radiobutton.Value {
			return true
		}
	}
	return false
}
