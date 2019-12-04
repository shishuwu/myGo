package interfacevalue

import (
	"fmt"
	"math"
	"testing"
)

type I interface {
	M()
}

// --------------------

type F float64

func (f F) M(){
	fmt.Println(f)
}

// --------------------

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

// --------------------

func Describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Test(t *testing.T) {
	var i I

	i = &T{"Hello"}
	Describe(i)
	i.M()


	i = F(math.Pi)
	Describe(i)
	i.M()

}