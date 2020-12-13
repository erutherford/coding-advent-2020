package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findTwoEntriesWithSum(t *testing.T) {
	tests := []struct {
		name        string
		entries     []int64
		expectedOne int
		expectedTwo int
	}{
		{
			name:        "should return the correct indexes",
			entries:     []int64{1721, 979, 366, 299, 675, 1456},
			expectedOne: 0,
			expectedTwo: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			resultOne, resultTwo := findTwoEntriesWithSum(tt.entries)
			assert.Equal(tt.expectedOne, resultOne)
			assert.Equal(tt.expectedTwo, resultTwo)
		})
	}
}

func Test_findThreeEntriesWithSum(t *testing.T) {
	tests := []struct {
		name          string
		entries       []int64
		expectedOne   int
		expectedTwo   int
		expectedThree int
	}{
		{
			name:          "should return the correct indexes",
			entries:       []int64{1721, 979, 366, 299, 675, 1456},
			expectedOne:   1,
			expectedTwo:   2,
			expectedThree: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			resultOne, resultTwo, resultThree := findThreeEntriesWithSum(tt.entries)
			assert.Equal(tt.expectedOne, resultOne)
			assert.Equal(tt.expectedTwo, resultTwo)
			assert.Equal(tt.expectedThree, resultThree)
		})
	}
}
