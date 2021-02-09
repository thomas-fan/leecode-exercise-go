package romain_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	Roman string
	Want  int
}

var cases = []testCase{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInteger(t *testing.T) {
	t.Run("solution1", func(t *testing.T) {
		for _, item := range cases {
			got := RomanToInteger1(item.Roman)
			assert.Equal(t, item.Want, got)
		}
	})

}
