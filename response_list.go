package fdsn

import (
	"fmt"
)

// Response: list of frequency, amplitude and phase values. Corresponds to SEED blockette 55.
type ResponseList struct {

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

	ResponseListElements []ResponseListElement `xml:"ResponseListElement,omitempty" json:",omitempty"`
}

func (r ResponseList) IsValid() error {
	if !(len(r.ResourceId) > 0) {
		return fmt.Errorf("empty response list resourceid")
	}
	if !(len(r.Name) > 0) {
		return fmt.Errorf("empty response list name")
	}

	if err := r.InputUnits.IsValid(); err != nil {
		return err
	}
	if err := r.OutputUnits.IsValid(); err != nil {
		return err
	}

	for _, e := range r.ResponseListElements {
		if err := e.IsValid(); err != nil {
			return err
		}
	}

	return nil
}
