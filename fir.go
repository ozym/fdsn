package fdsn

import (
	"encoding/json"
	"fmt"
)

// Response: FIR filter. Corresponds to SEED blockette 61. FIR filters are
// also commonly documented using the CoefficientsType element.
type FIR struct {
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

	Symmetry              Symmetry
	NumeratorCoefficients []NumeratorCoefficient `xml:"NumeratorCoefficient,omitempty" json:",omitempty"`
}

func (f FIR) String() string {

	j, err := json.Marshal(&f)
	if err != nil {
		return ""
	}
	return string(j)
}

func (f *FIR) IsValid() error {
	if f == nil {
		return nil
	}

	if !(len(f.ResourceId) > 0) {
		return fmt.Errorf("empty fir resourceid")
	}
	if !(len(f.Name) > 0) {
		return fmt.Errorf("empty fir name")
	}

	if err := Validate(&f.InputUnits); err != nil {
		return err
	}
	if err := Validate(&f.OutputUnits); err != nil {
		return err
	}

	if err := Validate(&f.Symmetry); err != nil {
		return err
	}

	for _, n := range f.NumeratorCoefficients {
		if err := Validate(&n); err != nil {
			return err
		}
	}

	return nil
}
