package emptyinterface

import (
	"fmt"
	"testing"
)

// empty interface
type I interface {
}

func Test(t *testing.T) {
	var i I
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}