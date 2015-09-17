package stationxml

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMarshalling_Channel(t *testing.T) {

	var tests = []struct {
		x string
		r FDSNStationXML
	}{
		{"testdata/channel.xml", FDSNStationXML{
			NameSpace:     "http://www.fdsn.org/xml/station/1",
			SchemaVersion: "1.0",
			Source:        "SeisComP3",
			Sender:        "WEL(GNS_Test)",
			Created:       MustParseDateTime("2015-08-28T11:05:52"),
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
					Channels: []Channel{{
						BaseNode: BaseNode{
							Code:             "EHZ",
							StartDate:        MustParseDateTimePtr("2008-10-13T04:00:00"),
							EndDate:          MustParseDateTimePtr("2010-03-15T02:00:00"),
							RestrictedStatus: StatusOpen,
						},
						LocationCode: "10",
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
						Elevation: Distance{Float: Float{Value: 74}},
						Depth:     Distance{Float: Float{Value: 0}},
						Azimuth:   &Azimuth{Float: Float{Value: 0}},
						Dip:       &Dip{Float: Float{Value: -90}},
						SampleRateGroup: SampleRateGroup{
							SampleRate: SampleRate{Float: Float{Value: 100}},
							SampleRateRatio: &SampleRateRatio{
								NumberSamples: 100,
								NumberSeconds: 1,
							},
						},
						StorageFormat: "Steim2",
						ClockDrift:    &ClockDrift{Float: Float{Value: 0.0001}},
						Sensor: &Equipment{
							ResourceId:  "Sensor#20150130114212.658908.42",
							Type:        "L4C-3D",
							Description: "L4C-3D",
							Model:       "L4C-3D",
						},
						DataLogger: &Equipment{
							ResourceId:  "Datalogger#20150130114212.658624.40",
							Description: "ABAZ.2008.287.EHZ10",
						},
						Response: &Response{
							InstrumentSensitivity: &Sensitivity{
								Gain: Gain{
									Value:     74574700,
									Frequency: 15,
								},
								InputUnits:  Units{Name: "M/S"},
								OutputUnits: Units{Name: ""},
							},
						},
					},
						{
							BaseNode: BaseNode{
								Code:             "EHZ",
								StartDate:        MustParseDateTimePtr("2010-03-15T02:15:00"),
								RestrictedStatus: StatusOpen,
							},
							LocationCode: "10",
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
							Elevation: Distance{Float: Float{Value: 74}},
							Depth:     Distance{Float: Float{Value: 0}},
							Azimuth:   &Azimuth{Float: Float{Value: 0}},
							Dip:       &Dip{Float: Float{Value: -90}},
							SampleRateGroup: SampleRateGroup{
								SampleRate: SampleRate{Float: Float{Value: 100}},
								SampleRateRatio: &SampleRateRatio{
									NumberSamples: 100,
									NumberSeconds: 1,
								},
							},
							StorageFormat: "Steim2",
							ClockDrift:    &ClockDrift{Float: Float{Value: 0.0001}},
							Sensor: &Equipment{
								ResourceId:  "Sensor#20150130114212.659492.46",
								Type:        "LE-3DliteMkII",
								Description: "LE-3DliteMkII",
								Model:       "LE-3DliteMkII",
							},
							DataLogger: &Equipment{
								ResourceId:  "Datalogger#20150130114212.659187.44",
								Description: "ABAZ.2010.074.EHZ10",
							},
							Response: &Response{
								InstrumentSensitivity: &Sensitivity{
									Gain: Gain{
										Value:     1.67772e+08,
										Frequency: 15,
									},
									InputUnits:  Units{Name: "M/S"},
									OutputUnits: Units{Name: ""},
								},
							},
						}},
				}},
			}},
		},
		},
	}

	for _, test := range tests {
		if err := Validate(test.r); err != nil {
			t.Errorf("Channel struct is not valid: %s", err)
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
			t.Errorf("Channel Marshal:\n\t%q\n\t%q\n", string(x), string(y))
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
			t.Errorf("Channel Unmarshal/Marshal:\n\t%q\n\t%q\n", string(x), string(y))
		}
	}
}
