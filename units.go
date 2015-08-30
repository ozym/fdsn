package fdsn

// A type to document units. Corresponds to SEED blockette 34.
type Units struct {
	// Name of units, e.g. "M/S", "V", "PA".
	Name string
	// Description of units, e.g. "Velocity in meters per second", "Volts", "Pascals".
	Description *string `xml:",omitempty"`
}
