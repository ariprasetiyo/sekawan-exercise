package playground

import (
	"fmt"
	"testing"
)

func TestMultipleParamsFunc(t *testing.T) {
	iniTest("a", "b", "c")
}

func iniTest(coba ...string) {
	for _, x := range coba {
		fmt.Println(x)
	}
}
