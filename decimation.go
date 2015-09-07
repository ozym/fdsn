package fdsn

import (
	"encoding/json"
)

// Corresponds to SEED blockette 57.
type Decimation struct {
	InputSampleRate Frequency
	Factor          int32
	Offset          int32
	Delay           Float
	Correction      Float
}

func (d *Decimation) String() string {

	j, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(j)
}

func (d *Decimation) IsValid() error {
	if d == nil {
		return nil
	}

	if err := Validate(&d.InputSampleRate); err != nil {
		return err
	}
	if err := Validate(&d.Delay); err != nil {
		return err
	}
	if err := Validate(&d.Correction); err != nil {
		return err
	}

	return nil
}
