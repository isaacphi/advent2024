package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid [][]string
type Loc [2]int

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func getPairs(locations []Loc) [][2]Loc {
	// Return permutations of antenna pairs
	pairs := make([][2]Loc, 0)
	for i, locA := range locations {
		for _, locB := range locations[i+1:] {
			pairs = append(pairs, [2]Loc{locA, locB})
		}
	}
	return pairs
}

func getAntinodes(pair [2]Loc) [2]Loc {
	a, b := pair[0], pair[1]
	distance := [2]int{b[0] - a[0], b[1] - a[1]}
	antinodeA := Loc{a[0] - distance[0], a[1] - distance[1]}
	antinodeB := Loc{b[0] + distance[0], b[1] + distance[1]}
	return [2]Loc{antinodeA, antinodeB}
}

func getCount(grid Grid) int {
	count := 0
	for _, row := range grid {
		for _, val := range row {
			if val == "#" {
				count++
			}
		}
	}
	return count
}

func countAntinodes(grid Grid) int {
	H, L := len(grid), len(grid[0])
	antennas := make(map[string][]Loc)
	for i, row := range grid {
		for j, val := range row {
			if val != "." {
				antennas[val] = append(antennas[val], Loc{i, j})
			}
		}
	}
	for _, locations := range antennas {
		pairs := getPairs(locations)
		for _, pair := range pairs {
			antinodes := getAntinodes(pair)
			for _, antinode := range antinodes {
				if antinode[0] >= 0 && antinode[0] < H && antinode[1] >= 0 && antinode[1] < L {
					grid[antinode[0]][antinode[1]] = "#"
				}
			}
		}
	}
	return getCount(grid)
}

func getAntinodesPt2(pair [2]Loc, H, L int) []Loc {
	a, b := pair[0], pair[1]
	antinodes := []Loc{a, b}
	distance := [2]int{b[0] - a[0], b[1] - a[1]}

	i, j := a[0], a[1]
	for i >= 0 && i < H && j >= 0 && j < L {
		antinode := Loc{i - distance[0], j - distance[1]}
		i, j = antinode[0], antinode[1]
		antinodes = append(antinodes, antinode)
	}
	i, j = b[0], b[1]
	for i >= 0 && i < H && j >= 0 && j < L {
		antinode := Loc{i + distance[0], j + distance[1]}
		i, j = antinode[0], antinode[1]
		antinodes = append(antinodes, antinode)
	}

	return antinodes
}

func countAntinodesPt2(grid Grid) int {
	H, L := len(grid), len(grid[0])
	antennas := make(map[string][]Loc)
	for i, row := range grid {
		for j, val := range row {
			if val != "." {
				antennas[val] = append(antennas[val], Loc{i, j})
			}
		}
	}
	for _, locations := range antennas {
		pairs := getPairs(locations)
		for _, pair := range pairs {
			antinodes := getAntinodesPt2(pair, H, L)
			for _, antinode := range antinodes {
				if antinode[0] >= 0 && antinode[0] < H && antinode[1] >= 0 && antinode[1] < L {
					grid[antinode[0]][antinode[1]] = "#"
				}
			}
		}
	}
	return getCount(grid)
}

func main() {
	content, _ := os.ReadFile("test_data")
	content, _ = os.ReadFile("data")

	lines := strings.Fields(string(content))
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	// fmt.Println(countAntinodes(grid))
	fmt.Println(countAntinodesPt2(grid))
}
