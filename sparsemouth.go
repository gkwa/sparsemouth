package sparsemouth

import (
	"fmt"
	"log/slog"
)

func Main() int {
	slog.Debug("sparsemouth", "test", true)

	// Building a traditional house
	builderA := &Traditional{}
	builderA.BuildHouse()

	fmt.Println()

	// Building a modern house
	builderB := &Modern{}
	builderB.BuildHouse()

	return 0
}
