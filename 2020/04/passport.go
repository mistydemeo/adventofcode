package main

import (
	"io/ioutil"
	"log"
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

func PassportValid(passport Passport) bool {
	required_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range required_fields {
		_, field_exists := passport[field]
		if !field_exists {
			return false
		}
	}

	return true
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
	println(CountValidPassports(passports))
}