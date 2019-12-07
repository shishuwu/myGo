package exittest

import(
	"fmt"
	"os"
	"testing"
)

func Test (t *testing.T)  {
	defer fmt.Println("!")

	os.Exit(3)
}