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

func (c Channel) String() string {

	j, err := json.Marshal(&c)
	if err != nil {
		return ""
	}
	return string(j)
}

func (c Channel) IsValid() error {
	if !(len(c.Code) > 0) {
		return fmt.Errorf("empty code element")
	}

	if c.StartDate != nil {
		if err := c.StartDate.IsValid(); err != nil {
			return fmt.Errorf("bad start date: %s", err)
		}
	}
	if c.EndDate != nil {
		if err := c.EndDate.IsValid(); err != nil {
			return fmt.Errorf("bad end date: %s", err)
		}
	}
	if c.RestrictedStatus != nil {
		if err := c.RestrictedStatus.IsValid(); err != nil {
			return err
		}
	}

	if err := c.Latitude.IsValid(); err != nil {
		return err
	}
	if err := c.Longitude.IsValid(); err != nil {
		return err
	}
	if err := c.Elevation.IsValid(); err != nil {
		return err
	}
	if err := c.Depth.IsValid(); err != nil {
		return err
	}

	if c.Dip != nil {
		if err := c.Dip.IsValid(); err != nil {
			return err
		}
	}
	if c.Azimuth != nil {
		if err := c.Azimuth.IsValid(); err != nil {
			return err
		}
	}

	for _, t := range c.Types {
		if err := t.IsValid(); err != nil {
			return err
		}
	}

	if err := c.SampleRate.IsValid(); err != nil {
		return nil
	}
	if c.SampleRateRatio != nil {
		if err := c.SampleRateRatio.IsValid(); err != nil {
			return nil
		}
	}
	if !(len(c.StorageFormat) > 0) {
		return fmt.Errorf("empty code element")
	}

	if c.ClockDrift != nil {
		if err := c.ClockDrift.IsValid(); err != nil {
			return nil
		}
	}

	if c.CalibrationUnits != nil {
		if err := c.CalibrationUnits.IsValid(); err != nil {
			return nil
		}
	}
	if c.Sensor != nil {
		if err := c.Sensor.IsValid(); err != nil {
			return nil
		}
	}
	if c.PreAmplifier != nil {
		if err := c.PreAmplifier.IsValid(); err != nil {
			return nil
		}
	}
	if c.DataLogger != nil {
		if err := c.DataLogger.IsValid(); err != nil {
			return nil
		}
	}
	if c.Equipment != nil {
		if err := c.Equipment.IsValid(); err != nil {
			return nil
		}
	}
	if c.Response != nil {
		if err := c.Response.IsValid(); err != nil {
			return nil
		}
	}

	return nil
}
