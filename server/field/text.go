package field

type Text string

func (text *Text) Validate() error {
	if text == nil {
		return ErrMissingValue
	}
	return nil
}

func (text *Text) IsEmpty() bool {
	return *text == ""
}
