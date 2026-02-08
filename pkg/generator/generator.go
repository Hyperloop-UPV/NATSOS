package generator

import "github.com/Hyperloop-UPV/NATSOS/pkg/adj"

// Generator is an interface that defines the method to generate a packet from a given board and packet definition.
type Generator interface {
	Generate(boardName string, packet adj.Packet) ([]byte, error)
}
