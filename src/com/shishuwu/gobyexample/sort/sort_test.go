package sort_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	strs := []string{"b", "c", "a"}
	sort.Strings(strs)
	fmt.Println(strs)

	nums := []int {8,5,9,4,1,0}
	sort.Ints(nums)
	fmt.Println(nums)


	sorted := sort.IntsAreSorted(nums)
	fmt.Println(sorted)
}
