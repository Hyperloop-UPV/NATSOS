package plate

import (
	"fmt"
	"io"

	"github.com/Hyperloop-UPV/NATSOS/pkg/adj"
	"github.com/Hyperloop-UPV/NATSOS/pkg/generator"
)

// NewMeasurementState creates a new MeasurementState for the given ADJ measurement. It initializes the generator for the measurement based on its type and range.
func NewMeasurementState(measurement adj.Measurement) *MeasurementState {

	return &MeasurementState{
		Measurement: measurement,
		Generator:   generator.SelectRandomGenerator(measurement),
	}
}

// given a mesurament state and a paylod writes its value in the payload
func (m *MeasurementState) WriteTo(w io.Writer) error {

	// Gets geneartor
	m.mu.RLock()
	gen := m.Generator
	m.mu.RUnlock()

	if gen == nil {
		return fmt.Errorf("generator not configured")
	}

	data, err := gen.Generate(m.Measurement)
	if err != nil {
		return err
	}

	// Write output into payload
	_, err = w.Write(data)
	return err

}
