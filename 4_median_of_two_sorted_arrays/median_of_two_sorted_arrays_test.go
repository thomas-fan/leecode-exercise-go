package median_of_two_sorted_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCases struct {
	Name   string
	Nums1  []int
	Nums2  []int
	Wanted float64
}

var cases = []testCases{
	{"case1", []int{1, 3}, []int{2}, 2},
	{"case2", []int{1, 2}, []int{3, 4}, 2.5},
	{"case3", []int{0, 0}, []int{0, 0}, 0},
	{"case4", []int{}, []int{1}, 1},
	{"case5", []int{2}, []int{}, 2},
}

func TestFindMedianSortedArrays1(t *testing.T) {
	for _, item := range cases {
		t.Run(item.Name, func(t *testing.T) {
			result := findMedianSortedArrays1(item.Nums1, item.Nums2)
			assert.Equal(t, item.Wanted, result)
		})
	}
}
