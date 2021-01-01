package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

func ParsePassports(lines []string) []Passport {
	passports := make([]Passport, 0)
	passport := make(Passport)
	for _, line := range lines {
		if line == "" {
			passports = append(passports, passport)
			passport = make(Passport)
			continue
		}

		passport_segments := strings.Split(line, " ")
		for _, segment := range passport_segments {
			split := strings.Split(segment, ":")
			passport[split[0]] = split[1]
		}
	}

	return passports
}

func RangeValid(value string, minimum int, maximum int) bool {
	number, err := strconv.Atoi(value)
	// Swallow errors for now, maybe change this later
	if err != nil {
		return false
	}

	if number < minimum || number > maximum {
		return false
	}

	return true
}

func HgtValid(value string) bool {
	length := len(value)
	unit := value[length - 2:length]
	number := value[0:length - 2]
	if unit == "cm" {
		return RangeValid(number, 150, 193)
	} else if unit == "in" {
		return RangeValid(number, 59, 76)
	} else {
		return false
	}
}

func HclValid(value string) bool {
	re := regexp.MustCompile(`#[0-9a-fA-F]{6}`)
	return re.Match([]byte(value))
}

func EclValid(value string) bool {
	switch value {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}

	return false
}

func PidValid(value string) bool {
	// Doesn't matter what the number is, but it has to be parseable as a number
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	// String length has to be 9, regardless of number value
	return len(value) == 9
}

func PassportFieldsValid(passport Passport) bool {
	required_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range required_fields {
		_, field_exists := passport[field]
		if !field_exists {
			return false
		}
	}

	return true
}

func PassportValid(passport Passport) bool {
	required_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range required_fields {
		value, field_exists := passport[field]
		if !field_exists {
			return false
		}

		result := true

		switch field {
		case "byr":
			result = RangeValid(value, 1920, 2002)
		case "iyr":
			result = RangeValid(value, 2010, 2020)
		case "eyr":
			result = RangeValid(value, 2020, 2030)
		case "hgt":
			result = HgtValid(value)
		case "hcl":
			result = HclValid(value)
		case "ecl":
			result = EclValid(value)
		case "pid":
			result = PidValid(value)
		}


		if !result {
			return false
		}
	}

	return true
}

func CountValidPassportFields(passports []Passport) int {
	valid := 0
	for _, passport := range passports {
		if PassportFieldsValid(passport) {
			valid++
		}
	}

	return valid
}

func CountValidPassports(passports []Passport) int {
	valid := 0
	for _, passport := range passports {
		if PassportValid(passport) {
			valid++
		}
	}

	return valid
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")

	passports := ParsePassports(lines)
	println(CountValidPassportFields(passports))
	println(CountValidPassports(passports))
}
