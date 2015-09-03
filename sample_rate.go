package fdsn

import (
	"encoding/json"
	"fmt"
)

// Sample rate in samples per second.
type SampleRate struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // SAMPLES/S

	Value float64 `xml:",chardata"`
}

func (s SampleRate) String() string {

	j, err := json.Marshal(&s)
	if err != nil {
		return ""
	}
	return string(j)
}

func (s SampleRate) IsValid() error {
	if s.Unit != "" && s.Unit != "SAMPLES/S" {
		return fmt.Errorf("invalid sample rate unit: %s", s.Unit)
	}
	return nil
}
