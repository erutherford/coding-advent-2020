package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newPasswordPolicy(t *testing.T) {
	tests := []struct {
		name           string
		policyArg      string
		expectedPolicy sledPasswordPolicy
		expectedErr    string
	}{
		{
			name:      "should return the expected policy with a proper argument",
			policyArg: "1-3 a",
			expectedPolicy: sledPasswordPolicy{
				min:  1,
				max:  3,
				char: "a",
			},
			expectedErr: "",
		},
		{
			name:      "should hand double digit numbers",
			policyArg: "10-19 c",
			expectedPolicy: sledPasswordPolicy{
				min:  10,
				max:  19,
				char: "c",
			},
			expectedErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			result, resultErr := newSledPasswordPolicy(tt.policyArg)
			if tt.expectedErr != "" {
				assert.Error(resultErr)
				assert.EqualError(resultErr, tt.expectedErr)
				return
			}

			assert.NoError(resultErr)
			assert.Equal(tt.expectedPolicy, result)
		})
	}
}

func Test_sledPasswordPolicy_Validate(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		policy      sledPasswordPolicy
		expectedErr string
	}{
		{
			name:     "should return no errors if the password meets the policy requirements",
			password: "applesabcCCC",
			policy: sledPasswordPolicy{
				min:  1,
				max:  3,
				char: "c",
			},
			expectedErr: "",
		},
		{
			name:     "should handle upper case characters correctly",
			password: "applesabCCCccc",
			policy: sledPasswordPolicy{
				min:  1,
				max:  3,
				char: "C",
			},
			expectedErr: "",
		},
		{
			name:     "should handle symbols correctly",
			password: "applesab%%",
			policy: sledPasswordPolicy{
				min:  1,
				max:  3,
				char: "%",
			},
			expectedErr: "",
		},
		{
			name:     "first example should pass",
			password: "abcde",
			policy: sledPasswordPolicy{
				min:  1,
				max:  3,
				char: "a",
			},
			expectedErr: "",
		},
		{
			name:     "second example should fail",
			password: "cdefg",
			policy: sledPasswordPolicy{
				min:  1,
				max:  3,
				char: "b",
			},
			expectedErr: "password does not meet policy requirements",
		},
		{
			name:     "third example should pass",
			password: "ccccccccc",
			policy: sledPasswordPolicy{
				min:  2,
				max:  9,
				char: "c",
			},
			expectedErr: "",
		},
		{
			name:     "should throw an error if password violates minimum requirements",
			password: "pwordisthebest",
			policy: sledPasswordPolicy{
				min:  1,
				max:  3,
				char: "a",
			},
			expectedErr: "password does not meet policy requirements",
		},
		{
			name:     "should throw an error if password violates maximum requirements",
			password: "applesarealwaysawesome",
			policy: sledPasswordPolicy{
				min:  1,
				max:  3,
				char: "a",
			},
			expectedErr: "password does not meet policy requirements",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			resultErr := tt.policy.Validate(tt.password)
			if len(tt.expectedErr) > 0 {
				assert.Error(resultErr)
				assert.EqualError(resultErr, tt.expectedErr)
				return
			}
			assert.NoError(resultErr)
		})
	}
}

func Test_tobagganPasswordPolicy_Validate(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		policy      tobagganPasswordPolicy
		expectedErr string
	}{
		{
			name:     "should return no errors if the password meets the policy requirements",
			password: "apclesabcCCC",
			policy: tobagganPasswordPolicy{
				firstPosition:  1,
				secondPosition: 3,
				char:           "c",
			},
			expectedErr: "",
		},
		{
			name:     "should handle upper case characters correctly",
			password: "CpplesabCCCccc",
			policy: tobagganPasswordPolicy{
				firstPosition:  1,
				secondPosition: 3,
				char:           "C",
			},
			expectedErr: "",
		},
		{
			name:     "should handle symbols correctly",
			password: "%p%lesab%%",
			policy: tobagganPasswordPolicy{
				firstPosition:  1,
				secondPosition: 4,
				char:           "%",
			},
			expectedErr: "",
		},
		{
			name:     "first example should pass",
			password: "abcde",
			policy: tobagganPasswordPolicy{
				firstPosition:  1,
				secondPosition: 3,
				char:           "a",
			},
			expectedErr: "",
		},
		{
			name:     "second example should fail",
			password: "cdefg",
			policy: tobagganPasswordPolicy{
				firstPosition:  1,
				secondPosition: 3,
				char:           "b",
			},
			expectedErr: "password does not meet policy requirements",
		},
		{
			name:     "third example should fail",
			password: "ccccccccc",
			policy: tobagganPasswordPolicy{
				firstPosition:  2,
				secondPosition: 9,
				char:           "c",
			},
			expectedErr: "password does not meet policy requirements",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			resultErr := tt.policy.Validate(tt.password)
			if len(tt.expectedErr) > 0 {
				assert.Error(resultErr)
				assert.EqualError(resultErr, tt.expectedErr)
				return
			}
			assert.NoError(resultErr)
		})
	}
}
