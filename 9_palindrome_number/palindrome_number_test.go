package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCases struct {
	Num int
	Want bool
}
func TestPalindromeNumber(t *testing.T) {
	var cases = []testCases{
		{1, true},
		{1001, true},
		{12001, false},
		{0, true},
	}

	for _, item := range cases {
		t.Run("Solution1", func(t *testing.T) {
			got := Solution1(item.Num)
			assert.Equal(t, got, item.Want)
		})
	}
}
