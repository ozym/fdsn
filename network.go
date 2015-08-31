package fdsn

import (
	"fmt"
	"strings"
)

// This type represents the Network layer, all station metadata is contained within this element.
// The official name of the network or other descriptive information can be included in the
// Description element. The Network can contain 0 or more Stations.
type Network struct {
	Code             string            `xml:"code,attr"`
	StartDate        *DateTime         `xml:"startDate,attr,omitempty"`
	EndDate          *DateTime         `xml:"endDate,attr,omitempty"`
	RestrictedStatus *RestrictedStatus `xml:"restrictedStatus,attr,omitempty"`

	// A code used for display or association, alternate to the SEED-compliant code.
	AlternateCode string `xml:"alternateCode,attr,omitempty"`

	// A previously used code if different from the current code.
	HistoricalCode string `xml:"historicalCode,attr,omitempty"`

	Description string    `xml:",omitempty"`
	Comments    []Comment `xml:"comment,omitempty"`

	// The total number of stations contained in this network, including inactive or terminated stations.
	TotalNumberStations uint32 `xml:",omitempty"`

	// The total number of stations in this network that were selected by the query that produced this document,
	// even if the stations do not appear in the document. (This might happen if the user only wants a document
	// that goes contains only information at the Network level.)
	SelectedNumberStations uint32 `xml:",omitempty"`

	Stations []Station `xml:"Station,omitempty"`
}

func (n Network) String() string {
	var parts []string

	parts = append(parts, fmt.Sprintf("Network: %s", n.Code))
	if n.StartDate != nil {
		parts = append(parts, fmt.Sprintf("StartDate: \"%s\"", *n.StartDate))
	}
	if n.EndDate != nil {
		parts = append(parts, fmt.Sprintf("EndDate: \"%s\"", *n.EndDate))
	}
	if n.RestrictedStatus != nil {
		parts = append(parts, fmt.Sprintf("RestrictedStatus: \"%s\"", *n.RestrictedStatus))
	}
	if n.AlternateCode != "" {
		parts = append(parts, fmt.Sprintf("AlternateCode: \"%s\"", n.AlternateCode))
	}
	if n.HistoricalCode != "" {
		parts = append(parts, fmt.Sprintf("HistoricalCode: \"%s\"", n.HistoricalCode))
	}
	if n.Description != "" {
		parts = append(parts, fmt.Sprintf("Description: \"%s\"", n.Description))
	}
	parts = append(parts, fmt.Sprintf("Comments: [%d]...", len(n.Comments)))
	if n.TotalNumberStations > 0 {
		parts = append(parts, fmt.Sprintf("TotalNumberStations: \"%d\"", n.TotalNumberStations))
	}
	if n.SelectedNumberStations > 0 {
		parts = append(parts, fmt.Sprintf("SelectedNumberStations: \"%d\"", n.SelectedNumberStations))
	}
	parts = append(parts, fmt.Sprintf("Stations: [%d]...", len(n.Stations)))

	return "<" + strings.Join(parts, "; ") + ">"
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
