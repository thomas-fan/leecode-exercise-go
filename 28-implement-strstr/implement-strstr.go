package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	delta := len(haystack) - len(needle)

	for i := 0; i <= delta; i++ {
		j := 0
		startIndex := i
		for j < len(needle) {
			if haystack[startIndex] == needle[j] {
				j++
				startIndex++
			} else {
				break
			}
		}

		if len(needle) == j {
			return i
		}

	}

	return -1
}
