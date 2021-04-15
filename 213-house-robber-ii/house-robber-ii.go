package main

import "fmt"

//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/house-robber-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var testCases = [][]int{
	{1},
	{1, 3, 1, 3, 100},
	{2, 3, 2},
	{1, 2, 3, 1},
	{0},
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	nums1 := nums[:len(nums)-1]
	nums2 := nums[1:]
	return maxVal(maxMoney(nums1), maxMoney(nums2))
}

func maxMoney(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	var dp []int
	dp = append(dp, nums[0])
	dp = append(dp, maxVal(dp[0], nums[1]))
	for i := 2; i < len(nums); i++ {
		dp = append(dp, maxVal(dp[i -2] + nums[i], dp[i -1]))
	}

	return dp[len(dp) - 1]
}

func maxVal(vals ... int) int {
	var max int
	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max
}

func main() {
	for _, cas := range testCases {
		fmt.Println(rob(cas))
	}
}
