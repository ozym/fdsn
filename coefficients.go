package fdsn

import (
	"encoding/json"
)

// Response: coefficients for FIR filter.
// Laplace transforms or IIR filters can be expressed using type as well but the
// PolesAndZerosType should be used instead. Corresponds to SEED blockette 54.
type Coefficients struct {
	// Same meaning as Equipment:resourceId.</xs:documentation>
	ResourceId string `xml:"resourceId,attr,omitempty" json:",omitempty"`

	// A name given to this filter.
	Name string `xml:"name,attr,omitempty" json:",omitempty"`

	Description string `xml:",omitempty" json:",omitempty"`

	// The units of the data as input from the perspective of data acquisition.
	// After correcting data for this response, these would be the resulting units.
	InputUnits Units

	// The units of the data as output from the perspective of data acquisition.
	// These would be the units of the data prior to correcting for this response.
	OutputUnits Units

	CfTransferFunctionType CfTransferFunctionType

	Numerators   []Float `xml:"Numerator,omitempty" json:",omitempty"`
	Denominators []Float `xml:"Denominator,omitempty" json:",omitempty"`
}

func (c Coefficients) String() string {

	j, err := json.Marshal(&c)
	if err != nil {
		return ""
	}
	return string(j)
}

func (c *Coefficients) IsValid() error {

	if c == nil {
		return nil
	}

	if err := Validate(&c.InputUnits); err != nil {
		return err
	}
	if err := Validate(&c.OutputUnits); err != nil {
		return err
	}

	if err := Validate(&c.CfTransferFunctionType); err != nil {
		return err
	}

	for _, n := range c.Numerators {
		if err := Validate(&n); err != nil {
			return err
		}
	}

	for _, d := range c.Denominators {
		if err := Validate(&d); err != nil {
			return err
		}
	}

	return nil
}
