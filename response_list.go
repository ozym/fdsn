package fdsn

import (
	"encoding/json"
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

func (r ResponseList) String() string {

	j, err := json.Marshal(&r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (r *ResponseList) IsValid() error {
	if r == nil {
		return nil
	}

	if !(len(r.ResourceId) > 0) {
		return fmt.Errorf("empty response list resourceid")
	}
	if !(len(r.Name) > 0) {
		return fmt.Errorf("empty response list name")
	}

	if err := Validate(&r.InputUnits); err != nil {
		return err
	}
	if err := Validate(&r.OutputUnits); err != nil {
		return err
	}

	for _, e := range r.ResponseListElements {
		if err := Validate(&e); err != nil {
			return err
		}
	}

	return nil
}
