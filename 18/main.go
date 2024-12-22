package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type spot struct {
	val     string
	steps   int
	visited bool
}

func printGrid(grid [][]spot) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val.val, val.steps, " ")
		}
		fmt.Println()
	}
}

func visit(i, j int, grid [][]spot, steps int) {
	L := len(grid)
	spot := &grid[i][j]
	if spot.val != "#" && (!spot.visited || spot.steps > steps) {
		spot.visited = true
		spot.steps = steps
		if i > 0 {
			visit(i-1, j, grid, steps+1)
		}
		if i < L-1 {
			visit(i+1, j, grid, steps+1)
		}
		if j > 0 {
			visit(i, j-1, grid, steps+1)
		}
		if j < L-1 {
			visit(i, j+1, grid, steps+1)
		}
	}
}

func clearGrid(grid [][]spot) {
	for i, row := range grid {
		for j, val := range row {
			if val.val == "." {
				grid[i][j] = spot{val: "."}
			}
		}
	}
}

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("data")
	input := string(data)

	L := 7
	L = 71

	grid := make([][]spot, L)
	for i := 0; i < L; i++ {
		for j := 0; j < L; j++ {
			grid[i] = append(grid[i], spot{val: "."})
		}
	}

	for n, row := range strings.Split(input, "\n") {
		// if n == 1024 {
		// 	break
		// }
		if row == "" {
			continue
		}
		fields := strings.Split(row, ",")
		i, _ := strconv.Atoi(fields[1])
		j, _ := strconv.Atoi(fields[0])
		grid[i][j] = spot{
			val: "#",
		}
		if n > 1024 {
			fmt.Println(n)
			clearGrid(grid)
			visit(L-1, L-1, grid, 0)
			if grid[0][0].steps == 0 {
				fmt.Println(j, i)
				break
			}
		}
	}

	// printGrid(grid)
	// visit(L-1, L-1, grid, 0)
	// printGrid(grid)
	// fmt.Println(grid[0][0].steps)
}
