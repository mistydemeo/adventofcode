package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type Form map[string]bool

func CountResponses(lines []string) []int {
	forms := make([]int, 0)
	form := make(Form)

	for _, line := range lines {
		if line == "" {
			forms = append(forms, len(form))
			form = make(Form)
			continue
		}

		for _, c := range line {
			char := string(c)
			form[char] = true
		}
	}

	return forms
}

func IdentifyCommonValues(forms []Form) []string {
	counts := make(map[string]int)
	for _, form := range forms {
		for key, _ := range form {
			// initialize if necessary
			_, exists := counts[key]
			if !exists {
				counts[key] = 1
			} else {
				counts[key]++
			}
		}
	}

	// Now do a second loop to identify values in `counts` where
	// the number of responses equals the number of forms
	responded_by_all := make([]string, 0)
	for key, count := range counts {
		if count == len(forms) {
			responded_by_all = append(responded_by_all, key)
		}
	}

	return responded_by_all
}

func CountUnanimousResponses(lines []string) []int {
	// Tracks an array of all forms for a given responses
	forms := make([]Form, 0)
	// Each individual response
	form := make(Form)
	unanimous_responses := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			unanimous := IdentifyCommonValues(forms)
			unanimous_responses = append(unanimous_responses, len(unanimous))

			forms = make([]Form, 0)
			form = make(Form)
			continue
		}

		for _, c := range line {
			char := string(c)
			form[char] = true
		}

		forms = append(forms, form)
		form = make(Form)
	}

	return unanimous_responses
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")
	counts := CountResponses(lines)
	sum := 0
	for _, count := range counts {
		sum += count
	}
	println(sum)

	unanimous := CountUnanimousResponses(lines)
	sum = 0
	for _, count := range unanimous {
		sum += count
	}
	println(sum)
}
