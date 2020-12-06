package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

var reqFields = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
	// "cid": true, ignore intentionally
}

func main() {
	scanner, close := openFile("input.txt")
	defer close()

	var passports []string
	var cur string
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			cur += line + " "
		} else {
			passports = append(passports, strings.TrimSpace(cur))
			cur = ""
		}
	}
	if cur != "" {
		passports = append(passports, strings.TrimSpace(cur))
	}

	fmt.Println("p1:", p1(passports))
	fmt.Println("p2:", p2(passports))
}

func p1(passports []string) (valid int) {
	for _, passport := range passports {
		gotFields := make(map[string]bool, len(reqFields))

		for _, gotPair := range strings.Split(passport, " ") {
			gotField := strings.Split(gotPair, ":")[0]

			// only add known required fields
			if _, isReqField := reqFields[gotField]; isReqField {
				gotFields[gotField] = true
			}
		}
		if len(gotFields) == len(reqFields) {
			valid++
		}
	}

	return valid
}

func p2(passports []string) (valid int) {
	for _, passport := range passports {
		gotFields := make(map[string]bool, len(reqFields))

		for _, gotPair := range strings.Split(passport, " ") {
			pair := strings.Split(gotPair, ":")
			key, val := pair[0], pair[1]

			// only add known required fields
			if _, isReqField := reqFields[key]; isReqField {
				if validField(key, val) {
					gotFields[key] = true
				}
			}
		}
		if len(gotFields) == len(reqFields) {
			valid++
		}
	}

	return valid
}

func validField(key, val string) bool {
	switch key {
	case "byr":
		return validNumIncl(val, 1920, 2002)
	case "iyr":
		return validNumIncl(val, 2010, 2020)
	case "eyr":
		return validNumIncl(val, 2020, 2030)
	case "hgt":
		return validHeight(val)
	case "hcl":
		return validHairColour(val)
	case "ecl":
		return validEyeColour(val)
	case "pid":
		return validPassportID(val)
	}
	return false
}

func validNumIncl(val string, min, max int) bool {
	if num, err := strconv.Atoi(val); err == nil {
		return num >= min && num <= max
	}
	return false
}

func validHeight(val string) bool {
	if strings.HasSuffix(val, "cm") {
		return validNumIncl(strings.TrimSuffix(val, "cm"), 150, 193)
	}
	if strings.HasSuffix(val, "in") {
		return validNumIncl(strings.TrimSuffix(val, "in"), 59, 76)
	}
	return false
}

var hairRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)

func validHairColour(val string) bool {
	return hairRegex.MatchString(val)
}

func validEyeColour(val string) bool {
	switch val {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

var pIDRegex = regexp.MustCompile(`^[0-9]{9}$`)

func validPassportID(val string) bool {
	return pIDRegex.MatchString(val)
}
