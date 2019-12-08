package chansynctest

import(
	"fmt"
	"testing"
	"time"
)

func work() (chan int) {
	done := make(chan int, 1)
	fmt.Println("work")
	time.Sleep(time.Second)
	fmt.Println("done")
	
	done <- 1

	return done
}


func Test(t *testing.T)  {
	done := go work()
}