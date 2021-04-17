package main

//给你一个整数数组 nums 和两个整数k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
//
//如果存在则返回 true，不存在返回 false。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/contains-duplicate-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	bucket := make(map[int]int)
	for i, num := range nums {
		bucketId := getBucketId(num, t)

		if i-(k+1) >= 0 {
			delete(bucket, getBucketId(nums[i-(k+1)], t))
		}

		_, ok := bucket[bucketId]
		if ok {
			return true
		}

		val, ok := bucket[bucketId-1]
		if ok && abs(val-num) <= t {
			return true
		}

		val, ok = bucket[bucketId+1]
		if ok && abs(val-num) <= t {
			return true
		}
		bucket[bucketId] = num
	}

	return false
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return n * -1
	}
}

func getBucketId(i, cap int) int {
	if i >= 0 {
		return i / (cap + 1)
	} else {
		return (i+1)/(cap+1) - 1
	}
}
