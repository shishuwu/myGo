package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	arr := [3]int{10,2,3}
	t.Log(arr[0], arr[1])
}

func TestArrayTravel(t *testing.T){
	arr3 := [...]int{1,2,3,4,5}

	for _, e:=range arr3 {
		t.Log(e)
	}
}


func TestArraySection(t *testing.T){
	arr3 := [...]int{1,2,3,4,5}
	arr4 := arr3 [2:]
	for _, e:=range arr4 {
		t.Log(e)
	}
}


