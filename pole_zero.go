package fdsn

import (
	"encoding/json"
	"fmt"
)

// Complex numbers used as poles or zeros in channel response.
type PoleZero struct {
	Number    int32 `xml:"number,attr"`
	Real      FloatNoUnit
	Imaginary FloatNoUnit
}

func (pz *PoleZero) String() string {

	j, err := json.Marshal(pz)
	if err != nil {
		return ""
	}
	return string(j)
}

func (pz *PoleZero) IsValid() error {

	if pz == nil {
		return nil
	}

	if pz.Number < 0 {
		return fmt.Errorf("invalid pole/zero number: %d", pz.Number)
	}
	if err := Validate(&pz.Real); err != nil {
		return err
	}
	if err := Validate(&pz.Imaginary); err != nil {
		return err
	}

	return nil
}
