package errortest

import (
	"fmt"
	"testing"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	result := doSqrt(x)
	return result, nil
}

/** private method - do sqrt algo */
func doSqrt(x float64) float64 {
	z := float64(2.)
	s := float64(0)
	for {
		z = z - (z*z - x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}
	return s
}

func Test(t *testing.T) {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

