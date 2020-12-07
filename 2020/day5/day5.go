package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type seat struct {
	row, col int
}

func (s seat) ID() int {
	return s.row*8 + s.col
}

type seats []seat

// sort.Interface methods
var _ sort.Interface = seats{}

func (s seats) Len() int           { return len(s) }
func (s seats) Less(i, j int) bool { return s[i].ID() < s[j].ID() }
func (s seats) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	scanner, close := openFile("input.txt")
	defer close()

	var seats []seat
	for scanner.Scan() {
		seat := codeToSeat(scanner.Text())
		seats = append(seats, seat)
	}
	fmt.Println("p1:", p1(seats))
	fmt.Println("p2:", p2(seats))
}

func codeToSeat(code string) seat {
	rowCode := code[:7]
	colCode := code[7:]

	return seat{
		row: logSearch(0, 127, 'B', rowCode),
		col: logSearch(0, 7, 'R', colCode),
	}
}

func logSearch(min, max int, upChar byte, code string) int {
	for min != max {
		char := code[0]
		if char == upChar {
			min = (max-min)/2 + min + 1
		} else {
			max = (max-min)/2 + min
		}
		code = code[1:]
	}
	return min
}

func p1(seats []seat) (max int) {
	for _, s := range seats {
		if id := s.ID(); id > max {
			max = id
		}
	}
	return max
}

func p2(seats seats) int {
	sort.Sort(seats)

	// not first or last, adjust one on both ends
	l := 1
	r := len(seats) - 2
	offset := seats[0].ID()

	for l != r {
		mid := (l-r)/2 + r
		midSeat := seats[mid]

		if midSeat.ID()-offset == mid {
			l = mid + 1
		} else {
			r = mid
		}

		if midSeat.ID()-seats[mid-1].ID() > 1 {
			return midSeat.ID() - 1
		}
	}

	return 0
}
