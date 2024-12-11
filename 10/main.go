package main

import (
	"fmt"
	"os"
	"strings"
)

type loc [2]int
type summits map[loc]bool
type summitsList []loc

func (s summits) addSummits(newS summits) {
	for k, v := range newS {
		s[k] = v
	}
}

func initSummits(grid [][]summits) {
	for i, row := range grid {
		for j := range row {
			grid[i][j] = make(summits)
		}
	}
}

func (s *summitsList) addSummitsList(newS summitsList) {
	*s = append(*s, newS...)
}

func initSummitsList(grid [][]summitsList) {
	for i, row := range grid {
		for j := range row {
			grid[i][j] = make(summitsList, 0)
		}
	}
}
func printGrid(grid []string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func countTrailheadsPt2(grid []string) int {
	H, L := len(grid), len(grid[0])
	levels := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	accessibleSummits := make([][]summitsList, H)
	accessibleSummitsNext := make([][]summitsList, H)
	for i := range accessibleSummits {
		accessibleSummits[i] = make([]summitsList, L)
		accessibleSummitsNext[i] = make([]summitsList, L)
	}
	initSummitsList(accessibleSummits)
	initSummitsList(accessibleSummitsNext)

	for _, level := range levels {
		for i, row := range grid {
			for j, val := range row {
				curVal := int(val - '0')
				if level == 9 && curVal == level {
					accessibleSummits[i][j] = append(accessibleSummits[i][j], loc{i, j})
				}
				if curVal == level {
					curSummits := accessibleSummits[i][j]
					if i > 0 && int(grid[i-1][j]-'0') == level-1 {
						accessibleSummitsNext[i-1][j].addSummitsList(curSummits)
					}
					if i < H-1 && int(grid[i+1][j]-'0') == level-1 {
						accessibleSummitsNext[i+1][j].addSummitsList(curSummits)
					}
					if j > 0 && int(grid[i][j-1]-'0') == level-1 {
						accessibleSummitsNext[i][j-1].addSummitsList(curSummits)
					}
					if j < L-1 && int(grid[i][j+1]-'0') == level-1 {
						accessibleSummitsNext[i][j+1].addSummitsList(curSummits)
					}
				}
			}
		}

		accessibleSummits, accessibleSummitsNext = accessibleSummitsNext, accessibleSummits
		initSummitsList(accessibleSummitsNext)
	}

	count := 0
	for _, row := range accessibleSummits {
		for _, val := range row {
			count += len(val)
		}
	}

	return count
}

func countTrailheads(grid []string) int {
	H, L := len(grid), len(grid[0])
	levels := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	accessibleSummits := make([][]summits, H)
	accessibleSummitsNext := make([][]summits, H)
	for i := range accessibleSummits {
		accessibleSummits[i] = make([]summits, L)
		accessibleSummitsNext[i] = make([]summits, L)
	}
	initSummits(accessibleSummits)
	initSummits(accessibleSummitsNext)

	for _, level := range levels {
		for i, row := range grid {
			for j, val := range row {
				curVal := int(val - '0')
				if level == 9 && curVal == level {
					accessibleSummits[i][j][loc{i, j}] = true
				}
				if curVal == level {
					curSummits := accessibleSummits[i][j]
					if i > 0 && int(grid[i-1][j]-'0') == level-1 {
						accessibleSummitsNext[i-1][j].addSummits(curSummits)
					}
					if i < H-1 && int(grid[i+1][j]-'0') == level-1 {
						accessibleSummitsNext[i+1][j].addSummits(curSummits)
					}
					if j > 0 && int(grid[i][j-1]-'0') == level-1 {
						accessibleSummitsNext[i][j-1].addSummits(curSummits)
					}
					if j < L-1 && int(grid[i][j+1]-'0') == level-1 {
						accessibleSummitsNext[i][j+1].addSummits(curSummits)
					}
				}
			}
		}

		accessibleSummits, accessibleSummitsNext = accessibleSummitsNext, accessibleSummits
		initSummits(accessibleSummitsNext)
	}

	count := 0
	for _, row := range accessibleSummits {
		for _, val := range row {
			count += len(val)
		}
	}

	return count
}

func main() {
	data, _ := os.ReadFile("data")
	// data, _ := os.ReadFile("test_data")
	input := string(data)

	grid := strings.Fields(input)
	printGrid(grid)

	// count := countTrailheads(grid)
	count := countTrailheadsPt2(grid)
	fmt.Println()
	fmt.Println(count)
}
