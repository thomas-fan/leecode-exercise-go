package two_sum

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.

// 暴力解法
func Solution1(nums []int, target int) []int {
	var result [] int
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1 + num2 == target {
				return append(result, i, j)
			}
		}
	}
	return result
}
