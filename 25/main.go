package main

import (
	"fmt"
	"os"
	"strings"
)

type key [5]int
type lock [5]int

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("data")
	input := string(data)

	keys := make([]key, 0)
	locks := make([]lock, 0)
	for _, diagram := range strings.Split(input, "\n\n") {
		grid := make([][]string, 0)
		for i, row := range strings.Fields(diagram) {
			grid = append(grid, make([]string, 0))
			for _, val := range row {
				grid[i] = append(grid[i], string(val))
			}
		}
		if grid[0][0] == "#" {
			// l
			var l lock
			for j := 0; j < 5; j++ {
				height := 0
				for grid[len(grid)-1-height][j] == "." {
					height++
				}
				l[j] = 5 - height + 1
			}
			locks = append(locks, l)
		} else {
			// k
			var k key
			for j := 0; j < 5; j++ {
				height := 0
				for grid[len(grid)-1-height][j] == "#" {
					height++
				}
				k[j] = height - 1
			}
			keys = append(keys, k)
		}
	}

	// part 1
	count := 0
	for _, l := range locks {
	outerloop:
		for _, k := range keys {
			for i := range l {
				if k[i]+l[i] > 5 {
					// fmt.Println(l, k, "overlap")
					continue outerloop
				}
			}
			// fmt.Println(k, l, "fit")
			count++
		}
	}
	fmt.Println(count)

	// part 2

}
