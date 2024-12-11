package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func getStartLocation(grid [][]string) [2]int {
	for i, row := range grid {
		for j, col := range row {
			if col == "^" || col == ">" || col == "v" || col == "<" {
				return [2]int{i, j}
			}
		}
	}
	panic("Start location not found")
}

func printGrid(grid [][]string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()

	for _, row := range grid {
		fmt.Println(row)
	}
}

func update(grid [][]string) bool {
	// Return false if next step leaves the grid and true if can continue
	currentLocation := getStartLocation(grid)
	i, j := currentLocation[0], currentLocation[1]
	direction := grid[i][j]

	// Mark as visited
	grid[i][j] = "X"

	H, L := len(grid), len(grid[0])

	if direction == "^" {
		if i == 0 {
			return false
		}
		if grid[i-1][j] == "#" {
			grid[i][j] = ">"
		} else {
			grid[i-1][j] = "^"
		}
	} else if direction == ">" {
		if j == L-1 {
			return false
		}
		if grid[i][j+1] == "#" {
			grid[i][j] = "v"
		} else {
			grid[i][j+1] = ">"
		}
	} else if direction == "v" {
		if i == H-1 {
			return false
		}
		if grid[i+1][j] == "#" {
			grid[i][j] = "<"
		} else {
			grid[i+1][j] = "v"
		}
	} else if direction == "<" {
		if j == 0 {
			return false
		}
		if grid[i][j-1] == "#" {
			grid[i][j] = "^"
		} else {
			grid[i][j-1] = "<"
		}
	} else {
		panic("Unknown direction")
	}

	return true
}

func countVisits(grid [][]string) int {
	count := 0
	for _, row := range grid {
		for _, col := range row {
			if col == "X" {
				count++
			}
		}
	}
	return count
}

func part1(grid [][]string) int {
	for update(grid) {
		// time.Sleep(100 * time.Millisecond)
		// printGrid(grid)
	}
	return countVisits(grid)
}

// Part 2

func isLoop(obstacles map[string]map[int][]int, start [2]int, direction string) bool {
	visited := make([]string, 0)
	i, j := start[0], start[1]
	for {
		willExit := true
		switch direction {
		case "^":
			for _, val := range obstacles["^"][j] {
				if val < i {
					i = val + 1
					direction = ">"
					willExit = false
					break
				}
			}
		case ">":
			for _, val := range obstacles[">"][i] {
				if val > j {
					j = val - 1
					direction = "v"
					willExit = false
					break
				}
			}
		case "v":
			for _, val := range obstacles["v"][j] {
				if val > i {
					i = val - 1
					direction = "<"
					willExit = false
					break
				}
			}
		case "<":
			for _, val := range obstacles["<"][i] {
				if val < j {
					j = val + 1
					direction = "^"
					willExit = false
					break
				}
			}
		}

		if willExit {
			return false
		}
		visitedLocation := fmt.Sprintf("%d,%d%s", i, j, direction)
		for _, val := range visited {
			if val == visitedLocation {
				return true
			}
		}
		visited = append(visited, visitedLocation)
	}
}

func makeObstacles(grid [][]string) map[string]map[int][]int {
	// obstacles holds the position of each obstacle as seen by each direction for each row or column
	obstacles := map[string]map[int][]int{
		"^": make(map[int][]int),
		">": make(map[int][]int),
		"v": make(map[int][]int),
		"<": make(map[int][]int),
	}
	for i := range grid {
		obstacles[">"][i] = []int{}
		obstacles["<"][i] = []int{}
	}
	for j := range grid[0] {
		obstacles["^"][j] = []int{}
		obstacles["v"][j] = []int{}
	}
	for i, row := range grid {
		for j, col := range row {
			if col == "#" {
				obstacles["^"][j] = append(obstacles["^"][j], i)
				obstacles[">"][i] = append(obstacles[">"][i], j)
				obstacles["v"][j] = append(obstacles["v"][j], i)
				obstacles["<"][i] = append(obstacles["<"][i], j)
			}
		}
	}
	for i := range grid {
		slices.Reverse(obstacles["<"][i])
	}
	for j := range grid[0] {
		slices.Reverse(obstacles["^"][j])
	}

	return obstacles
}

func getObstacleLocations(grid [][]string) [][2]int {
	tempGrid := make([][]string, len(grid))
	for i := range grid {
		tempGrid[i] = make([]string, len(grid[i]))
		copy(tempGrid[i], grid[i])
	}
	part1(tempGrid)
	locations := [][2]int{}
	for i, row := range tempGrid {
		for j, col := range row {
			if col == "X" {
				locations = append(locations, [2]int{i, j})
			}
		}
	}
	return locations
}

func part2(grid [][]string) int {
	count := 0
	startLocation := getStartLocation(grid)
	direction := grid[startLocation[0]][startLocation[1]]
	for _, obstacleLocation := range getObstacleLocations(grid) {
		// Make a temporary grid with a hypothetical obstacle
		tempGrid := make([][]string, len(grid))
		for i := range grid {
			tempGrid[i] = make([]string, len(grid[i]))
			copy(tempGrid[i], grid[i])
		}
		tempGrid[obstacleLocation[0]][obstacleLocation[1]] = "#"

		obstacles := makeObstacles(tempGrid)

		if isLoop(obstacles, startLocation, direction) {
			count++
		}
	}

	return count
}

func main() {
	content, _ := os.ReadFile("test_data")
	content, _ = os.ReadFile("data")

	lines := strings.Fields(string(content))
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	// part1(grid)
	count := part2(grid)
	fmt.Println(count)
}
