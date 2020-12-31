package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func RideSlope(right_step int, down_step int, lines []string) int {
	height := len(lines)
	width := len(lines[0])

	trees := 0
	step := 0
	for i := down_step; i < height; i += down_step {
		step++

		// Borrowed from https://github.com/coingraham/adventofcode/blob/master/2020/day3.py#L35
		index := (step * right_step) % width
		if string(lines[i][index]) == "#" {
			trees++
		}
	}

	return trees
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")

	trees := RideSlope(3, 1, lines)

	println(trees)

	multiplied := RideSlope(1, 1, lines) *
		RideSlope(3, 1, lines) *
		RideSlope(5, 1, lines) * 
		RideSlope(7, 1, lines) * 
		RideSlope(1, 2, lines)

	println(multiplied)
}
