package palindrome_number

func Solution1(x int) bool {
	// 负数直接 return, 或者末位是 0且不为 0 的 return
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 取最高位的除数
	highPart := 1

	// 取得最高位的除数
	for x/highPart >= 10 {
		highPart *= 10
	}

	for x > 0 {
		// 取末位
		right := x % 10
		// 取最高位
		left := x / highPart
		// 不等说明不对称
		if right != left {
			return false
		}

		// x 去掉最后一位和第一位
		x = (x % highPart) / 10
		// 最高位除数缩减 100
		highPart /= 100
	}

	return true
}
