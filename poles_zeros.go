package fdsn

import (
	"encoding/json"
	"fmt"
)

type PolesZeros struct {
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

	PzTransferFunctionType PzTransferFunctionType
	NormalizationFactor    FloatValue
	NormalizationFrequency Frequency
	Zeros                  []PoleZero `xml:"Zero,omitempty" json:",omitempty"`
	Poles                  []PoleZero `xml:"Pole,omitempty" json:",omitempty"`
}

func (pz PolesZeros) String() string {

	j, err := json.Marshal(&pz)
	if err != nil {
		return ""
	}
	return string(j)
}

func (pz *PolesZeros) IsValid() error {
	if pz == nil {
		return nil
	}

	if !(len(pz.ResourceId) > 0) {
		return fmt.Errorf("empty poles/zeros resourcedId")
	}
	if !(len(pz.Name) > 0) {
		return fmt.Errorf("empty poles/zeros resourcedId")
	}

	if err := Validate(&pz.InputUnits); err != nil {
		return err
	}
	if err := Validate(&pz.OutputUnits); err != nil {
		return err
	}

	if err := Validate(&pz.PzTransferFunctionType); err != nil {
		return err
	}
	if err := Validate(&pz.NormalizationFactor); err != nil {
		return err
	}
	if err := Validate(&pz.NormalizationFrequency); err != nil {
		return err
	}

	for _, z := range pz.Zeros {
		if err := Validate(&z); err != nil {
			return err
		}
	}

	for _, p := range pz.Poles {
		if err := Validate(&p); err != nil {
			return err
		}
	}

	return nil
}
