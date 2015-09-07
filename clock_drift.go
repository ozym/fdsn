package fdsn

import (
	"encoding/json"
	"fmt"
)

// A tolerance value, measured in seconds per sample, used as a
// threshold for time error detection in data from the channel.
type ClockDrift struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // SECONDS/SAMPLE

	Value float64 `xml:",chardata"`
}

func (c *ClockDrift) String() string {

	j, err := json.Marshal(&c)
	if err != nil {
		return ""
	}
	return string(j)
}

func (c *ClockDrift) IsValid() error {
	if c == nil {
		return nil
	}

	if c.Unit != "" && c.Unit != "SECONDS/SAMPLE" {
		return fmt.Errorf("invalid clock drift unit: %s", c.Unit)
	}

	return nil
}
