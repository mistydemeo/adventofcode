package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Rule struct {
	minimum   int
	maximum   int
	character string
}

func (self Rule) passwordValid(password string) bool {
	char_count := strings.Count(password, self.character)
	return char_count >= self.minimum && char_count <= self.maximum
}

func ParseRuleDefinition(definition string) (*Rule, error) {
	// TODO refactor this to use regular expressions
	// First, split by " " to fetch whatever the character is
	split_definition := strings.Split(definition, " ")
	char := split_definition[1]

	// Next, parse out the ranges
	ranges := strings.Split(split_definition[0], "-")
	var minimum int
	var maximum int
	var err error
	minimum, err = strconv.Atoi(ranges[0])
	if err != nil {
		return nil, err
	}

	maximum, err = strconv.Atoi(ranges[1])
	if err != nil {
		return nil, err
	}

	return &Rule{
		minimum:   minimum,
		maximum:   maximum,
		character: char,
	}, nil
}

func ParseLine(input string) (*Rule, string, error) {
	split := strings.Split(input, ": ")
	rule, err := ParseRuleDefinition(split[0])
	if err != nil {
		return nil, "", err
	}

	var password = split[1]

	return rule, password, nil
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")

	valid_passwords := 0
	for _, line := range lines {
		rule, password, err := ParseLine(line)
		if err != nil {
			log.Fatal("Error parsing input:", err)
		}

		if rule.passwordValid(password) {
			valid_passwords++
		}
	}

	println(valid_passwords)
}
