package fdsn

import (
	"encoding/json"
)

// Sample rate expressed as number of samples in a number of seconds.
type SampleRateRatio struct {
	NumberSamples int32
	NumberSeconds int32
}

func (s SampleRateRatio) String() string {

	j, err := json.Marshal(&s)
	if err != nil {
		return ""
	}
	return string(j)
}

func (s SampleRateRatio) IsValid() error {
	return nil
}
