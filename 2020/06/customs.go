package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func CountResponses(lines []string) []int {
	forms := make([]int, 0)
	form := make(map[string]bool)

	for _, line := range lines {
		if line == "" {
			forms = append(forms, len(form))
			form = make(map[string]bool)
			continue
		}

		for _, c := range line {
			char := string(c)
			form[char] = true
		}
	}

	return forms
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
}
