package fileutils_test

import (
	"testing"

	"github.com/erutherford/coding-advent-2020/pkg/fileutils"

	"github.com/stretchr/testify/assert"
)

func TestReadLinesFromFile(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		expected    []string
		expectedErr string
	}{
		{
			name:        "should read a file with strings correctly",
			path:        "testdata/simple_strings.txt",
			expected:    []string{"test", "with", "many", "strings", "apples", "oranges", "bananas", "kiwi", "cucumbers"},
			expectedErr: "",
		},
		{
			name:        "should read a file with strings correctly",
			path:        "testdata/strings_with_spaces.txt",
			expected:    []string{"test apples", "apples test", "test with multiple spaces", "sp3c!@l ch@r@ct3rs!%%@#d}P{13"},
			expectedErr: "",
		},
		{
			name:        "should should return an error if the file doesn't exist",
			path:        "testdata/does_not_exist.txt",
			expected:    []string{},
			expectedErr: "error opening file: open testdata/does_not_exist.txt: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			result, resultErr := fileutils.ReadLinesFromFile(tt.path)
			if len(tt.expectedErr) > 0 {
				assert.Error(resultErr)
				assert.EqualError(resultErr, tt.expectedErr)
				return
			}

			assert.NoError(resultErr)
			assert.Equal(tt.expected, result)
		})
	}
}
