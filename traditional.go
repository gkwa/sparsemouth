package sparsemouth

import "fmt"

type Traditional struct {
	*DefaultHouseBuilder
}

func (b *Traditional) BuildHouse() {
	fmt.Println("Building a traditional house:")
	b.LayFoundation()
	b.ConstructWalls()
	b.AddRoof()
}

func (b *Traditional) ConstructWalls() {
	fmt.Println("Constructing brick walls.")
}

func (b *Traditional) AddRoof() {
	fmt.Println("Adding a tiled roof.")
}
