package fdsn

// Complex numbers used as poles or zeros in channel response.
type PoleZero struct {
	Number    int32 `xml:"number,attr"`
	Real      FloatNoUnit
	Imaginary FloatNoUnit
}
