package pointertest

import (
	"fmt"
	"testing"
)

func Test(t *testing.T){
	i,j := 1,21
	p := &i
	fmt.Println(*p)
	
	*p = 9
	fmt.Println(i)

	p = &j
	*p = *p / 7
	fmt.Println(j)

}