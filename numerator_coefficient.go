package fdsn

type NumeratorCoefficient struct {
	Coefficient int32      `xml:"i,attr"`
	Value       FloatValue `xml:",chardata"`
}