package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func CalculateRange(boarding_pass string, low int, high int) []int {
	result := []int{low, high}

	for _, c := range boarding_pass {
		switch string(c) {
		case "F", "L":
			result[1] = (result[0] + result[1]) / 2
		case "B", "R":
			result[0] = (result[0] + result[1]) / 2
		}
	}

	return result
}

func DecodePass(boarding_pass string) []int {
	horizontal_range := boarding_pass[0:7]
	vertical_range := boarding_pass[7:10]

	row_range := CalculateRange(horizontal_range, 0, 127)
	column_range := CalculateRange(vertical_range, 0, 7)

	horizontal_seat := row_range[1]
	vertical_seat := column_range[1]

	return []int{horizontal_seat, vertical_seat}
}

func CalculateBoardingPassID(seat []int) int {
	return seat[0]*8 + seat[1]
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")

	highest := 0
	seen := []int{}
	for _, boarding_pass := range lines {
		seat := DecodePass(boarding_pass)
		boarding_id := CalculateBoardingPassID(seat)
		seen = append(seen, boarding_id)
		if boarding_id > highest {
			highest = boarding_id
		}
	}

	println(highest)

	// Next, check for any missing IDs. We'll start by sorting so we can
	// identify missing IDs positionally.
	sort.Ints(seen)
	for i, id := range seen {
		// We won't be at the very front or the very end
		if i == 0 || i == len(seen)-1 {
			continue
		}

		if seen[i-1] != id-1 {
			println(id-1)
			break
		}

		if seen[i+1] != id+1 {
			 println(id+1)
			 break
		}
	}
}
