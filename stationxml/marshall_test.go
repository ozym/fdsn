package stationxml

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestFDSNStationXML_Marshalling(t *testing.T) {

	var tests = []struct {
		n string
		x string
	}{
		{"ac2a", "testdata/ac1a.xml"},
		{"akus", "testdata/akus.xml"},
		{"mqz", "testdata/mqz.xml"},
		{"covz", "testdata/covz.xml"},
		{"glkz", "testdata/glkz.xml"},
	}

	for _, test := range tests {
		x, err := ioutil.ReadFile(test.x)
		if err != nil {
			t.Error(err)
		}

		var a FDSNStationXML
		err = xml.Unmarshal(x, &a)
		if err != nil {
			t.Error(err)
		}

		y, err := a.Marshal()
		if err != nil {
			t.Error(err)
		}

		var b FDSNStationXML
		err = xml.Unmarshal(y, &b)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(a, b) {
			t.Errorf("Unmarshal/Marshal:\n\t%q\n\t%q\n", string(x), string(y))
		}
	}
}
