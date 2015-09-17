package stationxml

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMarshalling_Station(t *testing.T) {

	var tests = []struct {
		x string
		r FDSNStationXML
	}{
		{"testdata/station.xml", FDSNStationXML{

			NameSpace:     "http://www.fdsn.org/xml/station/1",
			SchemaVersion: "1.0",
			Source:        "SeisComP3",
			Sender:        "WEL(GNS_Test)",
			Created:       MustParseDateTime("2015-08-28T11:10:02"),
			Networks: []Network{{
				BaseNode: BaseNode{
					Code:             "NZ",
					Description:      "New Zealand National Seismograph Network",
					RestrictedStatus: StatusOpen,
					StartDate:        MustParseDateTimePtr("1980-01-01T00:00:00"),
				},
				Stations: []Station{{
					BaseNode: BaseNode{
						Code:             "ABAZ",
						StartDate:        MustParseDateTimePtr("2008-10-13T00:00:00"),
						RestrictedStatus: StatusOpen,
					},
					Site: Site{Name: "Army Bay"},
					Latitude: Latitude{
						LatitudeBase: LatitudeBase{
							Float: Float{
								Value: -36.600224,
							},
						},
					},
					Longitude: Longitude{
						LongitudeBase: LongitudeBase{
							Float: Float{
								Value: 174.832333,
							},
						},
					},
					Elevation:    Distance{Float: Float{Value: 74}},
					CreationDate: MustParseDateTime("2008-10-13T00:00:00"),
				}},
			}},
		},
		},
	}

	for _, test := range tests {
		if err := Validate(test.r); err != nil {
			t.Errorf("Station struct is not valid: %s", err)
		}
	}

	for _, test := range tests {
		x, err := ioutil.ReadFile(test.x)
		if err != nil {
			t.Error(err)
		}

		y, err := test.r.Marshal()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(string(x), string(y)) {
			t.Errorf("Station Marshal:\n\t%q\n\t%q\n", string(x), string(y))
		}
	}

	for _, test := range tests {

		x, err := ioutil.ReadFile(test.x)
		if err != nil {
			t.Error(err)
		}

		var s FDSNStationXML
		err = xml.Unmarshal(x, &s)
		if err != nil {
			t.Error(err)
		}

		y, err := s.Marshal()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(string(x), string(y)) {
			t.Errorf("Station Unmarshal/Marshal: \n\t%q\n\t%q\n", string(x), string(y))
		}
	}
}
