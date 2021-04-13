package median_of_two_sorted_arrays

import "sort"

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

//解法 1,暴力合并两数组,根据奇偶求出中位数, 时间复杂度O(m+n)
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	for _, num := range nums2 {
		nums1 = append(nums1, num)
	}

	sort.Ints(nums1)
	length := len(nums1)
	if length == 1 {
		return float64(nums1[0])
	}

	if length%2 == 0 {
		return (float64(nums1[length/2-1] + nums1[length/2])) / 2.0
	}

	return float64(nums1[length/2])
}
