package playground

import (
	"fmt"
	"testing"
)

func TestEnum(t *testing.T) {
	// Declaring a variable myDirection with type Direction
	var myDirection Direction
	myDirection = West

	if myDirection == West {
		fmt.Println("myDirection is West:", myDirection)
	}
}

type Direction int

const (
	North Direction = iota
	South Direction = iota
	East  Direction = iota
	West  Direction = iota
)
