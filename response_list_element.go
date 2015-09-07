package fdsn

import (
	"encoding/json"
)

type ResponseListElement struct {
	Frequency Frequency
	Amplitude Float
	Phase     Angle
}

func (r *ResponseListElement) String() string {

	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (r *ResponseListElement) IsValid() error {

	if r == nil {
		return nil
	}

	if err := Validate(&r.Frequency); err != nil {
		return err
	}
	if err := Validate(&r.Amplitude); err != nil {
		return err
	}
	if err := Validate(&r.Phase); err != nil {
		return err
	}

	return nil
}
