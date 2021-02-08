package two_sum

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.

// 暴力解法, 时间复杂度 O(n^2)
func Solution1(nums []int, target int) []int {
	var result []int
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i != j && num1+num2 == target {
				return append(result, i, j)
			}
		}
	}
	return result
}

// map 解法, key nums 中的int, value 为序号 index
// 时间复杂度 O(n)
func Solution2(nums []int, target int) []int {
	var result []int
	// 临时存根
	var stub = map[int]int{}
	for idx, num := range nums {
		// 另一个需要的数为 target - num
		sub := target - num
		// 判断存根中知否有另一个值, 有说明满足两数和为 target, 直接返回
		_, exist := stub[sub]
		if exist {
			return append(result, stub[sub], idx)
		}
		// 如果不存在,将当前num 和 idx 写入存根中
		stub[num] = idx
	}

	return result
}
