package fdsn

import (
	"fmt"
)

// Response: expressed as a polynomial (allows non-linear sensors to be described).
// Corresponds to SEED blockette 62. Can be used to describe a stage of acquisition or a complete system.
type Polynomial struct {
	// Same meaning as Equipment:resourceId.</xs:documentation>
	ResourceId string `xml:"resourceId,attr"`

	// A name given to this filter.
	Name string `xml:"name,attr"`

	Description string `xml:",omitempty" json:",omitempty"`

	// The units of the data as input from the perspective of data acquisition.
	// After correcting data for this response, these would be the resulting units.
	InputUnits Units

	// The units of the data as output from the perspective of data acquisition.
	// These would be the units of the data prior to correcting for this response.
	OutputUnits Units

	ApproximationType       ApproximationType
	FrequencyLowerBound     Frequency
	FrequencyUpperBound     Frequency
	ApproximationLowerBound string
	ApproximationUpperBound string
	MaximumError            float64

	Coefficients []Coefficient `xml:"Coefficient,omitempty" json:",omitempty"`
}

func (p Polynomial) IsValid() error {

	if !(len(p.ResourceId) > 0) {
		return fmt.Errorf("empty polynomial resourceid")
	}
	if !(len(p.Name) > 0) {
		return fmt.Errorf("empty polynomial name")
	}

	if err := Validate(p.InputUnits); err != nil {
		return err
	}
	if err := Validate(p.OutputUnits); err != nil {
		return err
	}

	if err := Validate(p.ApproximationType); err != nil {
		return err
	}

	if err := Validate(p.FrequencyLowerBound); err != nil {
		return err
	}
	if err := Validate(p.FrequencyUpperBound); err != nil {
		return err
	}

	for _, c := range p.Coefficients {
		if err := Validate(c); err != nil {
			return err
		}
	}

	return nil
}
