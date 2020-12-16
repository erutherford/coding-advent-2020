package main

import (
	"fmt"
	"strconv"
)

type PasswordValidators interface {
	Validate(password string) error
}

func newSledPasswordPolicy(policy string) (sledPasswordPolicy, error) {
	parsedPolicy := sledPasswordPolicy{}
	minFound := false
	maxFound := false
	lastIdx := 0
	for idx, char := range policy {
		if !minFound && string(char) == `-` {
			minFound = true
			min, err := strconv.ParseInt(policy[:idx], 10, 64)
			if err != nil {
				return parsedPolicy, fmt.Errorf("error parsing min: %w", err)
			}
			lastIdx = idx
			parsedPolicy.min = int(min)
		}

		if !maxFound && string(char) == ` ` {
			maxFound = true
			max, err := strconv.ParseInt(policy[lastIdx+1:idx], 10, 64)
			if err != nil {
				return parsedPolicy, fmt.Errorf("errror parsing max: %w", err)
			}
			lastIdx = idx
			parsedPolicy.max = int(max)
		}

		parsedPolicy.char = string(char)
	}

	return parsedPolicy, nil
}

type sledPasswordPolicy struct {
	min  int
	max  int
	char string
}

func (p sledPasswordPolicy) Validate(password string) error {
	if len(password) < p.min {
		return ErrInvalidPassword
	}

	charTracker := map[string]int{}
	for _, char := range password {
		normalizedChar := string(char)
		if _, ok := charTracker[normalizedChar]; !ok {
			charTracker[normalizedChar] = 0
		}
		charTracker[normalizedChar]++

		if charTracker[p.char] > p.max {
			return ErrInvalidPassword
		}
	}

	if charTracker[p.char] < p.min {
		return ErrInvalidPassword
	}

	return nil
}

func newTobbaganPasswordPolicy(policy string) (tobagganPasswordPolicy, error) {
	parsedPolicy := tobagganPasswordPolicy{}
	firstFound := false
	secondFound := false
	lastIdx := 0
	for idx, char := range policy {
		if !firstFound && string(char) == `-` {
			firstFound = true
			min, err := strconv.ParseInt(policy[:idx], 10, 64)
			if err != nil {
				return parsedPolicy, fmt.Errorf("error parsing min: %w", err)
			}
			lastIdx = idx
			parsedPolicy.firstPosition = int(min)
		}

		if !secondFound && string(char) == ` ` {
			secondFound = true
			max, err := strconv.ParseInt(policy[lastIdx+1:idx], 10, 64)
			if err != nil {
				return parsedPolicy, fmt.Errorf("errror parsing max: %w", err)
			}
			lastIdx = idx
			parsedPolicy.secondPosition = int(max)
		}

		parsedPolicy.char = string(char)
	}

	return parsedPolicy, nil
}

type tobagganPasswordPolicy struct {
	firstPosition  int
	secondPosition int
	char           string
}

func (p tobagganPasswordPolicy) Validate(password string) error {
	pwordBytes := []byte(password)
	matches := 0

	if len(pwordBytes) < p.secondPosition {
		return ErrInvalidPassword
	}

	if string(pwordBytes[p.firstPosition-1]) != p.char {
		matches++
	}

	if string(pwordBytes[p.secondPosition-1]) != p.char {
		matches++
	}

	if matches != 1 {
		return ErrInvalidPassword
	}

	return nil
}
