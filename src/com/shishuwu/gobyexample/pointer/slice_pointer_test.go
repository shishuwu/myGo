package pointertest

import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
    arr := [5]int{1, 3, 5, 6, 7}
    fmt.Printf("addr:%p\n", &arr)// addr:0xc42001a1e0
	
	s1 := arr[:]
    fmt.Printf("addr:%p\n", &s1)// addr:0xc42000a060

	changeSlice(s1)
	fmt.Println(s1)
}

func changeSlice(s []int) {
    fmt.Printf("addr:%p\n", &s)// addr:0xc42000a080
	fmt.Printf("addr:%p\n", &s[0])// addr:0xc42001a1e0
	s[0] = 100
}