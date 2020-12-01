package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func StringsToInt(input []string) ([]int, error) {
	var output []int
	for _, string := range input {
		i, err := strconv.Atoi(string)
		if err != nil {
			return nil, err
		}

		output = append(output, i)
	}

	return output, nil
}

func DoubleIterationSum(numbers []int) int {
	// yeah, a double iteration isn't especially fast, but at least
	// we can break early
	for _, i1 := range numbers {
		for _, i2 := range numbers {
			if i1+i2 == 2020 {
				return i1 * i2
				break
			}
		}
	}

	return 0
}

func TripleIterationSum(numbers []int) int {
	for _, i1 := range numbers {
		for _, i2 := range numbers {
			for _, i3 := range numbers {
				if i1+i2+i3 == 2020 {
					return i1 * i2 * i3
					break
				}
			}
		}
	}

	return 0
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	string_numbers := strings.Split(string(data), "\n")
	numbers, err := StringsToInt(string_numbers)
	if err != nil {
		log.Fatal("Error parsing input:", err)
	}

	var output int
	if len(os.Args) > 1 && os.Args[1] == "--three" {
		output = TripleIterationSum(numbers)
	} else {
		output = DoubleIterationSum(numbers)
	}

	if output != 0 {
		println("Result:", output)
	} else {
		println("Unable to find the result!")
	}
}
