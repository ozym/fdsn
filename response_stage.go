package fdsn

// This complex type represents channel response and covers SEED blockettes 53 to 56.
type ResponseStage struct {
	Number int32 `xml:"number,attr"`

	// A choice of response types. There should be one response per stage.
	PolesZeros   *PolesZeros   `xml:",omitempty"`
	Coefficients *Coefficients `xml:",omitempty"`
	ResponseList *ResponseList `xml:,omitempty"`
	FIR          *FIR          `xml:",omitempty"`
	Polynomial   *Polynomial   `xml:",omitempty"`

	Decimation *Decimation `xml:",omitempty"`
	StageGain  Gain
}
