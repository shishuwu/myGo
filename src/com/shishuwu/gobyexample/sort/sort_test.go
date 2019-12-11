package sort_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_sort(t *testing.T) {
	strs := []string{"b", "c", "a"}
	sort.Strings(strs)
	fmt.Println(strs)

	nums := []int{8, 5, 9, 4, 1, 0}
	sort.Ints(nums)
	fmt.Println(nums)

	sorted := sort.IntsAreSorted(nums)
	fmt.Println(sorted)
}

// -------------------------------------------
// customized sorter
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Test_customized_sort(t *testing.T) {
	fruits := []string {"apple", "orange", "abcde", "lemon", "pear"}
	sort.Sort(ByLength(fruits));

	fmt.Println(fruits)
}
