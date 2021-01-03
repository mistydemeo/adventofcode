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
	count  int
}

type Rule struct {
	colour string
	bags   []Bag
}

func ParseBagCondition(text string) Bag {
	re := regexp.MustCompile(`(\d+) (.+) bags?`)
	match := re.FindStringSubmatch(text)
	// TODO: handle this
	number, _ := strconv.Atoi(match[1])
	return Bag{
		colour: match[2],
		count:  number,
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
	text = text[0 : len(text)-1]
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
		bags:   bags,
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

func BuildRuleMap(rules []Rule) map[string]Rule {
	rule_map := make(map[string]Rule)
	for _, rule := range rules {
		rule_map[rule.colour] = rule
	}

	return rule_map
}

func CheckBag(bag Bag) bool {
	if bag.colour == "shiny gold" {
		return true
	}

	return false
}

func RecursivelyCheckContents(bag Bag, rule_map map[string]Rule) bool {
	rule, _ := rule_map[bag.colour]
	for _, bag := range rule.bags {
		if CheckBag(bag) || RecursivelyCheckContents(bag, rule_map) {
			return true
		}
	}

	return false
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	lines := strings.Split(string(data), "\n")
	rules := ParseRules(lines)
	rule_map := BuildRuleMap(rules)

	count := 0
	for _, rule := range rules {
		valid := false
		for _, bag := range rule.bags {
			// We're looking not just for first-level bag contents, but
			// any possible depth of bag contents that could hypothetically
			// contain our bag.
			if CheckBag(bag) || RecursivelyCheckContents(bag, rule_map) {
				valid = true
			}
		}
		if valid {
			count++
		}
	}

	println(count)
}
