package fdsn

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func MapString(s string) *string {
	return &s
}

func MapDateTime(s string) *DateTime {
	t := MustParse(s)
	return &t
}

func TestMarshalling(t *testing.T) {

	f := &FDSNStationXML{
		NameSpace:     "http://www.fdsn.org/xml/station/1",
		SchemaVersion: "1.0",
		Source:        "SeisComP3",
		Sender:        MapString("WEL(GNS_Test)"),
		Created:       MustParse("2015-08-28T11:10:40"),
	}

	n := &FDSNStationXML{
		NameSpace:     "http://www.fdsn.org/xml/station/1",
		SchemaVersion: "1.0",
		Source:        "SeisComP3",
		Sender:        MapString("WEL(GNS_Test)"),
		Created:       MustParse("2015-08-28T11:10:40"),
		Networks: []Network{{
			Code:             "NZ",
			Description:      MapString("New Zealand National Seismograph Network"),
			RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
			StartDate:        MapDateTime("1980-01-01T00:00:00"),
		}},
	}

	tc, err := Parse("2008-10-13T00:00:00")
	if err != nil {
		t.Error(err)
	}

	s := &FDSNStationXML{
		NameSpace:     "http://www.fdsn.org/xml/station/1",
		SchemaVersion: "1.0",
		Source:        "SeisComP3",
		Sender:        MapString("WEL(GNS_Test)"),
		Created:       MustParse("2015-08-28T11:10:02"),
		Networks: []Network{{
			Code:             "NZ",
			Description:      MapString("New Zealand National Seismograph Network"),
			RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
			StartDate:        MapDateTime("1980-01-01T00:00:00"),
			Stations: []Station{{
				Code:             "ABAZ",
				Site:             Site{Name: "Army Bay"},
				StartDate:        MapDateTime("2008-10-13T00:00:00"),
				RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
				Latitude:         Latitude{Value: -36.600224},
				Longitude:        Longitude{Value: 174.832333},
				Elevation:        Distance{Value: 74},
				CreationDate:     tc,
			}},
		}},
	}

	c := &FDSNStationXML{
		NameSpace:     "http://www.fdsn.org/xml/station/1",
		SchemaVersion: "1.0",
		Source:        "SeisComP3",
		Sender:        MapString("WEL(GNS_Test)"),
		Created:       MustParse("2015-08-28T11:05:52"),
		Networks: []Network{{
			Code:             "NZ",
			Description:      MapString("New Zealand National Seismograph Network"),
			RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
			StartDate:        MapDateTime("1980-01-01T00:00:00"),
			Stations: []Station{{Code: "ABAZ",
				Site:             Site{Name: "Army Bay"},
				StartDate:        MapDateTime("2008-10-13T00:00:00"),
				RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
				Latitude:         Latitude{Value: -36.600224},
				Longitude:        Longitude{Value: 174.832333},
				Elevation:        Distance{Value: 74},
				CreationDate:     tc,
				Channels: []Channel{{
					Code:             "EHZ",
					LocationCode:     "10",
					RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
					Latitude:         Latitude{Value: -36.600224},
					Longitude:        Longitude{Value: 174.832333},
					Elevation:        Distance{Value: 74},
					Depth:            Distance{Value: 0},
					Azimuth:          &Azimuth{Value: 0},
					Dip:              &Dip{Value: -90},
					SampleRate:       SampleRate{Value: 100},
					SampleRateRatio: &SampleRateRatio{
						NumberSamples: 100,
						NumberSeconds: 1,
					},
					StorageFormat: "Steim2",
					ClockDrift:    &ClockDrift{Value: 0.0001},
					StartDate:     MapDateTime("2008-10-13T04:00:00"),
					EndDate:       MapDateTime("2010-03-15T02:00:00"),
					Sensor: &Equipment{
						ResourceId:  MapString("Sensor#20150130114212.658908.42"),
						Type:        MapString("L4C-3D"),
						Description: MapString("L4C-3D"),
						Model:       MapString("L4C-3D"),
					},
					DataLogger: &Equipment{
						ResourceId:  MapString("Datalogger#20150130114212.658624.40"),
						Description: MapString("ABAZ.2008.287.EHZ10"),
					},
					Response: &Response{
						InstrumentSensitivity: &Sensitivity{
							Value:       74574700,
							Frequency:   15,
							InputUnits:  Units{Name: "M/S"},
							OutputUnits: Units{Name: ""},
						},
					},
				},
					{
						Code:             "EHZ",
						LocationCode:     "10",
						RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
						Latitude:         Latitude{Value: -36.600224},
						Longitude:        Longitude{Value: 174.832333},
						Elevation:        Distance{Value: 74},
						Depth:            Distance{Value: 0},
						Azimuth:          &Azimuth{Value: 0},
						Dip:              &Dip{Value: -90},
						SampleRate:       SampleRate{Value: 100},
						SampleRateRatio: &SampleRateRatio{
							NumberSamples: 100,
							NumberSeconds: 1,
						},
						StorageFormat: "Steim2",
						ClockDrift:    &ClockDrift{Value: 0.0001},
						StartDate:     MapDateTime("2010-03-15T02:15:00"),
						Sensor: &Equipment{
							ResourceId:  MapString("Sensor#20150130114212.659492.46"),
							Type:        MapString("LE-3DliteMkII"),
							Description: MapString("LE-3DliteMkII"),
							Model:       MapString("LE-3DliteMkII"),
						},
						DataLogger: &Equipment{
							ResourceId:  MapString("Datalogger#20150130114212.659187.44"),
							Description: MapString("ABAZ.2010.074.EHZ10"),
						},
						Response: &Response{
							InstrumentSensitivity: &Sensitivity{
								Value:       167772000,
								Frequency:   15,
								InputUnits:  Units{Name: "M/S"},
								OutputUnits: Units{Name: ""},
							},
						},
					}},
			}},
		}},
	}

	r := &FDSNStationXML{
		NameSpace:     "http://www.fdsn.org/xml/station/1",
		SchemaVersion: "1.0",
		Source:        "SeisComP3",
		Sender:        MapString("WEL(GNS_Test)"),
		Created:       MustParse("2015-08-28T23:11:23"),
		Networks: []Network{{
			Code:             "NZ",
			Description:      MapString("New Zealand National Seismograph Network"),
			RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
			StartDate:        MapDateTime("1980-01-01T00:00:00"),
			Stations: []Station{{Code: "ABAZ",
				Site:             Site{Name: "Army Bay"},
				StartDate:        MapDateTime("2008-10-13T00:00:00"),
				RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
				Latitude:         Latitude{Value: -36.600224},
				Longitude:        Longitude{Value: 174.832333},
				Elevation:        Distance{Value: 74},
				CreationDate:     tc,
				Channels: []Channel{{
					Code:             "EHZ",
					LocationCode:     "10",
					RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
					Latitude:         Latitude{Value: -36.600224},
					Longitude:        Longitude{Value: 174.832333},
					Elevation:        Distance{Value: 74},
					Depth:            Distance{Value: 0},
					Azimuth:          &Azimuth{Value: 0},
					Dip:              &Dip{Value: -90},
					SampleRate:       SampleRate{Value: 100},
					SampleRateRatio: &SampleRateRatio{
						NumberSamples: 100,
						NumberSeconds: 1,
					},
					StorageFormat: "Steim2",
					ClockDrift:    &ClockDrift{Value: 0.0001},
					StartDate:     MapDateTime("2008-10-13T04:00:00"),
					EndDate:       MapDateTime("2010-03-15T02:00:00"),
					Sensor: &Equipment{
						ResourceId:  MapString("Sensor#20150130114212.658908.42"),
						Type:        MapString("L4C-3D"),
						Description: MapString("L4C-3D"),
						Model:       MapString("L4C-3D"),
					},
					DataLogger: &Equipment{
						ResourceId:  MapString("Datalogger#20150130114212.658624.40"),
						Description: MapString("ABAZ.2008.287.EHZ10"),
					},
					Response: &Response{
						InstrumentSensitivity: &Sensitivity{
							Value:       74574700,
							Frequency:   15,
							InputUnits:  Units{Name: "M/S"},
							OutputUnits: Units{Name: "COUNTS"},
						},
						Stages: []ResponseStage{
							{
								Number: 1,
								PolesZeros: &PolesZeros{
									ResourceId:             "ResponsePAZ#20150130114212.658955.43",
									Name:                   "ABAZ.2008.287.HZ10",
									InputUnits:             Units{Name: "M/S"},
									OutputUnits:            Units{Name: "V"},
									PzTransferFunctionType: PzTransferFunctionType{FUNCTION_LAPLACE_RADIANS_PER_SECOND},
									NormalizationFactor:    0.999556,
									NormalizationFrequency: Frequency{Value: 15},
									Zeros: []PoleZero{
										{Number: 2, Real: FloatNoUnit{Value: 0}, Imaginary: FloatNoUnit{Value: 0}},
										{Number: 3, Real: FloatNoUnit{Value: 0}, Imaginary: FloatNoUnit{Value: 0}},
									},
									Poles: []PoleZero{
										{Number: 0, Real: FloatNoUnit{Value: -4.2097}, Imaginary: FloatNoUnit{Value: 4.6644}},
										{Number: 1, Real: FloatNoUnit{Value: -4.2097}, Imaginary: FloatNoUnit{Value: -4.6644}},
									},
								},
								StageGain: Gain{
									Value:     177.8,
									Frequency: 15,
								},
							},
							{
								Number: 2,
								Coefficients: &Coefficients{
									InputUnits:             Units{Name: "V"},
									OutputUnits:            Units{Name: "COUNTS"},
									CfTransferFunctionType: "DIGITAL",
								},
								Decimation: &Decimation{
									InputSampleRate: Frequency{Value: 100},
									Factor:          1,
									Offset:          0,
									Delay:           Float{Value: 0.0},
									Correction:      Float{Value: 0.0},
								},
								StageGain: Gain{
									Value:     419430,
									Frequency: 0,
								},
							},
							{
								Number: 3,
								FIR: &FIR{
									ResourceId:  "ResponseFIR#20150130114212.658669.41",
									Name:        "ABAZ.10.EHZ.2008.287.stage_3",
									InputUnits:  Units{Name: "COUNTS"},
									OutputUnits: Units{Name: "COUNTS"},
									Symmetry:    "NONE",
									NumeratorCoefficients: []NumeratorCoefficient{
										{Coefficient: 0, Value: 1.31549e-11},
										{Coefficient: 1, Value: 0.000150107},
										{Coefficient: 2, Value: 0.0133968},
										{Coefficient: 3, Value: 0.164429},
										{Coefficient: 4, Value: 0.568809},
										{Coefficient: 5, Value: 0.517383},
										{Coefficient: 6, Value: -0.260836},
										{Coefficient: 7, Value: -0.122033},
										{Coefficient: 8, Value: 0.257181},
										{Coefficient: 9, Value: -0.202903},
										{Coefficient: 10, Value: 0.0707588},
										{Coefficient: 11, Value: 0.0387967},
										{Coefficient: 12, Value: -0.114313},
										{Coefficient: 13, Value: 0.13548},
										{Coefficient: 14, Value: -0.111447},
										{Coefficient: 15, Value: 0.0670548},
										{Coefficient: 16, Value: -0.0192712},
										{Coefficient: 17, Value: -0.0209313},
										{Coefficient: 18, Value: 0.0476806},
										{Coefficient: 19, Value: -0.0593383},
										{Coefficient: 20, Value: 0.0575793},
										{Coefficient: 21, Value: -0.0462333},
										{Coefficient: 22, Value: 0.0297771},
										{Coefficient: 23, Value: -0.0124829},
										{Coefficient: 24, Value: -0.00236608},
										{Coefficient: 25, Value: 0.0127882},
										{Coefficient: 26, Value: -0.0184698},
										{Coefficient: 27, Value: 0.0187973},
										{Coefficient: 28, Value: -0.0171387},
										{Coefficient: 29, Value: 0.012782},
										{Coefficient: 30, Value: -0.00767579},
										{Coefficient: 31, Value: 0.00325516},
										{Coefficient: 32, Value: -8.94756e-05},
										{Coefficient: 33, Value: -0.00177876},
										{Coefficient: 34, Value: 0.00259604},
										{Coefficient: 35, Value: -0.00266617},
										{Coefficient: 36, Value: 0.0023074},
										{Coefficient: 37, Value: -0.00177052},
										{Coefficient: 38, Value: 0.00121864},
										{Coefficient: 39, Value: -0.000746049},
										{Coefficient: 40, Value: 0.000392175},
										{Coefficient: 41, Value: -0.000158366},
										{Coefficient: 42, Value: 2.4378e-05},
										{Coefficient: 43, Value: 3.80757e-05},
										{Coefficient: 44, Value: -5.61805e-05},
										{Coefficient: 45, Value: 5.15277e-05},
										{Coefficient: 46, Value: -3.85647e-05},
										{Coefficient: 47, Value: 2.53029e-05},
										{Coefficient: 48, Value: -1.51246e-05},
										{Coefficient: 49, Value: 8.7398e-06},
										{Coefficient: 50, Value: -4.64812e-06},
										{Coefficient: 51, Value: 1.37628e-06},
										{Coefficient: 52, Value: 7.04206e-07},
										{Coefficient: 53, Value: 2.24187e-07},
										{Coefficient: 54, Value: -1.25103e-06},
										{Coefficient: 55, Value: 1.06677e-07},
										{Coefficient: 56, Value: 2.64288e-07},
										{Coefficient: 57, Value: 3.22664e-07},
										{Coefficient: 58, Value: -8.07416e-08},
										{Coefficient: 59, Value: -1.09905e-07},
										{Coefficient: 60, Value: -3.3252e-08},
										{Coefficient: 61, Value: 1.38851e-08},
										{Coefficient: 62, Value: 1.05627e-08},
										{Coefficient: 63, Value: 2.57791e-09},
										{Coefficient: 64, Value: -7.01862e-10},
									},
								},
								Decimation: &Decimation{
									InputSampleRate: Frequency{Value: 100},
									Factor:          1,
									Offset:          0,
									Delay:           Float{Value: 0.04167},
									Correction:      Float{Value: 0.04167},
								},
								StageGain: Gain{
									Value:     1,
									Frequency: 0,
								},
							},
						},
					},
				},
					{
						Code:             "EHZ",
						LocationCode:     "10",
						RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
						Latitude:         Latitude{Value: -36.600224},
						Longitude:        Longitude{Value: 174.832333},
						Elevation:        Distance{Value: 74},
						Depth:            Distance{Value: 0},
						Azimuth:          &Azimuth{Value: 0},
						Dip:              &Dip{Value: -90},
						SampleRate:       SampleRate{Value: 100},
						SampleRateRatio: &SampleRateRatio{
							NumberSamples: 100,
							NumberSeconds: 1,
						},
						StorageFormat: "Steim2",
						ClockDrift:    &ClockDrift{Value: 0.0001},
						StartDate:     MapDateTime("2010-03-15T02:15:00"),
						Sensor: &Equipment{
							ResourceId:  MapString("Sensor#20150130114212.659492.46"),
							Type:        MapString("LE-3DliteMkII"),
							Description: MapString("LE-3DliteMkII"),
							Model:       MapString("LE-3DliteMkII"),
						},
						DataLogger: &Equipment{
							ResourceId:  MapString("Datalogger#20150130114212.659187.44"),
							Description: MapString("ABAZ.2010.074.EHZ10"),
						},
						Response: &Response{
							InstrumentSensitivity: &Sensitivity{
								Value:       167772000,
								Frequency:   15,
								InputUnits:  Units{Name: "M/S"},
								OutputUnits: Units{Name: "COUNTS"},
							},
							Stages: []ResponseStage{
								{
									Number: 1,
									PolesZeros: &PolesZeros{
										ResourceId:             "ResponsePAZ#20150130114212.659544.47",
										Name:                   "ABAZ.2010.074.HZ10",
										InputUnits:             Units{Name: "M/S"},
										OutputUnits:            Units{Name: "V"},
										PzTransferFunctionType: PzTransferFunctionType{FUNCTION_LAPLACE_RADIANS_PER_SECOND},
										NormalizationFactor:    1.00008,
										NormalizationFrequency: Frequency{Value: 15},
										Zeros: []PoleZero{
											{Number: 3, Real: FloatNoUnit{Value: 0}, Imaginary: FloatNoUnit{Value: 0}},
											{Number: 4, Real: FloatNoUnit{Value: 0}, Imaginary: FloatNoUnit{Value: 0}},
											{Number: 5, Real: FloatNoUnit{Value: 0}, Imaginary: FloatNoUnit{Value: 0}},
										},
										Poles: []PoleZero{
											{Number: 0, Real: FloatNoUnit{Value: -4.444}, Imaginary: FloatNoUnit{Value: 4.444}},
											{Number: 1, Real: FloatNoUnit{Value: -4.444}, Imaginary: FloatNoUnit{Value: -4.444}},
											{Number: 2, Real: FloatNoUnit{Value: -1.083}, Imaginary: FloatNoUnit{Value: 0}},
										},
									},
									StageGain: Gain{
										Value:     400,
										Frequency: 15,
									},
								},
								{
									Number: 2,
									Coefficients: &Coefficients{
										InputUnits:             Units{Name: "V"},
										OutputUnits:            Units{Name: "COUNTS"},
										CfTransferFunctionType: "DIGITAL",
									},
									Decimation: &Decimation{
										InputSampleRate: Frequency{Value: 100},
										Factor:          1,
										Offset:          0,
										Delay:           Float{Value: 0.0},
										Correction:      Float{Value: 0.0},
									},
									StageGain: Gain{
										Value:     419430,
										Frequency: 0,
									},
								},
								{
									Number: 3,
									FIR: &FIR{
										ResourceId:  "ResponseFIR#20150130114212.659238.45",
										Name:        "ABAZ.10.EHZ.2010.074.stage_3",
										InputUnits:  Units{Name: "COUNTS"},
										OutputUnits: Units{Name: "COUNTS"},
										Symmetry:    "NONE",
										NumeratorCoefficients: []NumeratorCoefficient{
											{Coefficient: 0, Value: 1.31549e-11},
											{Coefficient: 1, Value: 0.000150107},
											{Coefficient: 2, Value: 0.0133968},
											{Coefficient: 3, Value: 0.164429},
											{Coefficient: 4, Value: 0.568809},
											{Coefficient: 5, Value: 0.517383},
											{Coefficient: 6, Value: -0.260836},
											{Coefficient: 7, Value: -0.122033},
											{Coefficient: 8, Value: 0.257181},
											{Coefficient: 9, Value: -0.202903},
											{Coefficient: 10, Value: 0.0707588},
											{Coefficient: 11, Value: 0.0387967},
											{Coefficient: 12, Value: -0.114313},
											{Coefficient: 13, Value: 0.13548},
											{Coefficient: 14, Value: -0.111447},
											{Coefficient: 15, Value: 0.0670548},
											{Coefficient: 16, Value: -0.0192712},
											{Coefficient: 17, Value: -0.0209313},
											{Coefficient: 18, Value: 0.0476806},
											{Coefficient: 19, Value: -0.0593383},
											{Coefficient: 20, Value: 0.0575793},
											{Coefficient: 21, Value: -0.0462333},
											{Coefficient: 22, Value: 0.0297771},
											{Coefficient: 23, Value: -0.0124829},
											{Coefficient: 24, Value: -0.00236608},
											{Coefficient: 25, Value: 0.0127882},
											{Coefficient: 26, Value: -0.0184698},
											{Coefficient: 27, Value: 0.0187973},
											{Coefficient: 28, Value: -0.0171387},
											{Coefficient: 29, Value: 0.012782},
											{Coefficient: 30, Value: -0.00767579},
											{Coefficient: 31, Value: 0.00325516},
											{Coefficient: 32, Value: -8.94756e-05},
											{Coefficient: 33, Value: -0.00177876},
											{Coefficient: 34, Value: 0.00259604},
											{Coefficient: 35, Value: -0.00266617},
											{Coefficient: 36, Value: 0.0023074},
											{Coefficient: 37, Value: -0.00177052},
											{Coefficient: 38, Value: 0.00121864},
											{Coefficient: 39, Value: -0.000746049},
											{Coefficient: 40, Value: 0.000392175},
											{Coefficient: 41, Value: -0.000158366},
											{Coefficient: 42, Value: 2.4378e-05},
											{Coefficient: 43, Value: 3.80757e-05},
											{Coefficient: 44, Value: -5.61805e-05},
											{Coefficient: 45, Value: 5.15277e-05},
											{Coefficient: 46, Value: -3.85647e-05},
											{Coefficient: 47, Value: 2.53029e-05},
											{Coefficient: 48, Value: -1.51246e-05},
											{Coefficient: 49, Value: 8.7398e-06},
											{Coefficient: 50, Value: -4.64812e-06},
											{Coefficient: 51, Value: 1.37628e-06},
											{Coefficient: 52, Value: 7.04206e-07},
											{Coefficient: 53, Value: 2.24187e-07},
											{Coefficient: 54, Value: -1.25103e-06},
											{Coefficient: 55, Value: 1.06677e-07},
											{Coefficient: 56, Value: 2.64288e-07},
											{Coefficient: 57, Value: 3.22664e-07},
											{Coefficient: 58, Value: -8.07416e-08},
											{Coefficient: 59, Value: -1.09905e-07},
											{Coefficient: 60, Value: -3.3252e-08},
											{Coefficient: 61, Value: 1.38851e-08},
											{Coefficient: 62, Value: 1.05627e-08},
											{Coefficient: 63, Value: 2.57791e-09},
											{Coefficient: 64, Value: -7.01862e-10},
										},
									},
									Decimation: &Decimation{
										InputSampleRate: Frequency{Value: 100},
										Factor:          1,
										Offset:          0,
										Delay:           Float{Value: 0.04167},
										Correction:      Float{Value: 0.04167},
									},
									StageGain: Gain{
										Value:     1,
										Frequency: 0,
									},
								},
							},
						},
					}},
			}},
		}},
	}

	var tests = []struct {
		n string
		x string
		s *FDSNStationXML
	}{
		{"fdsn", "testdata/fdsn.xml", f},
		{"network", "testdata/network.xml", n},
		{"station", "testdata/station.xml", s},
		{"channel", "testdata/channel.xml", c},
		{"response", "testdata/response.xml", r},
		/*
			{"ac1a", "testdata/ac1a.xml", nil},
			{"akus", "testdata/akus.xml", nil},
			{"mqz", "testdata/mqz.xml", nil},
			{"covz", "testdata/covz.xml", nil},
			{"glkz", "testdata/glkz.xml", nil},
		*/
	}

	var remaps = []struct {
		f string
		t string
	}{
		{"<Name/>", "<Name></Name>"},
		{"<ApproximationLowerBound/>", "<ApproximationLowerBound></ApproximationLowerBound>"},
		{"<ApproximationUpperBound/>", "<ApproximationUpperBound></ApproximationUpperBound>"},
		{"&apos;", "&#39;"},
	}

	/*
		for _, test := range tests {
			if test.s != nil {
				x, err := ioutil.ReadFile(test.x)
				if err != nil {
					t.Error(err)
				}

				y := string(x)
				for _, n := range remaps {
					y = strings.Replace(y, n.f, n.t, -1)
				}

				s, err := test.s.Marshal()
				if err != nil {
					t.Error(err)
				}

				if !reflect.DeepEqual(string(s), y) {
					t.Errorf("Marshal %s: \n\t%q\n\t%q\n", test.n, string(s), y)
				}
			}
		}
	*/

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

		a := string(x)
		for _, n := range remaps {
			a = strings.Replace(a, n.f, n.t, -1)
		}

		if !reflect.DeepEqual(a, string(y)) {
			t.Errorf("Unmarshal/Marshal %s: \n\t%q\n\t%q\n", test.n, a, string(y))
		}
	}
}
