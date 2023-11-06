package sparsemouth

import "fmt"

// AbstractHouseBuilder defines the template method for building a house.
type AbstractHouseBuilder interface {
	BuildHouse()
	ConstructWalls()
	AddRoof()
}

// DefaultHouseBuilder provides default implementations for building a house.
type DefaultHouseBuilder struct{}

func (d *DefaultHouseBuilder) LayFoundation() {
	fmt.Println("Adding a default foundation.")
}
