package fdsn

// Instrument sensitivities, or the complete system sensitivity, can be expressed
// using either a sensitivity value or a polynomial. The information can be used
// to convert raw data to Earth at a specified frequency or within a range of frequencies.
type Response struct {
	// Same meaning as Equipment:resourceId.
	ResourceId string `xml:"resourceId,attr,omitempty"`

	// The total sensitivity for a channel, representing the complete acquisition system expressed as a scalar.
	// Equivalent to SEED stage 0 gain with (blockette 58) with the ability to specify a frequency range.
	InstrumentSensitivity *Sensitivity `xml:",omitempty"`

	// The total sensitivity for a channel, representing the complete acquisition system expressed as a polynomial.
	// Equivalent to SEED stage 0 polynomial (blockette 62).
	InstrumentPolynomial *Polynomial `xml:",omitempty"`

	Stages []ResponseStage `xml:"Stage,omitempty"`
}