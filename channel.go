package fdsn

import (
	"encoding/json"
	"fmt"
)

// Equivalent to SEED blockette 52 and parent element for the related the response blockettes.
type Channel struct {
	Code             string            `xml:"code,attr"`
	StartDate        *DateTime         `xml:"startDate,attr,omitempty" json:",omitempty"`
	EndDate          *DateTime         `xml:"endDate,attr,omitempty" json:",omitempty"`
	RestrictedStatus *RestrictedStatus `xml:"restrictedStatus,attr,omitempty" json:",omitempty"`
	LocationCode     string            `xml:"locationCode,attr"`

	// A code used for display or association, alternate to the SEED-compliant code.
	AlternateCode string `xml:"alternateCode,attr,omitempty" json:",omitempty"`

	// A previously used code if different from the current code.
	HistoricalCode string `xml:"historicalCode,attr,omitempty" json:",omitempty"`

	Description string    `xml:"description,omitempty" json:",omitempty"`
	Comments    []Comment `xml:"comment,omitempty" json:",omitempty"`

	// URI of any type of external report, such as data quality reports.
	ExternalReferences []ExternalReference `xml:"ExternalReference,omitempty" json:",omitempty"`

	// Latitude coordinate of this channel's sensor.
	Latitude Latitude

	//Longitude coordinate of this channel's sensor.
	Longitude Longitude

	// Elevation of the sensor.
	Elevation Distance

	// The local depth or overburden of the instrument's location.
	// For downhole instruments, the depth of the instrument under the surface ground level.
	// For underground vaults, the distance from the instrument to the local ground level above.
	Depth Distance

	// Azimuth of the sensor in degrees from north, clockwise.
	Azimuth *Azimuth `xml:",omitempty" json:",omitempty"`

	// Dip of the instrument in degrees, down from horizontal
	Dip *Dip `xml:",omitempty" json:",omitempty"`

	// The type of data this channel collects. Corresponds to channel flags in SEED blockette 52.
	// The SEED volume producer could use the first letter of an Output value as the SEED channel flag.
	Types []Type `xml:"Type,omitempty" json:",omitempty"`

	// This is a group of elements that represent sample rate.
	// If this group is included, then SampleRate, which is the sample rate in samples per second, is required.
	// SampleRateRatio, which is expressed as a ratio of number of samples in a number of seconds, is optional.
	// If both are included, SampleRate should be considered more definitive.
	SampleRate      SampleRate
	SampleRateRatio *SampleRateRatio `xml:",omitempty" json:",omitempty"`

	// The storage format of the recorded data (e.g. SEED).
	StorageFormat string
	// A tolerance value, measured in seconds per sample, used as a threshold for time
	// error detection in data from the channel.
	ClockDrift *ClockDrift `xml:",omitempty" json:",omitempty"`

	CalibrationUnits *Units     `xml:",omitempty" json:",omitempty"`
	Sensor           *Equipment `xml:",omitempty" json:",omitempty"`
	PreAmplifier     *Equipment `xml:",omitempty" json:",omitempty"`
	DataLogger       *Equipment `xml:",omitempty" json:",omitempty"`
	Equipment        *Equipment `xml:",omitempty" json:",omitempty"`
	Response         *Response  `xml:",omitempty" json:",omitempty"`
}

func (c *Channel) String() string {

	j, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(j)
}

func (c *Channel) IsValid() error {
	if c == nil {
		return nil
	}

	if !(len(c.Code) > 0) {
		return fmt.Errorf("empty code element")
	}

	if err := Validate(c.StartDate); err != nil {
		return fmt.Errorf("bad start date: %s", err)
	}
	if err := Validate(c.EndDate); err != nil {
		return fmt.Errorf("bad end date: %s", err)
	}
	if err := Validate(c.RestrictedStatus); err != nil {
		return err
	}

	if err := Validate(&c.Latitude); err != nil {
		return err
	}
	if err := Validate(&c.Longitude); err != nil {
		return err
	}
	if err := Validate(&c.Elevation); err != nil {
		return err
	}
	if err := Validate(&c.Depth); err != nil {
		return err
	}

	if err := Validate(c.Dip); err != nil {
		return err
	}
	if err := Validate(c.Azimuth); err != nil {
		return err
	}

	for _, t := range c.Types {
		if err := Validate(&t); err != nil {
			return err
		}
	}

	if err := Validate(&c.SampleRate); err != nil {
		return nil
	}
	if err := Validate(c.SampleRateRatio); err != nil {
		return nil
	}
	if !(len(c.StorageFormat) > 0) {
		return fmt.Errorf("empty code element")
	}

	if err := Validate(c.ClockDrift); err != nil {
		return nil
	}

	if err := Validate(c.CalibrationUnits); err != nil {
		return nil
	}
	if err := Validate(c.Sensor); err != nil {
		return nil
	}
	if err := Validate(c.PreAmplifier); err != nil {
		return nil
	}
	if err := Validate(c.DataLogger); err != nil {
		return nil
	}
	if err := Validate(c.Equipment); err != nil {
		return nil
	}
	if err := Validate(c.Response); err != nil {
		return nil
	}

	return nil
}

func (c *Channel) Copy(level Level) *Channel {

	var r *Response

	if level >= CHANNEL_LEVEL {
		r = c.Response.Copy(level)
	}

	return &Channel{
		Code:               c.Code,
		StartDate:          c.StartDate,
		EndDate:            c.EndDate,
		RestrictedStatus:   c.RestrictedStatus,
		LocationCode:       c.LocationCode,
		AlternateCode:      c.AlternateCode,
		HistoricalCode:     c.HistoricalCode,
		Description:        c.Description,
		Comments:           c.Comments,
		ExternalReferences: c.ExternalReferences,
		Latitude:           c.Latitude,
		Longitude:          c.Longitude,
		Elevation:          c.Elevation,
		Depth:              c.Depth,
		Azimuth:            c.Azimuth,
		Dip:                c.Dip,
		Types:              c.Types,
		SampleRate:         c.SampleRate,
		SampleRateRatio:    c.SampleRateRatio,
		StorageFormat:      c.StorageFormat,
		ClockDrift:         c.ClockDrift,
		CalibrationUnits:   c.CalibrationUnits,
		Sensor:             c.Sensor,
		PreAmplifier:       c.PreAmplifier,
		DataLogger:         c.DataLogger,
		Equipment:          c.Equipment,

		Response: r,
	}
}
