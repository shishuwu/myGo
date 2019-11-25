package interfaceimplicity

import (
	"fmt"
	"testing"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func Test(t *testing.T) {
	var i I = T {"hello world"}
	i.M()
}