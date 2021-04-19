package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}


	p, q := 0, 0
	for q < len(nums) {
		if nums[q] != val {
			nums[p] = nums[q]
			p++
		}

		q++
	}

	return p
}
