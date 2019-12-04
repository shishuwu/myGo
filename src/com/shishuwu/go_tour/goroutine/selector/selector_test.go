package selecttest

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T){
	c1 := make(chan string, 0)
	c2 := make(chan string, 0)

	go func(){
		time.Sleep(5 * time.Second)
		c1 <- "one"
	}()

	go func(){
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()


	for i:=0; i<2; i++ {
		select {
		case msg1 := <- c1:
			fmt.Println("received: " + msg1)

		case msg2 := <- c2:
			fmt.Println("received: " + msg2)
		}
	}
}
