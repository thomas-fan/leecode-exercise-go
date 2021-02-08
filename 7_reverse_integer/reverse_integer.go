package reverse_integer

import "math"

//给你一个 32 位的有符号整数 x ，返回 x 中每位上的数字反转后的结果。
//
//如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31− 1] ，就返回 0。
//
//假设环境不允许存储 64 位整数（有符号或无符号）。

// 整体思路,对 10 取余取得每个末位,然后*10 倒序拼接
func Solution1(x int) int {
	result := 0
	for x >= 10 || x <= -10 {
		item := x % 10
		x = x / 10
		result = result * 10 + item
	}

	result = result * 10 + x
	if result > int(math.Pow(2.0, 31) - 1) || result < int(math.Pow(2.0, 31) * -1) {
		return 0
	}

	return result
}