package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestChannel_Marshal(t *testing.T) {

	testChannel := strings.Join([]string{
		"<Channel code=\"EHZ\" startDate=\"2008-10-13T04:00:00\" endDate=\"2010-03-15T02:00:00\" restrictedStatus=\"open\" locationCode=\"10\">",
		"<Latitude>-36.600224</Latitude>",
		"<Longitude>174.832333</Longitude>",
		"<Elevation>74</Elevation>",
		"<Depth>0</Depth>",
		"<Azimuth>0</Azimuth>",
		"<Dip>-90</Dip>",
		"<Type>GEOPHYSICAL</Type>",
		"<SampleRate>100</SampleRate>",
		"<SampleRateRatio>",
		"<NumberSamples>100</NumberSamples>",
		"<NumberSeconds>1</NumberSeconds>",
		"</SampleRateRatio>",
		"<StorageFormat>Steim2</StorageFormat>",
		"<ClockDrift>0.0001</ClockDrift>",
		"</Channel>",
	}, "")

	c := &Channel{
		Code:             "EHZ",
		LocationCode:     "10",
		RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
		Latitude:         Latitude{Value: -36.600224},
		Longitude:        Longitude{Value: 174.832333},
		Elevation:        Distance{Value: 74},
		Depth:            Distance{Value: 0},
		Azimuth:          &Azimuth{Value: 0},
		Dip:              &Dip{Value: -90},
		Types:            []Type{{TYPE_GEOPHYSICAL}},
		SampleRate:       SampleRate{Value: 100},
		SampleRateRatio: &SampleRateRatio{
			NumberSamples: 100,
			NumberSeconds: 1,
		},
		StorageFormat: "Steim2",
		ClockDrift:    &ClockDrift{Value: 0.0001},
		StartDate:     MapDateTime("2008-10-13T04:00:00"),
		EndDate:       MapDateTime("2010-03-15T02:00:00"),
	}

	x, err := xml.Marshal(&c)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testChannel {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testChannel, ""}, "\n=========\n"))
	}
}
