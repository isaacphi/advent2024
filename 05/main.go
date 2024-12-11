package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func pagesSatisfyRule(rule []string, pages []string) bool {
	foundSecondPage := false
	for _, page := range pages {
		if page == rule[1] {
			foundSecondPage = true
			continue
		}
		if foundSecondPage && page == rule[0] {
			return false
		}
	}
	return true
}

func pagesSatisfyRules(rules [][]string, pages []string) bool {
	for _, rule := range rules {
		isValid := pagesSatisfyRule(rule, pages)
		if !isValid {
			return false
		}
	}
	return true
}

func part1(rules [][]string, updates [][]string) {
	sum := 0
	for _, pages := range updates {
		isValid := pagesSatisfyRules(rules, pages)
		if isValid {
			middleNumber, _ := strconv.Atoi(pages[len(pages)/2])
			sum += middleNumber
		}
	}
	fmt.Println(sum)
}

// Part 2

func fixPages(rules [][]string, pages []string) {
	for _, rule := range rules {
		isValid := pagesSatisfyRule(rule, pages)
		if !isValid {
			p1 := slices.Index(pages, rule[0])
			p2 := slices.Index(pages, rule[1])
			pages[p1], pages[p2] = pages[p2], pages[p1]
			fixPages(rules, pages)
		}
	}
}

func part2(rules [][]string, updates [][]string) {
	sum := 0
	for _, pages := range updates {
		isValid := pagesSatisfyRules(rules, pages)
		if !isValid {
			fixPages(rules, pages)
			middleNumber, _ := strconv.Atoi(pages[len(pages)/2])
			sum += middleNumber
		}
	}
	fmt.Println(sum)
}

func main() {
	// file, _ := os.Open("data")
	file, _ := os.Open("test_data")

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rules [][]string
	var updates [][]string

	isParsingRules := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isParsingRules = false
			continue
		}
		if isParsingRules {
			rules = append(rules, strings.Split(line, "|"))
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	// part1(rules, updates)
	part2(rules, updates)
}
