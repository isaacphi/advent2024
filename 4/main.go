package main

import (
	"fmt"
	"os"
)

func part1(wordSearch string) {
	fmt.Println(wordSearch)
}

func part2() {}

func main() {
	content, _ := os.ReadFile("test_data")
	wordSearch := string(content)

	part1(wordSearch)
}
