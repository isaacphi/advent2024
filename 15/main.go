package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type vec struct {
	x, y int
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
}

func move(grid [][]string, dir string, start vec) vec {
	if grid[start.x][start.y] != "@" {
		panic("wrong start")
	}
	switch dir {
	case "^":
		i := start.x - 1
		j := start.y
		for {
			next := grid[i][j]
			if next == "#" {
				return start
			} else if next == "O" {
				i--
			} else if next == "." {
				for grid[i][j] != "@" {
					grid[i][j] = grid[i+1][j]
					i++
				}
				grid[i][j] = "."
				return vec{x: i - 1, y: j}
			}
		}
	case "v":
		i := start.x + 1
		j := start.y
		for {
			next := grid[i][j]
			if next == "#" {
				return start
			} else if next == "O" {
				i++
			} else if next == "." {
				for grid[i][j] != "@" {
					grid[i][j] = grid[i-1][j]
					i--
				}
				grid[i][j] = "."
				return vec{x: i + 1, y: j}
			}
		}
	case "<":
		i := start.x
		j := start.y - 1
		for {
			next := grid[i][j]
			if next == "#" {
				return start
			} else if next == "O" {
				j--
			} else if next == "." {
				for grid[i][j] != "@" {
					grid[i][j] = grid[i][j+1]
					j++
				}
				grid[i][j] = "."
				return vec{x: i, y: j - 1}
			}
		}
	case ">":
		i := start.x
		j := start.y + 1
		for {
			next := grid[i][j]
			if next == "#" {
				return start
			} else if next == "O" {
				j++
			} else if next == "." {
				for grid[i][j] != "@" {
					grid[i][j] = grid[i][j-1]
					j--
				}
				grid[i][j] = "."
				return vec{x: i, y: j + 1}
			}
		}
	default:
		panic("Invalid direction")
	}
}

func countScore(grid [][]string) int {
	count := 0
	for i, row := range grid {
		for j, val := range row {
			if val == "O" {
				count += 100*i + j
			}
		}
	}
	return count
}

func countScore2(grid [][]string) int {
	count := 0
	for i, row := range grid {
		for j, val := range row {
			if val == "[" {
				count += 100*i + j
			}
		}
	}
	return count
}

func pushUp(grid [][]string, blocks []vec) (newBlocks []vec, err error) {
	for _, b := range blocks {
		next := grid[b.x-1][b.y]
		if next == "#" {
			return make([]vec, 0), errors.New("no push")
		}
		if next == "." {
			newBlocks = append(newBlocks, b)
		}
		if next == "[" {
			bs := []vec{
				vec{x: b.x - 1, y: b.y},
				vec{x: b.x - 1, y: b.y + 1},
			}
			moreBlocks, err := pushUp(grid, bs)
			newBlocks = append(newBlocks, moreBlocks...)
			newBlocks = append(newBlocks, bs...)
			if err != nil {
				return make([]vec, 0), errors.New("no push")
			}
		}
		if next == "]" {
			bs := []vec{
				vec{x: b.x - 1, y: b.y},
				vec{x: b.x - 1, y: b.y - 1},
			}
			moreBlocks, err := pushUp(grid, bs)
			newBlocks = append(newBlocks, moreBlocks...)
			newBlocks = append(newBlocks, bs...)
			if err != nil {
				return make([]vec, 0), errors.New("no push")
			}
		}
	}
	return newBlocks, nil
}

func pushDown(grid [][]string, blocks []vec) (newBlocks []vec, err error) {
	for _, b := range blocks {
		next := grid[b.x+1][b.y]
		if next == "#" {
			return make([]vec, 0), errors.New("no push")
		}
		if next == "." {
			newBlocks = append(newBlocks, b)
		}
		if next == "[" {
			bs := []vec{
				vec{x: b.x + 1, y: b.y},
				vec{x: b.x + 1, y: b.y + 1},
			}
			moreBlocks, err := pushDown(grid, bs)
			newBlocks = append(newBlocks, moreBlocks...)
			newBlocks = append(newBlocks, bs...)
			if err != nil {
				return make([]vec, 0), errors.New("no push")
			}
		}
		if next == "]" {
			bs := []vec{
				vec{x: b.x + 1, y: b.y},
				vec{x: b.x + 1, y: b.y - 1},
			}
			moreBlocks, err := pushDown(grid, bs)
			newBlocks = append(newBlocks, moreBlocks...)
			newBlocks = append(newBlocks, bs...)
			if err != nil {
				return make([]vec, 0), errors.New("no push")
			}
		}
	}
	return newBlocks, nil
}

func move2(grid [][]string, start vec, dir string) (newGrid [][]string, newStart vec) {
	if grid[start.x][start.y] != "@" {
		panic("wrong start")
	}
	newGrid = make([][]string, len(grid))
	for i := range grid {
		newGrid[i] = make([]string, len(grid[0]))
		copy(newGrid[i], grid[i])
	}
	switch dir {
	case "^":
		i := start.x - 1
		j := start.y
		for {
			next := grid[i][j]
			if next == "#" {
				return grid, start
			} else if next == "." {
				newGrid[start.x][start.y] = "."
				newGrid[start.x-1][start.y] = "@"
				return newGrid, vec{x: start.x - 1, y: start.y}
			} else if next == "[" || next == "]" {
				blocks, err := pushUp(grid, []vec{start})
				if err != nil {
					return grid, start
				}
				for _, b := range blocks {
					newGrid[b.x-1][b.y] = grid[b.x][b.y]
					newGrid[b.x][b.y] = "."
					for _, otherB := range blocks {
						if b.y == otherB.y && b.x+1 == otherB.x {
							newGrid[b.x][b.y] = grid[otherB.x][otherB.y]
						}
					}
				}
				if next == "[" {
					newGrid[start.x-1][start.y+1] = "."
				} else {
					newGrid[start.x-1][start.y-1] = "."
				}
				newGrid[start.x][start.y] = "."
				newGrid[start.x-1][start.y] = "@"
				return newGrid, vec{x: start.x - 1, y: start.y}
			}
		}
	case "v":
		i := start.x + 1
		j := start.y
		for {
			next := grid[i][j]
			if next == "#" {
				return grid, start
			} else if next == "." {
				newGrid[start.x][start.y] = "."
				newGrid[start.x+1][start.y] = "@"
				return newGrid, vec{x: start.x + 1, y: start.y}
			} else if next == "[" || next == "]" {
				blocks, err := pushDown(grid, []vec{start})
				if err != nil {
					return grid, start
				}
				for _, b := range blocks {
					newGrid[b.x+1][b.y] = grid[b.x][b.y]
					newGrid[b.x][b.y] = "."
					for _, otherB := range blocks {
						if b.y == otherB.y && b.x-1 == otherB.x {
							newGrid[b.x][b.y] = grid[otherB.x][otherB.y]
						}
					}
				}
				if next == "[" {
					newGrid[start.x+1][start.y+1] = "."
				} else {
					newGrid[start.x+1][start.y-1] = "."
				}
				newGrid[start.x][start.y] = "."
				newGrid[start.x+1][start.y] = "@"
				return newGrid, vec{x: start.x + 1, y: start.y}
			}
		}
	case "<":
		i := start.x
		j := start.y - 1
		for {
			next := grid[i][j]
			if next == "#" {
				return grid, start
			} else if next == "[" || next == "]" {
				j--
			} else if next == "." {
				for grid[i][j] != "@" {
					grid[i][j] = grid[i][j+1]
					j++
				}
				grid[i][j] = "."
				return grid, vec{x: i, y: j - 1}
			}
		}
	case ">":
		i := start.x
		j := start.y + 1
		for {
			next := grid[i][j]
			if next == "#" {
				return grid, start
			} else if next == "[" || next == "]" {
				j++
			} else if next == "." {
				for grid[i][j] != "@" {
					grid[i][j] = grid[i][j-1]
					j--
				}
				grid[i][j] = "."
				return grid, vec{x: i, y: j + 1}
			}
		}
	default:
		panic("Invalid direction")
	}
}

func main() {
	// data, _ := os.ReadFile("test_data")
	data, _ := os.ReadFile("data")

	input := string(data)

	var start vec
	var start2 vec
	grid := make([][]string, 0)
	grid2 := make([][]string, 0)
	moves := make([]string, 0)
	isBuildingGrid := true

	for i, row := range strings.Split(input, "\n") {
		if row == "" {
			isBuildingGrid = false
		}
		if isBuildingGrid {
			grid = append(grid, make([]string, 0))
			grid2 = append(grid2, make([]string, 0))
			for j, val := range row {
				if val == '@' {
					start = vec{x: i, y: j}
					start2 = vec{x: i, y: j * 2}
				}
				grid[i] = append(grid[i], string(val))
				if val == '@' {
					grid2[i] = append(grid2[i], "@", ".")
				} else if val == '#' {
					grid2[i] = append(grid2[i], "#", "#")
				} else if val == 'O' {
					grid2[i] = append(grid2[i], "[", "]")
				} else if val == '.' {
					grid2[i] = append(grid2[i], ".", ".")
				}
			}
		} else {
			for _, val := range row {
				moves = append(moves, string(val))
			}
		}
	}
	// printGrid(grid)
	for _, m := range moves {
		start = move(grid, m, start)
	}
	printGrid(grid2)
	fmt.Println(start2)
	// fmt.Println(countScore(grid))
	for _, m := range moves {
		grid2, start2 = move2(grid2, start2, m)
	}
	printGrid(grid2)
	fmt.Println(countScore2(grid2))
}
