package fdsn

import (
	"fmt"
)

type PolesZeros struct {
	// Same meaning as Equipment:resourceId.</xs:documentation>
	ResourceId string `xml:"resourceId,attr"`

	// A name given to this filter.
	Name string `xml:"name,attr"`

	Description string `xml:",omitempty"`

	// The units of the data as input from the perspective of data acquisition.
	// After correcting data for this response, these would be the resulting units.
	InputUnits Units

	// The units of the data as output from the perspective of data acquisition.
	// These would be the units of the data prior to correcting for this response.
	OutputUnits Units

	PzTransferFunctionType PzTransferFunctionType
	NormalizationFactor    FloatValue
	NormalizationFrequency Frequency
	Zeros                  []PoleZero `xml:"Zero,omitempty"`
	Poles                  []PoleZero `xml:"Pole,omitempty"`
}

func (pz PolesZeros) IsValid() error {
	if !(len(pz.ResourceId) > 0) {
		return fmt.Errorf("empty poles/zeros resourcedId")
	}
	if !(len(pz.Name) > 0) {
		return fmt.Errorf("empty poles/zeros resourcedId")
	}

	if err := pz.InputUnits.IsValid(); err != nil {
		return err
	}
	if err := pz.OutputUnits.IsValid(); err != nil {
		return err
	}

	if err := pz.PzTransferFunctionType.IsValid(); err != nil {
		return err
	}
	if err := pz.NormalizationFactor.IsValid(); err != nil {
		return err
	}
	if err := pz.NormalizationFrequency.IsValid(); err != nil {
		return err
	}

	for _, z := range pz.Zeros {
		if err := z.IsValid(); err != nil {
			return err
		}
	}

	for _, p := range pz.Poles {
		if err := p.IsValid(); err != nil {
			return err
		}
	}

	return nil
}
