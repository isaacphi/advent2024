package main

import (
	"fmt"
	"os"
	"strings"
)

type spot struct {
	val          string
	stepsToEnd   int
	stepsToStart int
}

type vec struct {
	i, j int
}

func printGrid(grid [][]spot) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val.val)
		}
		fmt.Println()
	}
}

func visit(i, j int, grid [][]spot, steps int, fromEnd bool) {
	L := len(grid)
	spot := &grid[i][j]
	if spot.val != "#" && (fromEnd && (spot.stepsToEnd == -1 || spot.stepsToEnd > steps) || (!fromEnd && (spot.stepsToStart == -1 || spot.stepsToStart > steps))) {
		if fromEnd {
			spot.stepsToEnd = steps
		} else {
			spot.stepsToStart = steps
		}
		if i > 0 {
			visit(i-1, j, grid, steps+1, fromEnd)
		}
		if i < L-1 {
			visit(i+1, j, grid, steps+1, fromEnd)
		}
		if j > 0 {
			visit(i, j-1, grid, steps+1, fromEnd)
		}
		if j < L-1 {
			visit(i, j+1, grid, steps+1, fromEnd)
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("data")
	input := string(data)

	H := len(strings.Split(input, "\n")) - 1
	var L int
	var start vec
	var end vec

	grid := make([][]spot, H)
	for i, row := range strings.Split(input, "\n") {
		if row == "" {
			break
		}
		L = len(row)
		grid[i] = make([]spot, L)
		for j, val := range row {
			if val == 'S' {
				start = vec{i, j}
			}
			if val == 'E' {
				end = vec{i, j}
			}
			grid[i][j] = spot{val: string(val), stepsToEnd: -1, stepsToStart: -1}
		}
	}

	visit(end.i, end.j, grid, 0, true)
	visit(start.i, start.j, grid, 0, false)

	M := grid[start.i][start.j].stepsToEnd
	diff := 100
	maxCheatLen := 20

	count := 0
	for i, row := range grid {
		for j, s := range row {
			for di := -maxCheatLen; di <= maxCheatLen; di++ {
				for dj := -maxCheatLen; dj <= maxCheatLen; dj++ {
					if abs(di)+abs(dj) <= maxCheatLen {
						if i+di >= 0 && i+di < H && j+dj >= 0 && j+dj < L {
							cheatDistance := getDistance(s, grid[i+di][j+dj], abs(dj)+abs(di))
							if cheatDistance <= M-diff {
								count++
							}
						}
					}
				}
			}
		}
	}

	// printGrid(grid)
	fmt.Println(count)
}

func getDistance(s1, s2 spot, cheatLen int) int {
	if s1.stepsToStart >= 0 && s2.stepsToEnd >= 0 {
		return s1.stepsToStart + s2.stepsToEnd + cheatLen
	} else {
		return 99999999
	}
}
