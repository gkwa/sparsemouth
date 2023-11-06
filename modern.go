package sparsemouth

import "fmt"

// Modern is a concrete implementation for building a modern house.
type Modern struct{}

func (b *Modern) BuildHouse() {
	fmt.Println("Building a modern house:")
	b.LayFoundation()
	b.ConstructWalls()
	b.AddRoof()
}

func (b *Modern) LayFoundation() {
	fmt.Println("Laying a reinforced foundation with steel beams.")
}

func (b *Modern) ConstructWalls() {
	fmt.Println("Constructing glass walls with steel frames.")
}

func (b *Modern) AddRoof() {
	fmt.Println("Adding a flat, energy-efficient roof.")
}
