package fdsn

import (
	"encoding/json"
)

type ResponseListElement struct {
	Frequency Frequency
	Amplitude Float
	Phase     Angle
}

func (r ResponseListElement) String() string {

	j, err := json.Marshal(&r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (r ResponseListElement) IsValid() error {
	if err := r.Frequency.IsValid(); err != nil {
		return err
	}
	if err := r.Amplitude.IsValid(); err != nil {
		return err
	}
	if err := r.Phase.IsValid(); err != nil {
		return err
	}
	return nil
}
