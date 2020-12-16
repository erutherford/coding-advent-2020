package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countValidPasswords(t *testing.T) {
	tests := []struct {
		name            string
		passwordEntries []string
		policyType      string
		expected        int
		expectedErr     string
	}{
		{
			name: "with proper input should return the expected response for sleds",
			passwordEntries: []string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			policyType:  "sled",
			expected:    2,
			expectedErr: "",
		},
		{
			name: "with proper input should return the expected response for tobaggans",
			passwordEntries: []string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			policyType:  "tobaggan",
			expected:    1,
			expectedErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			result, err := countValidPasswords(tt.passwordEntries, tt.policyType)
			if len(tt.expectedErr) > 0 {
				assert.Error(err)
				assert.EqualError(err, tt.expectedErr)
				return
			}
			assert.Equal(tt.expected, result)
		})
	}
}
