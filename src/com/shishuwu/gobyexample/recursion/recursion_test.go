package recurtest

import (
	"fmt"
	"testing"
)


func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n - 1)
}


func Test (t *testing.T){
	fmt.Println(fact(7))
}