package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/erutherford/coding-advent-2020/pkg/fileutils"
)

/**
  Without Sorting and Exiting Early
  combinations seen: 7068
  2020/12/13 22:02:04 indexes are: 35, 103
  2020/12/13 22:02:04 values are: 399, 1621
  2020/12/13 22:02:04 product is: 646779
  combinations seen: 2220360
  2020/12/13 22:02:04 indexes are: 56, 70, 187
  2020/12/13 22:02:04 values are: 591, 1021, 408
  2020/12/13 22:02:04 product is: 246191688
*/

/**
  WIth Sorting and Exiting Early
  combinations seen: 274
  2020/12/13 22:04:38 indexes are: 1, 108
  2020/12/13 22:04:38 values are: 399, 1621
  2020/12/13 22:04:38 product is: 646779
  combinations seen: 1411
  2020/12/13 22:04:38 indexes are: 2, 6, 11
  2020/12/13 22:04:38 values are: 408, 591, 1021
  2020/12/13 22:04:38 product is: 246191688
*/

type entries []int64

func (e entries) Len() int           { return len(e) }
func (e entries) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e entries) Less(i, j int) bool { return e[i] < e[j] }

func getInputData(inputPath string) (entries, error) {
	lines, err := fileutils.ReadLinesFromFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var input []int64
	for _, line := range lines {
		lineEntry, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing int from line: %w", err)
		}

		input = append(input, lineEntry)
	}

	return input, nil
}

var inputPath = flag.String("i", "", "input path for file")

func main() {
	flag.Parse()

	path := filepath.Clean(*inputPath)
	input, err := getInputData(path)
	if err != nil {
		log.Fatalf("error occurred getting input data: %#v", err)
	}
	sort.Sort(input)

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
	combinationsSeen := 0
	for firstIdx, firstEntry := range entries {
		for secondIdx, secondEntry := range entries {
			if firstIdx == secondIdx {
				continue
			}
			combinationsSeen += 1
			if firstEntry+secondEntry > 2020 {
				break
			}

			if firstEntry+secondEntry == 2020 {
				fmt.Printf("combinations seen: %d\n", combinationsSeen)
				return firstIdx, secondIdx
			}
		}
	}
	return 0, 0
}

func findThreeEntriesWithSum(entries []int64) (int, int, int) {
	combinationsSeen := 0
	for firstIdx, firstEntry := range entries {
		for secondIdx, secondEntry := range entries {
			if secondIdx == firstIdx {
				continue
			}
			if firstEntry+secondEntry > 2020 {
				break
			}

			for thirdIdx, thirdEntry := range entries {
				if thirdEntry == firstEntry || thirdEntry == secondEntry {
					continue
				}
				combinationsSeen += 1

				if firstEntry+secondEntry+thirdEntry > 2020 {
					break
				}

				if firstEntry+secondEntry+thirdEntry == 2020 {
					fmt.Printf("combinations seen: %d\n", combinationsSeen)
					return firstIdx, secondIdx, thirdIdx
				}
			}
		}
	}
	return 0, 0, 0
}
