package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/erutherford/coding-advent-2020/pkg/fileutils"
)

var ErrInvalidPassword = errors.New("password does not meet policy requirements")

func getInputData(inputPath string) ([]string, error) {
	lines, err := fileutils.ReadLinesFromFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

func countValidPasswords(passwordEntries []string, policyType string) (int, error) {
	validCount := 0
	for _, entry := range passwordEntries {
		password, policy, err := parseEntry(entry, policyType)
		if err != nil {
			log.Print("error parsing entry", err)
			return validCount, err
		}
		if err := policy.Validate(password); err == nil {
			validCount++
		}
	}
	return validCount, nil
}

func parseEntry(entry, policyType string) (string, PasswordValidators, error) {
	strs := strings.Split(entry, ": ")
	if len(strs) != 2 {
		return "", sledPasswordPolicy{}, fmt.Errorf("invalid password entry")
	}

	var policy PasswordValidators
	var err error
	if policyType == "sled" {
		policy, err = newSledPasswordPolicy(strs[0])
		if err != nil {
			return "", sledPasswordPolicy{}, fmt.Errorf("unable to parse password policy: %w", err)
		}
	} else if policyType == "tobaggan" {
		policy, err = newTobbaganPasswordPolicy(strs[0])
		if err != nil {
			return "", tobagganPasswordPolicy{}, fmt.Errorf("unable to parse password policy: %w", err)
		}
	}

	return strs[1], policy, nil
}

var inputPath = flag.String("i", "", "input path for file")
var policyType = flag.String("t", "sled", "policy type for password validation (sled or toboggan)")

func main() {
	flag.Parse()

	path := filepath.Clean(*inputPath)
	input, err := getInputData(path)
	if err != nil {
		log.Fatalf("error occurred getting input data: %#v", err)
	}
	fmt.Printf("validating %d passwords\n", len(input))

	validPasswordCount, countErr := countValidPasswords(input, *policyType)
	if countErr != nil {
		log.Fatalf("error counting valid passwords: %#v", countErr)
	}

	fmt.Printf("validPasswordCount: %d\n", validPasswordCount)
}
