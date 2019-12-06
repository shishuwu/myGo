package closuretest

import (
	"testing"
	"fmt"
)

func intSeq() func() int {

	i := 0

	// 函数中的函数
	// 可以访问外边变量
	return func() int {
		i = i + 1
		return i
	}
}

func Test (t *testing.T) {

	addOne := intSeq()

	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())
}