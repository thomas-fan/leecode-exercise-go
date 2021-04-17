package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}
	p := 0
	q := 1

	for q < len(nums) {
		if nums[p] != nums[q] {
			nums[p + 1] = nums[q]
			p++
		}

		q ++
	}

	return p + 1
}