package readertest

import (
	"fmt"
	"testing"
	"io"
	"strings"
)

func Test (t *testing.T){
	r := strings.NewReader("Hello world!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v, err = %v, b=%v \n", n, err, b)
		fmt.Printf("b[n] = %q \n", b[:n])
		if err == io.EOF {
			break
		}
	}


}