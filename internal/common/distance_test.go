package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsDistance(t *testing.T) {
	type testCase struct {
		name              string
		first             int
		second            int
		expectedDistance  int
		expectedDirection Direction
	}

	testCases := []testCase{
		{"both positive, first < second", 4, 6, 2, DirectionAscending},
		{"both positive, first > second", 9, 6, 3, DirectionDescending},
		{"both positive, first = second", 5, 5, 0, DirectionUnknown},
		{"both negative, first < second", -4, -1, 3, DirectionAscending},
		{"both negative, first > second", -7, -8, 1, DirectionDescending},
		{"both negative, first = second", -6, -6, 0, DirectionUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualDistance, actualDirection := AbsDistance(tc.first, tc.second)
			assert.Equal(t, tc.expectedDistance, actualDistance)
			assert.Equal(t, tc.expectedDirection, actualDirection)
		})
	}
}
