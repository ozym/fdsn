package fdsn

// Corresponds to SEED blockette 57.
type Decimation struct {
	InputSampleRate Frequency
	Factor          int32
	Offset          int32
	Delay           Float
	Correction      Float
}
