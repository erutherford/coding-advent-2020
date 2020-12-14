package fileutils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadLinesFromFile(path string) ([]string, error) {
	cleanPath := filepath.Clean(path)

	file, err := os.Open(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, file.Close()
}
