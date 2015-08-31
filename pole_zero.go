package fdsn

import (
	"fmt"
)

// Complex numbers used as poles or zeros in channel response.
type PoleZero struct {
	Number    int32 `xml:"number,attr"`
	Real      FloatNoUnit
	Imaginary FloatNoUnit
}

func (pz PoleZero) IsValid() error {
	if pz.Number < 0 {
		return fmt.Errorf("invalid pole/zero number: %d", pz.Number)
	}
	if err := pz.Real.IsValid(); err != nil {
		return err
	}
	if err := pz.Imaginary.IsValid(); err != nil {
		return err
	}

	return nil
}
