package sha1_test

import (
	"fmt"
	"testing"
	"crypto/sha1"
)

func Test(t *testing.T){
	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}