package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
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

func (self Rule) passwordValidByPosition(password string) bool {
	first := self.minimum - 1
	second := self.maximum - 1
	if string(password[first]) == self.character && string(password[second]) == self.character {
		return false
	}

	return string(password[first]) == self.character || string(password[second]) == self.character
}

func ParseRuleDefinition(definition string) (*Rule, error) {
	re := regexp.MustCompile(`(\d+)-(\d+) (\S)`)
	match := re.FindStringSubmatch(definition)
	char := match[3]

	var minimum int
	var maximum int
	var err error
	minimum, err = strconv.Atoi(match[1])
	if err != nil {
		return nil, err
	}

	maximum, err = strconv.Atoi(match[2])
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
	position_method := len(os.Args) > 1 && os.Args[1] == "--correct"

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

		if position_method {
			if rule.passwordValidByPosition(password) {
				valid_passwords++
			}
		} else {
			if rule.passwordValid(password) {
				valid_passwords++
			}
		}
	}

	println(valid_passwords)
}
