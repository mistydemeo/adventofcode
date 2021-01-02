package main

import (
	"io/ioutil"
	"log"
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
	for _, boarding_pass := range lines {
		seat := DecodePass(boarding_pass)
		boarding_id := CalculateBoardingPassID(seat)
		if boarding_id > highest {
			highest = boarding_id
		}
	}

	println(highest)
}
