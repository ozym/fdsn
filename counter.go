package fdsn

// Integers greater than or equal to 0
type Counter uint32

func (c Counter) IsValid() error {
	return nil
}
