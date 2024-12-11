package main

import (
	"fmt"
	"strconv"
	"strings"
)

func blink(rocks []string) []string {
	newRocks := make([]string, 0)
	var numLength int
	var rockVal int
	for _, rock := range rocks {
		numLength = len(rock)
		if rock == "0" {
			newRocks = append(newRocks, "1")
		} else if numLength%2 == 0 {
			left := rock[:numLength/2]
			right := rock[numLength/2:]
			i := 0
			for i < len(right)-1 && right[i] == '0' {
				i++
			}
			right = right[i:]
			newRocks = append(newRocks, left, right)
		} else {
			rockVal, _ = strconv.Atoi(rock)
			newRocks = append(newRocks, strconv.Itoa(2024*rockVal))
		}
	}
	return newRocks
}

var knownCounts map[string]map[int]int

func getCountAfter(rock string, blinks int) int {
	if blinks == 0 {
		return 1
	}
	if countsForRock, exists := knownCounts[rock]; exists {
		if countsForBlinks, exists := countsForRock[blinks]; exists {
			return countsForBlinks
		}
	}

	rocks := blink([]string{rock})
	count := 0
	for _, rock := range rocks {
		count += getCountAfter(rock, blinks-1)
	}
	if _, exists := knownCounts[rock]; !exists {
		knownCounts[rock] = make(map[int]int)
	}
	knownCounts[rock][blinks] = count
	return count
}

func main() {
	data := "125 17"                              // test data
	data = "8435 234 928434 14 0 7 92446 8992692" // puzzle data
	N := 75

	// rocks := strings.Fields(data)
	// fmt.Println(rocks)
	// for i := 0; i < N; i++ {
	// 	rocks = blink(rocks)
	// }
	// fmt.Println(len(rocks))

	rocks2 := strings.Fields(data)
	count := 0
	knownCounts = make(map[string]map[int]int)
	for _, rock := range rocks2 {
		count += getCountAfter(rock, N)
	}
	fmt.Println(count)
}
