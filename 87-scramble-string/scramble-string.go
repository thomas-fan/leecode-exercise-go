package main

import "fmt"

type testItem struct {
	s1 string
	s2 string
	expect bool
}

var testItems = []testItem {
	{"great", "rgeat", true},
	{"abc", "bca", true},
	{"eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd", false},
}
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	if !countEq(s1, s2) {
		return false
	}

	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}

		if isScramble(s1[i:], s2[:len(s2)-i]) && isScramble(s1[:i], s2[len(s2)-i:]) {
			return true
		}
	}

	return false

}

func countEq(s1, s2 string) bool {
	tmp := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		tmp[s1[i]] ++
		tmp[s2[i]] --
	}

	for _, v := range tmp {
		if v != 0 {
			return false
		}
	}

	return true
}
func main() {
	for _, item := range testItems {
		res := isScramble(item.s1, item.s2)
		if res != item.expect {
			fmt.Printf("expected %v, but got %v, case is %v", item.expect, res, item)
		}
	}
}
