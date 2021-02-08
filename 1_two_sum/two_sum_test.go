package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var nums = []int{2, 7, 11, 15}
var target = 9
var wanted = []int{0, 1}

func TestTwoSum(t *testing.T) {
	t.Run("Solution 1", func(t *testing.T) {
		got := Solution1(nums, target)
		assert.Equal(t, wanted, got)
	})
}
