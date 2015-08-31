package fdsn

type ResponseListElement struct {
	Frequency Frequency
	Amplitude Float
	Phase     Angle
}

func (r ResponseListElement) IsValid() error {
	if err := r.Frequency.IsValid(); err != nil {
		return err
	}
	if err := r.Amplitude.IsValid(); err != nil {
		return err
	}
	if err := r.Phase.IsValid(); err != nil {
		return err
	}
	return nil
}
