package fdsn

import (
	"encoding/json"
	"fmt"
)

// This type represents a Station epoch.
// It is common to only have a single station epoch with the station's creation
// and termination dates as the epoch start and end dates.
type Station struct {
	Code             string            `xml:"code,attr"`
	StartDate        *DateTime         `xml:"startDate,attr,omitempty" json:",omitempty"`
	EndDate          *DateTime         `xml:"endDate,attr,omitempty" json:",omitempty"`
	RestrictedStatus *RestrictedStatus `xml:"restrictedStatus,attr,omitempty" json:",omitempty"`

	// A code used for display or association, alternate to the SEED-compliant code.
	AlternateCode string `xml:"alternateCode,attr,omitempty" json:",omitempty"`

	// A previously used code if different from the current code.
	HistoricalCode string `xml:"historicalCode,attr,omitempty" json:",omitempty"`

	Description string    `xml:"description,omitempty" json:",omitempty"`
	Comments    []Comment `xml:"comment,omitempty" json:",omitempty"`

	Latitude  Latitude
	Longitude Longitude
	Elevation Distance

	// These fields describe the location of the station using geopolitical
	// entities (country, city, etc.).
	Site Site

	// Type of vault, e.g. WWSSN, tunnel, transportable array, etc.
	Vault string `xml:",omitempty" json:",omitempty"`

	// Type of rock and/or geologic formation.
	Geology string `xml:",omitempty" json:",omitempty"`

	// Equipment used by all channels at a station.
	Equipments []Equipment `xml:"Equipment,omitempty" json:",omitempty"`

	// An operating agency and associated contact persons.
	// If there multiple operators, each one should be encapsulated within an Operator tag.
	// Since the Contact element is a generic type that represents any contact person,
	// it also has its own optional Agency element.
	Operators []Operator `xml:"Operator,omitempty" json:",omitempty"`

	// Date and time (UTC) when the station was first installed.
	CreationDate DateTime

	// Date and time (UTC) when the station was terminated or will be terminated.
	// A blank value should be assumed to mean that the station is still active.
	TerminationDate *DateTime `xml:",omitempty" json:",omitempty"`

	// Total number of channels recorded at this station.
	TotalNumberChannels uint32 `xml:",omitempty" json:",omitempty"`

	// Number of channels recorded at this station and selected by the query
	// that produced this document.
	SelectedNumberChannels uint32 `xml:",omitempty" json:",omitempty"`

	// URI of any type of external report, such as IRIS data reports or dataless SEED volumes.
	ExternalReferences []ExternalReference `xml:"ExternalReference,omitempty" json:",omitempty"`

	Channels []Channel `xml:"Channel,omitempty" json:",omitempty"`
}

func (s *Station) String() string {

	j, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(j)
}

func (s *Station) IsValid() error {
	if s == nil {
		return nil
	}

	if !(len(s.Code) > 0) {
		return fmt.Errorf("empty code element")
	}

	if err := Validate(s.StartDate); err != nil {
		return fmt.Errorf("bad start date: %s", err)
	}
	if err := Validate(s.EndDate); err != nil {
		return fmt.Errorf("bad end date: %s", err)
	}
	if err := Validate(s.RestrictedStatus); err != nil {
		return err
	}

	if err := Validate(&s.Latitude); err != nil {
		return err
	}
	if err := Validate(&s.Longitude); err != nil {
		return err
	}
	if err := Validate(&s.Elevation); err != nil {
		return err
	}
	if err := Validate(&s.Site); err != nil {
		return err
	}

	for _, e := range s.Equipments {
		if err := Validate(&e); err != nil {
			return err
		}
	}

	for _, o := range s.Operators {
		if err := Validate(&o); err != nil {
			return err
		}
	}

	if err := Validate(&s.CreationDate); err != nil {
		return err
	}

	if err := Validate(s.TerminationDate); err != nil {
		return err
	}

	for _, x := range s.ExternalReferences {
		if err := Validate(&x); err != nil {
			return err
		}
	}

	for _, c := range s.Channels {
		if err := Validate(&c); err != nil {
			return err
		}
	}

	return nil
}

func (s *Station) Copy(level Level) *Station {

	var c []Channel

	if level > STATION_LEVEL {
		for _, y := range s.Channels {
			c = append(c, *y.Copy(level))
		}
	}

	return &Station{

		Code:                   s.Code,
		StartDate:              s.StartDate,
		EndDate:                s.EndDate,
		RestrictedStatus:       s.RestrictedStatus,
		AlternateCode:          s.AlternateCode,
		HistoricalCode:         s.HistoricalCode,
		Description:            s.Description,
		Comments:               s.Comments,
		Latitude:               s.Latitude,
		Longitude:              s.Longitude,
		Elevation:              s.Elevation,
		Site:                   s.Site,
		Vault:                  s.Vault,
		Geology:                s.Geology,
		Equipments:             s.Equipments,
		Operators:              s.Operators,
		CreationDate:           s.CreationDate,
		TerminationDate:        s.TerminationDate,
		TotalNumberChannels:    s.TotalNumberChannels,
		SelectedNumberChannels: s.SelectedNumberChannels,
		ExternalReferences:     s.ExternalReferences,

		Channels: c,
	}
}
