package buffchantest

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ch := make(chan int, 2) // set length, buffer channel
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}