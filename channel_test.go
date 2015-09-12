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
		BaseNode: BaseNode{
			Code:             "EHZ",
			StartDate:        MustParse("2008-10-13T04:00:00"),
			EndDate:          MustParse("2010-03-15T02:00:00"),
			RestrictedStatus: StatusOpen,
		},
		LocationCode: "10",
		Latitude: Latitude{
			LatitudeBase: LatitudeBase{
				Value: -36.600224,
			},
		},
		Longitude: Longitude{
			LongitudeBase: LongitudeBase{
				Value: 174.832333,
			},
		},
		Elevation: Distance{Float: Float{Value: 74}},
		Depth:     Distance{Float: Float{Value: 0}},
		Azimuth:   &Azimuth{Value: 0},
		Dip:       &Dip{Value: -90},
		Types:     []Type{{TYPE_GEOPHYSICAL}},
		SampleRateGroup: SampleRateGroup{
			SampleRate: SampleRate{Float: Float{Value: 100}},
			SampleRateRatio: &SampleRateRatio{
				NumberSamples: 100,
				NumberSeconds: 1,
			},
		},
		StorageFormat: "Steim2",
		ClockDrift:    &ClockDrift{Value: 0.0001},
	}

	x, err := xml.Marshal(&c)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testChannel {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testChannel, ""}, "\n=========\n"))
	}
}
