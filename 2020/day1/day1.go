package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func openFile(fileName string) (*bufio.Scanner, func() error) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file), file.Close
}

const _sumTo = 2020

func main() {
	scanner, close := openFile("day1.txt")
	defer close()

	var ints []int

	// load ints
	for scanner.Scan() {
		got, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, got)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("twoSum result:", twoSum(ints))
	fmt.Println("threeSum result:", threeSum(ints))
}

func twoSum(ints []int) int {
	// store ints
	seen := make(map[int]bool, len(ints))
	for _, entry := range ints {
		seen[entry] = true
	}

	// find the two
	for _, entry := range ints {
		if _, found := seen[_sumTo-entry]; found {
			return (_sumTo - entry) * entry
		}
	}

	panic("not found")
}

func threeSum(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == _sumTo {
				return nums[i] * nums[j] * nums[k]
			}

			if sum < _sumTo {
				j++
			} else {
				k--
			}
		}
	}

	panic("not found")
}
