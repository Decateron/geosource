package field

type Number float64

func (number *Number) Validate() error {
	return nil
}

func (number *Number) IsEmpty() bool {
	return number == nil
}
