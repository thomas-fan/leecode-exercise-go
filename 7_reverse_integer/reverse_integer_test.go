package reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
type testCase struct {
	Num int
	Want int
}
var cases = []testCase{
	{123, 321},
	{-246, -642},
}

func TestReverse(t *testing.T) {
	for _, item := range cases {
		t.Run("Solution1", func(t *testing.T) {
			got := Solution1(item.Num)
			assert.Equal(t, item.Want, got)
		})
	}
}
