package main

import "fmt"

type Per struct {
	name string
	bh   int
}

func main() {
	// p := &Per{}

	// for i := 0; i < 100; i++ {
	// 	p.bh = i
	// 	go func(p *Per) {
	// 		fmt.Println(p)
	// 	}(p)
	// }


	for i := 0;i < 100; i++{
        go func(i int) {
            fmt.Println(i)
        }(i)
	}

}
