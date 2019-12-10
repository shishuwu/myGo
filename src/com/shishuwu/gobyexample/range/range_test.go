package rangetest

import (
	"fmt"
	"testing"
)

func Test(t *testing.T){

	// range array
	arr := []int {2,3,4}
	num := 0

	for _, v:= range arr {
		num = num + v
	}

	fmt.Println("num ", num)
	// range map

	m := map[int]string {1:"a", 2:"b"}

	for k,v := range m {
		fmt.Println("k:", k, ", v:", v)
	}

	s := "abc"
	for i,c := range s {
		fmt.Println("index:", i, ", char", c)
	}

}