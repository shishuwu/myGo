package chansynctest

import(
	"fmt"
	"testing"
	"time"
)

func work( done chan int) {
	
	fmt.Println("work")
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	
	done <- 1
}


func Test(t *testing.T)  {
	done := make(chan int, 1)

	go work(done)

	<- done
}