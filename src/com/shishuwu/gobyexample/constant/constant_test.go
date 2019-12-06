package constanttest

import (
	"fmt"
	"math"
	"testing"
)

const s string = "constant"

func Test(t *testing.T){
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}

