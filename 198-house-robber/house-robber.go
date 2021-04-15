package main

import "fmt"

var testCases = [][]int{
	{1},
	{1, 3, 1, 3, 100},
	{2, 3, 2},
	{1, 2, 3, 1},
	{0},
}


func rob(nums []int) int {
	L := len(nums)
	if L == 0 {
		return 0
	}

	if L == 1 {
		return nums[0]
	}

	if L == 2 {
		return maxVal(nums[0], nums[1])
	}

	var prev ,cur, next int

	prev = nums[0]
	cur = maxVal(prev, nums[1])
	for i := 2; i < len(nums); i++ {
		next = maxVal(prev + nums[i], cur)
		prev = cur
		cur = next
	}

	return next
}

func maxVal(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func main() {
	for _, cas := range testCases {
		fmt.Println(rob(cas))
	}
}
