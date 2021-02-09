package longest_common_prefix

import (
	"math"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
type stub struct {
	Count  int
	Prefix string
}

func LongestCommonPrefix1(strs []string) string {
	var record stub
	// 只有一组字符串则直接返回
	if len(strs) == 1 {
		return strs[0]
	}

	for index, str := range strs {
		if index == 0 {
			record.Prefix = str
			record.Count = len(str)
			continue
		}
		// 取当前字符串和缓存字符串最小长度
		length := math.Min(float64(record.Count), float64(len(str)))
		var count int
		var tmp string
		for i := 0; i < int(length); i++ {
			// 如果两个字符串有一个不等则退出循环
			if record.Prefix[i] != str[i] {
				break
			}

			tmp += string(str[i])
			count++
		}

		// 只要相同的数量 count 比之前存的小,则替换
		if count <= record.Count {
			record.Count = count
			record.Prefix = tmp
		}
	}

	return record.Prefix
}
