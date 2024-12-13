package main

import (
	"fmt"
	"os"
	"strings"
)

type square struct {
	visited bool
	top     bool
	right   bool
	bottom  bool
	left    bool
}

var visited2 [][]square

func visitArea(i, j int, grid []string, visited [][]bool) (area int, perimeter int) {
	H, L := len(grid), len(grid[0])
	areaType := grid[i][j]

	visited[i][j] = true

	area, perimeter = 1, 0

	if i > 0 {
		if grid[i-1][j] == areaType {
			if !visited[i-1][j] {
				a, p := visitArea(i-1, j, grid, visited)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}

	if i < H-1 {
		if grid[i+1][j] == areaType {
			if !visited[i+1][j] {
				a, p := visitArea(i+1, j, grid, visited)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}

	if j > 0 {
		if grid[i][j-1] == areaType {
			if !visited[i][j-1] {
				a, p := visitArea(i, j-1, grid, visited)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}

	if j < L-1 {
		if grid[i][j+1] == areaType {
			if !visited[i][j+1] {
				a, p := visitArea(i, j+1, grid, visited)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	} else {
		perimeter++
	}

	return area, perimeter
}

func visitLine(i, j int, areaType byte, side string, grid []string) {
	H, L := len(grid), len(grid[0])
	switch side {
	case "top":
		if visited2[i][j].top {
			return
		}
		visited2[i][j].top = true
		if i > 0 {
			if j > 0 && grid[i][j-1] == areaType && grid[i-1][j-1] != areaType {
				visitLine(i, j-1, areaType, "top", grid)
			}
			if j < L-1 && grid[i][j+1] == areaType && grid[i-1][j+1] != areaType {
				visitLine(i, j+1, areaType, "top", grid)
			}
		} else {
			if j > 0 && grid[i][j-1] == areaType {
				visitLine(i, j-1, areaType, "top", grid)
			}
			if j < L-1 && grid[i][j+1] == areaType {
				visitLine(i, j+1, areaType, "top", grid)
			}
		}

	case "bottom":
		if visited2[i][j].bottom {
			return
		}
		visited2[i][j].bottom = true
		if i < H-1 {
			if j > 0 && grid[i][j-1] == areaType && grid[i+1][j-1] != areaType {
				visitLine(i, j-1, areaType, "bottom", grid)
			}
			if j < L-1 && grid[i][j+1] == areaType && grid[i+1][j+1] != areaType {
				visitLine(i, j+1, areaType, "bottom", grid)
			}
		} else {
			if j > 0 && grid[i][j-1] == areaType {
				visitLine(i, j-1, areaType, "bottom", grid)
			}
			if j < L-1 && grid[i][j+1] == areaType {
				visitLine(i, j+1, areaType, "bottom", grid)
			}
		}

	case "left":
		if visited2[i][j].left {
			return
		}
		visited2[i][j].left = true
		if j > 0 {
			if i > 0 && grid[i-1][j] == areaType && grid[i-1][j-1] != areaType {
				visitLine(i-1, j, areaType, "left", grid)
			}
			if i < H-1 && grid[i+1][j] == areaType && grid[i+1][j-1] != areaType {
				visitLine(i+1, j, areaType, "left", grid)
			}
		} else {
			if i > 0 && grid[i-1][j] == areaType {
				visitLine(i-1, j, areaType, "left", grid)
			}
			if i < H-1 && grid[i+1][j] == areaType {
				visitLine(i+1, j, areaType, "left", grid)
			}
		}

	case "right":
		if visited2[i][j].right {
			return
		}
		visited2[i][j].right = true
		if j < L-1 {
			if i > 0 && grid[i-1][j] == areaType && grid[i-1][j+1] != areaType {
				visitLine(i-1, j, areaType, "right", grid)
			}
			if i < H-1 && grid[i+1][j] == areaType && grid[i+1][j+1] != areaType {
				visitLine(i+1, j, areaType, "right", grid)
			}
		} else {
			if i > 0 && grid[i-1][j] == areaType {
				visitLine(i-1, j, areaType, "right", grid)
			}
			if i < H-1 && grid[i+1][j] == areaType {
				visitLine(i+1, j, areaType, "right", grid)
			}
		}
	}
}

func visitArea2(i, j int, grid []string, visited [][]bool) (area int, perimeter int) {
	H, L := len(grid), len(grid[0])
	areaType := grid[i][j]

	visited[i][j] = true

	area, perimeter = 1, 0

	if i > 0 {
		if grid[i-1][j] == areaType {
			if !visited[i-1][j] {
				a, p := visitArea2(i-1, j, grid, visited)
				area += a
				perimeter += p
			}
		} else if !visited2[i][j].top {
			visitLine(i, j, areaType, "top", grid)
			perimeter++
		}
	} else if !visited2[i][j].top {
		visitLine(i, j, areaType, "top", grid)
		perimeter++
	}
	if i < H-1 {
		if grid[i+1][j] == areaType {
			if !visited[i+1][j] {
				a, p := visitArea2(i+1, j, grid, visited)
				area += a
				perimeter += p
			}
		} else if !visited2[i][j].bottom {
			visitLine(i, j, areaType, "bottom", grid)
			perimeter++
		}
	} else if !visited2[i][j].bottom {
		visitLine(i, j, areaType, "bottom", grid)
		perimeter++
	}
	if j > 0 {
		if grid[i][j-1] == areaType {
			if !visited[i][j-1] {
				a, p := visitArea2(i, j-1, grid, visited)
				area += a
				perimeter += p
			}
		} else if !visited2[i][j].left {
			visitLine(i, j, areaType, "left", grid)
			perimeter++
		}
	} else if !visited2[i][j].left {
		visitLine(i, j, areaType, "left", grid)
		perimeter++
	}
	if j < L-1 {
		if grid[i][j+1] == areaType {
			if !visited[i][j+1] {
				a, p := visitArea2(i, j+1, grid, visited)
				area += a
				perimeter += p
			}
		} else if !visited2[i][j].right {
			visitLine(i, j, areaType, "right", grid)
			perimeter++
		}
	} else if !visited2[i][j].right {
		visitLine(i, j, areaType, "right", grid)
		perimeter++
	}

	return area, perimeter
}

func searchAreas(grid []string) int {
	H, L := len(grid), len(grid[0])
	visited := make([][]bool, H)
	visited2 = make([][]square, H)
	for i := range visited {
		visited[i] = make([]bool, L)
		visited2[i] = make([]square, L)
	}

	areas, perimeters := make([]int, 0), make([]int, 0)
	for i, row := range grid {
		for j := range row {
			if !visited[i][j] {
				// a, p := visitArea(i, j, grid, visited)
				a, p := visitArea2(i, j, grid, visited)
				areas = append(areas, a)
				perimeters = append(perimeters, p)
			}
		}
	}

	total := 0
	for i := range areas {
		total += areas[i] * perimeters[i]
	}
	return total
}

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("data")

	input := string(data)
	grid := strings.Fields(input)

	answer := searchAreas(grid)
	fmt.Println(answer)
}
