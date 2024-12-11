package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(stringContent string) {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(stringContent, -1)

	var results [][]int
	for _, match := range matches {
		if len(match) > 2 {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			results = append(results, []int{a, b})
		}
	}

	result := 0
	for _, pair := range results {
		result += pair[0] * pair[1]
	}

	fmt.Println(results)
	fmt.Println(result)
}

func part2(stringContent string) {
	shouldRunPattern := `(?s)(?:^|do\(\))(.*?)(?:$|don't\(\))`
	re := regexp.MustCompile(shouldRunPattern)
	enabledStrings := re.FindAllStringSubmatch(stringContent, -1)

	multiplicationPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	multiplicationRe := regexp.MustCompile(multiplicationPattern)

	var results [][]int
	for _, match := range enabledStrings {
		if len(match) > 1 {
			enabledString := match[1]
			fmt.Println(enabledString)
			matchedMultiplications := multiplicationRe.FindAllStringSubmatch(enabledString, -1)
			for _, matchedMultiplication := range matchedMultiplications {
				if len(matchedMultiplication) > 2 {
					a, _ := strconv.Atoi(matchedMultiplication[1])
					b, _ := strconv.Atoi(matchedMultiplication[2])
					results = append(results, []int{a, b})
				}
			}
		}
	}

	result := 0
	for _, pair := range results {
		result += pair[0] * pair[1]
	}

	fmt.Println(results)
	fmt.Println(result)
}

func main() {
	content, _ := os.ReadFile("data")
	stringContent := string(content)

	fmt.Println(stringContent)
	part1(stringContent)
	fmt.Println()
	part2(stringContent)
}
