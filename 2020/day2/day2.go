package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type input struct {
	num1, num2 int
	char       string
	password   string
}

func main() {
	scanner, close := openFile("day2.txt")
	defer close()

	var inputs []input

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		minMax := strings.Split(line[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		char := strings.TrimSuffix(line[1], ":")
		password := line[2]

		inputs = append(inputs, input{
			num1:     min,
			num2:     max,
			char:     char,
			password: password,
		})
	}

	fmt.Println("valid passwords p1:", policy1(inputs))
	fmt.Println("valid passwords p2:", policy2(inputs))
}

func policy1(inputs []input) (numValid int) {
	for _, in := range inputs {
		count := strings.Count(in.password, in.char)
		if in.num1 <= count && count <= in.num2 {
			numValid++
		}
	}
	return
}

func policy2(inputs []input) (numValid int) {
	for _, in := range inputs {
		at1 := string(in.password[in.num1-1]) == in.char
		at2 := string(in.password[in.num2-1]) == in.char
		if at1 != at2 && (at1 || at2) {
			numValid++
		}
	}
	return
}
