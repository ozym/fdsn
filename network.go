package fdsn

import (
	"encoding/json"
	"fmt"
)

// Network allows grouping of station metadata.
//
// This type represents the Network layer, all station metadata is contained within this element.
// The official name of the network or other descriptive information can be included in the
// Description element. The Network can contain 0 or more Stations.
type Network struct {
	Code             string            `xml:"code,attr"`
	StartDate        *DateTime         `xml:"startDate,attr,omitempty" json:",omitempty"`
	EndDate          *DateTime         `xml:"endDate,attr,omitempty" json:",omitempty"`
	RestrictedStatus *RestrictedStatus `xml:"restrictedStatus,attr,omitempty" json:",omitempty"`

	// A code used for display or association, alternate to the SEED-compliant code.
	AlternateCode string `xml:"alternateCode,attr,omitempty" json:",omitempty"`

	// A previously used code if different from the current code.
	HistoricalCode string `xml:"historicalCode,attr,omitempty" json:",omitempty"`

	Description string    `xml:",omitempty" json:",omitempty"`
	Comments    []Comment `xml:"comment,omitempty" json:",omitempty"`

	// The total number of stations contained in this network, including inactive or terminated stations.
	TotalNumberStations uint32 `xml:",omitempty" json:",omitempty"`

	// The total number of stations in this network that were selected by the query that produced this document,
	// even if the stations do not appear in the document. (This might happen if the user only wants a document
	// that goes contains only information at the Network level.)
	SelectedNumberStations uint32 `xml:",omitempty" json:",omitempty"`

	Stations []Station `xml:"Station,omitempty" json:",omitempty"`
}

func (n Network) String() string {

	j, err := json.Marshal(&n)
	if err != nil {
		return ""
	}
	return string(j)
}

func (n Network) IsValid() error {

	if !(len(n.Code) > 0) {
		return fmt.Errorf("empty code element")
	}

	if n.StartDate != nil {
		if err := n.StartDate.IsValid(); err != nil {
			return fmt.Errorf("bad start date: %s", err)
		}
	}
	if n.EndDate != nil {
		if err := n.EndDate.IsValid(); err != nil {
			return fmt.Errorf("bad end date: %s", err)
		}
	}
	if n.RestrictedStatus != nil {
		if err := n.RestrictedStatus.IsValid(); err != nil {
			return err
		}
	}

	for _, s := range n.Stations {
		if err := Validate(s); err != nil {
			return err
		}
	}

	return nil
}
