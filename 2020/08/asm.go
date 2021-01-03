package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  int
}

func ParseInstruction(source string) (*Instruction, error) {
	operation := source[0:3]
	argument_string := source[4:len(source)]
	argument, err := strconv.Atoi(argument_string)
	if err != nil {
		return nil, err
	}

	instruction := Instruction{
		operation: operation,
		argument:  argument,
	}

	return &instruction, nil
}

func ParseInstructions(lines []string) ([]Instruction, error) {
	instructions := make([]Instruction, 0)
	for _, line := range lines {
		instruction, err := ParseInstruction(line)
		if err != nil {
			return nil, err
		} else {
			instructions = append(instructions, *instruction)
		}
	}

	return instructions, nil
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")
	instructions, err := ParseInstructions(lines)
	if err != nil {
		log.Fatal("Error parsing instructions:", err)
	}

	// It's emulation time!
	// Global variable that asm instructions will mutate. Starts at 0.
	accumulator := 0
	seen_instructions := make(map[int]bool)
	for i := 0; i < len(instructions); i++ {
		// Break without operating if already seen
		_, exists := seen_instructions[i]
		if exists {
			break
		}
		seen_instructions[i] = true

		instruction := instructions[i]
		switch instruction.operation {
		case "nop":
			continue
		case "acc":
			accumulator += instruction.argument
		case "jmp":
			i += instruction.argument - 1
		}
	}

	println(accumulator)
}
