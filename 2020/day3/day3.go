package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(fileName string) (*bufio.Scanner, func()) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	return scanner, func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}

type slope struct {
	right, down int
}

var p2Slopes = []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

func main() {
	scanner, close := openFile("input.txt")
	defer close()

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("p1:", p1(input))
	fmt.Println("p1:", p2(input, slope{3, 1}))

	resultP2 := 1
	for _, s := range p2Slopes {
		resultP2 *= p2(input, s)
	}
	fmt.Println("p2:", resultP2)
}

func p1(input []string) (trees int) {
	col := 3
	for _, line := range input[1:] {
		if got := line[col]; got == '#' {
			trees++
		}
		col = (col + 3) % len(line)
	}
	return trees
}

func p2(input []string, s slope) (trees int) {
	col := s.right
	for row, line := range input[s.down:] {
		if row%s.down != 0 {
			continue
		}

		if got := line[col]; got == '#' {
			trees++
		}
		col = (col + s.right) % len(line)
	}
	return trees
}
