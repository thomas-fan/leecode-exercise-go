package longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	Words []string
	Wanted string
}

var cases = []testCase{
	{[]string{"flower", "flight", "flow"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
	{[]string{"a"}, "a"},
	{[]string{"", "b"}, ""},
	{[]string{"c", "acc", "ccc"}, ""},
	{[]string{"ab", "a"}, "a"},
}

func TestLongestCommonPrefix1(t *testing.T) {
	for _, item := range cases {
		got := LongestCommonPrefix1(item.Words)
		assert.Equal(t, item.Wanted, got)
	}
}