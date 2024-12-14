package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGradualProgression(t *testing.T) {
	type testCase struct {
		name     string
		levels   []int
		expected bool
	}

	testCases := []testCase{
		{"levels = {7, 6, 4, 2, 1}", []int{7, 6, 4, 2, 1}, true},
		{"levels = {1, 2, 7, 8, 9}", []int{1, 2, 7, 8, 9}, false},
		{"levels = {9, 7, 6, 2, 1}", []int{9, 7, 6, 2, 1}, false},
		{"levels = {1, 3, 2, 4, 5}", []int{1, 3, 2, 4, 5}, false},
		{"levels = {8, 6, 4, 4, 1}", []int{8, 6, 4, 4, 1}, false},
		{"levels = {1, 3, 6, 7, 9}", []int{1, 3, 6, 7, 9}, true},
		{"levels = {3, 3, 6, 7, 9}", []int{3, 3, 6, 7, 9}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := checkGradualProgression(tc.levels)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
