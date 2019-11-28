package chanrangeclosetest

import (
	"fmt"
	"testing"
)

func fib(n int, c chan int){
	x, y := 0, 1

	for i:=0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func Test(t *testing.T){
	c := make(chan int, 10)

	// go fib(10, c)
	go fib(cap(c), c)

	for x:= range c {
		fmt.Println(x)
	}
}