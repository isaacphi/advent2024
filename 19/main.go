package main

import (
	"fmt"
	"os"
	"strings"
)

var possibleDesigns map[string]bool
var impossibleDesigns map[string]bool

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("data")
	input := string(data)

	var towels []string
	designs := make([]string, 0)
	gettingTowels := true
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			gettingTowels = false
			continue
		}
		if gettingTowels {
			towels = strings.Split(row, ", ")
		} else {
			designs = append(designs, row)
		}
	}

	possibleDesigns = make(map[string]bool)
	impossibleDesigns = make(map[string]bool)

	fmt.Println(towels)

	count := 0
	for _, d := range designs {
		designOptions := make(map[string]int)
		fmt.Println(d)
		p := isPossible(d, towels)
		// if p {
		// 	count++
		// }
		c := numOptions(d, towels, designOptions)
		fmt.Println(p, c)
		count += c
	}
	fmt.Println(count)
}

func numOptions(d string, towels []string, designOptions map[string]int) int {
	count := 0

	if c, exists := designOptions[d]; exists {
		return c
	}
	if _, exists := impossibleDesigns[d]; exists {
		return 0
	}

	for _, t := range towels {
		if d == t {
			count++
		}
		if len(t) > len(d) {
			continue
		}
		if d[:len(t)] == t {
			count += numOptions(d[len(t):], towels, designOptions)
		}
	}
	if count == 0 {
		impossibleDesigns[d] = true
	}
	designOptions[d] += count
	return count
}

func isPossible(d string, towels []string) bool {
	if _, exists := possibleDesigns[d]; exists {
		return true
	}
	if _, exists := impossibleDesigns[d]; exists {
		return false
	}

	for _, t := range towels {
		if d == t {
			possibleDesigns[d] = true
			return true
		}
		if len(t) > len(d) {
			continue
		}
		if d[:len(t)] == t {
			p := isPossible(d[len(t):], towels)
			if p {
				possibleDesigns[d] = true
				return true
			}
		}
	}

	impossibleDesigns[d] = true
	return false
}
