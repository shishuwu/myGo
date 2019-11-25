package stringertest

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}



func Test(t *testing.T) {
	jason := Person{"Jason", 18}
	maggie := Person{"Maggie", 17}

	fmt.Println(jason, maggie)
}
