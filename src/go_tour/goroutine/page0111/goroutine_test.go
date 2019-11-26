package routinetest

import (
	"fmt"
	"time"
	"testing"
)

func Say(s string) {

	for i:=0; i<5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func Test(t *testing.T) {
	go Say("-----------")
	Say("++++++++++++++")
}