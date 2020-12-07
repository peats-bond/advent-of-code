package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	scanner, close := openFile("input.txt")
	defer close()

	var groups []string
	var cur string

	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			cur += line + " "
		} else if cur != "" {
			groups = append(groups, cur)
			cur = ""
		}
	}
	if cur != "" {
		groups = append(groups, cur)
	}

	fmt.Println("p1:", p1(groups))
	fmt.Println("p2:", p2(groups))
}

func p1(groups []string) (counts int) {
	for _, groupAnswers := range groups {
		m := make(map[rune]struct{})
		for _, ans := range groupAnswers {
			if ans != ' ' {
				m[ans] = struct{}{}
			}
		}
		counts += len(m)
	}
	return counts
}

func p2(groups []string) (counts int) {
	for _, groupAnswers := range groups {
		byPerson := strings.Split(strings.TrimSpace(groupAnswers), " ")

		m := personAnswers(byPerson[0])
		for _, ans := range byPerson[1:] {
			m = intersect(m, personAnswers(ans))
		}

		counts += len(m)
	}

	return counts
}

func personAnswers(pAnswers string) map[rune]struct{} {
	result := make(map[rune]struct{})
	for _, ans := range pAnswers {
		result[ans] = struct{}{}
	}
	return result
}

func intersect(m1, m2 map[rune]struct{}) map[rune]struct{} {
	intersection := make(map[rune]struct{}, len(m1))
	for k := range m1 {
		if _, found := m2[k]; found {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}
