package fdsn

import (
	"fmt"
)

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

func (r ResponseStage) IsValid() error {
	if r.Number < 0 {
		return fmt.Errorf("invalid response stage number: %d", r.Number)
	}

	if r.PolesZeros != nil {
		if err := r.PolesZeros.IsValid(); err != nil {
			return err
		}
	}
	if r.Coefficients != nil {
		if err := r.Coefficients.IsValid(); err != nil {
			return err
		}
	}
	if r.ResponseList != nil {
		if err := r.ResponseList.IsValid(); err != nil {
			return err
		}
	}
	if r.FIR != nil {
		if err := r.FIR.IsValid(); err != nil {
			return err
		}
	}
	if r.Polynomial != nil {
		if err := r.Polynomial.IsValid(); err != nil {
			return err
		}
	}
	if r.Decimation != nil {
		if err := r.Decimation.IsValid(); err != nil {
			return err
		}
	}
	if err := r.StageGain.IsValid(); err != nil {
		return err
	}

	return nil
}
