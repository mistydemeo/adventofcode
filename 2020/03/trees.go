package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")

	height := len(lines)
	width := len(lines[0])

	trees := 0
	step := 0
	for i := 1; i < height; i++ {
		step++

		// Borrowed from https://github.com/coingraham/adventofcode/blob/master/2020/day3.py#L35
		index := (step * 3) % width
		if string(lines[i][index]) == "#" {
			trees++
		}
	}

	println(trees)
}
