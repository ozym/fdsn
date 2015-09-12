package fdsn

// Sample rate expressed as number of samples in a number of seconds.
type SampleRateRatio struct {
	NumberSamples int32
	NumberSeconds int32
}

func (s SampleRateRatio) IsValid() error {
	return nil
}
