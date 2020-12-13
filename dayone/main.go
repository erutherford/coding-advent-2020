package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
)

const separator = "\n"

func GetInputData(inputPath string) ([]int64, error) {
	fileContents, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fileErr := fmt.Errorf("error reading file: %#v", err)
		log.Println(fileErr)
		return nil, fileErr
	}

	var input []int64
	splitContents := bytes.Split(fileContents, []byte(separator))
	for _, m := range splitContents {
		massStr := string(m[:])

		mass, err := strconv.ParseInt(massStr, 10, 64)
		if err != nil {
			massErr := fmt.Errorf("error converting to int: %w", err)
			log.Println(massErr)
			return nil, massErr
		}

		input = append(input, mass)
	}

	return input, nil
}

var inputPath = flag.String("i", "", "input path for file")

func main() {
	flag.Parse()

	path := filepath.Clean(*inputPath)
	input, err := GetInputData(path)
	if err != nil {
		log.Fatalf("error occurred getting input data: %#v", err)
	}

	idxOne, idxTwo := findTwoEntriesWithSum(input)

	log.Printf("indexes are: %d, %d\n", idxOne, idxTwo)
	log.Printf("values are: %d, %d\n", input[idxOne], input[idxTwo])
	log.Printf("product is: %d\n", input[idxOne]*input[idxTwo])

	idxOne, idxTwo, idxThree := findThreeEntriesWithSum(input)

	log.Printf("indexes are: %d, %d, %d\n", idxOne, idxTwo, idxThree)
	log.Printf("values are: %d, %d, %d\n", input[idxOne], input[idxTwo], input[idxThree])
	log.Printf("product is: %d\n", input[idxOne]*input[idxTwo]*input[idxThree])
}

func findTwoEntriesWithSum(entries []int64) (int, int) {
	for pIdx, pEntry := range entries {
		for sIdx, sEntry := range entries {
			if pEntry+sEntry == 2020 {
				return pIdx, sIdx
			}
		}
	}
	return 0, 0
}

func findThreeEntriesWithSum(entries []int64) (int, int, int) {
	for pIdx, pEntry := range entries {
		for sIdx, sEntry := range entries {
			for tIdx, tEntry := range entries {
				if pEntry+sEntry+tEntry == 2020 {
					return pIdx, sIdx, tIdx
				}
			}
		}
	}
	return 0, 0, 0
}
