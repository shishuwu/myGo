package readertest

import (
	"golang.org/x/tour/reader"
	"testing"
)

type MyReader struct {}

func (r MyReader) Read(b []byte) (int, error){
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func Test (t *testing.T){
	reader.Validate(MyReader{})
}
