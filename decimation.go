package fdsn

// Corresponds to SEED blockette 57.
type Decimation struct {
	InputSampleRate Frequency
	Factor          int32
	Offset          int32
	Delay           Float
	Correction      Float
}

func (d Decimation) IsValid() error {
	if err := d.InputSampleRate.IsValid(); err != nil {
		return err
	}
	if err := d.Delay.IsValid(); err != nil {
		return err
	}
	if err := d.Correction.IsValid(); err != nil {
		return err
	}
	return nil
}
