package mapext

import "testing"

func TestMapWithFunValue(t *testing.T){
	m := map[int] func(op int) int {}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](2), m[2](2), m[3](2))
	t.Log(m[1](4), m[1](6), m[1](8))

}


func TestMapAsSet(t *testing.T){
	mySet := map[int]bool{}
	mySet[1] = true

	n := 1
	exist := mySet[n]

	if exist {
		t.Logf("value %d exists", n)
	} else {
		t.Logf("value %d not exists", n)
	}

	// ----------------
	delete(mySet, 1)
	exist = mySet[n]
	
	if exist {
		t.Logf("value %d exists", n)
	} else {
		t.Logf("value %d not exists", n)
	}
}