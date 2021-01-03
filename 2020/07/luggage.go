package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	colour string
	count int
}

type Rule struct {
	colour string
	bags []Bag
}

func ParseBagCondition(text string) Bag {
	re := regexp.MustCompile(`(\d+) (.+) bags?`)
	match := re.FindStringSubmatch(text)
	// TODO: handle this
	number, _ := strconv.Atoi(match[1])
	return Bag{
		colour: match[2],
		count: number,
	}
}

func ParseBagConditions(text string) []Bag {
	conditions := strings.Split(text, ", ")
	bags := make([]Bag, 0)
	for _, condition := range conditions {
		bags = append(bags, ParseBagCondition(condition))
	}

	return bags
}

func ParseRule(text string) Rule {
	// Strip the trailing period
	text = text[0:len(text) - 1]
	split := strings.Split(text, " bags contain ")
	colour := split[0]
	conditions := split[1]
	bags := make([]Bag, 0)
	// If the text is "no other bags", then we parse no rules
	if conditions != "no other bags" {
		bags = ParseBagConditions(conditions)
	}

	return Rule{
		colour: colour,
		bags: bags,
	}
}

func ParseRules(text []string) []Rule {
	rules := make([]Rule, 0)
	for _, line := range text {
		rule := ParseRule(line)
		rules = append(rules, rule)
	}

	return rules
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")
	rules := ParseRules(lines)

	count := 0
	for _, rule := range rules {
		for _, bag := range rule.bags {
			if bag.colour == "shiny gold" {
				count++
			}
		}
	}

	println(count)
}
