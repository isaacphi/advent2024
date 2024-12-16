package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	i, j           int
	N, S, E, W     int
	isStart, isEnd bool
	val            string
}

type Map [][]*Point
type vec [2]int

func printGrid(grid Map, path map[vec]bool) {
	for i, row := range grid {
		for j, p := range row {
			if p != nil {
				if _, exists := path[vec{i, j}]; exists {
					fmt.Print("O")
				} else {
					fmt.Print(p.val)
				}
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

var ROTATION = 1000
var MOVEMENT = 1

func explore(m Map, i int, j int, dir string, newScore int) {
	p := m[i][j]
	var score int

	switch dir {
	case "N":
		score = p.N
	case "S":
		score = p.S
	case "E":
		score = p.E
	case "W":
		score = p.W
	}

	if score < 0 || newScore <= score {
		switch dir {
		case "N":
			p.N = newScore
			explore(m, i, j, "E", newScore+ROTATION)
			explore(m, i, j, "W", newScore+ROTATION)
			i++
		case "S":
			p.S = newScore
			explore(m, i, j, "E", newScore+ROTATION)
			explore(m, i, j, "W", newScore+ROTATION)
			i--
		case "E":
			p.E = newScore
			explore(m, i, j, "N", newScore+ROTATION)
			explore(m, i, j, "S", newScore+ROTATION)
			j--
		case "W":
			p.W = newScore
			explore(m, i, j, "N", newScore+ROTATION)
			explore(m, i, j, "S", newScore+ROTATION)
			j++
		}

		if m[i][j] != nil {
			explore(m, i, j, dir, newScore+MOVEMENT)
		}
	}
}

func findPath(m Map, path map[vec]bool, p *Point, dir string) {
	var nextMin int
	path[vec{p.i, p.j}] = true

	if p.isEnd {
		return
	}

	switch dir {
	case "N":
		i := p.i - 1
		j := p.j
		nextMin = min(p.E, p.W)
		if m[i][j] != nil {
			nextMin = min(nextMin, m[i][j].N)
			if m[i][j].N == nextMin {
				findPath(m, path, m[i][j], dir)
			}
		}
		if p.E == nextMin {
			findPath(m, path, p, "E")
		}
		if p.W == nextMin {
			findPath(m, path, p, "W")
		}

	case "S":
		i := p.i + 1
		j := p.j
		nextMin = min(p.E, p.W)
		if m[i][j] != nil {
			nextMin = min(nextMin, m[i][j].S)
			if m[i][j].S == nextMin {
				findPath(m, path, m[i][j], dir)
			}
		}
		if p.E == nextMin {
			findPath(m, path, p, "E")
		}
		if p.W == nextMin {
			findPath(m, path, p, "W")
		}

	case "E":
		i := p.i
		j := p.j + 1
		nextMin = min(p.N, p.S)
		if m[i][j] != nil {
			nextMin = min(nextMin, m[i][j].E)
			if m[i][j].E == nextMin {
				findPath(m, path, m[i][j], dir)
			}
		}
		if p.N == nextMin {
			findPath(m, path, p, "N")
		}
		if p.S == nextMin {
			findPath(m, path, p, "S")
		}

	case "W":
		i := p.i
		j := p.j - 1
		nextMin = min(p.N, p.S)
		if m[i][j] != nil {
			nextMin = min(nextMin, m[i][j].W)
			if m[i][j].W == nextMin {
				findPath(m, path, m[i][j], dir)
			}
		}
		if p.N == nextMin {
			findPath(m, path, p, "N")
		}
		if p.S == nextMin {
			findPath(m, path, p, "S")
		}
	}
}

func main() {
	data, _ := os.ReadFile("test_data")
	// data, _ = os.ReadFile("data")
	input := string(data)

	H := len(strings.Split(input, "\n"))
	var end *Point
	var start *Point
	grid := make([][]*Point, H)

	for i, row := range strings.Split(input, "\n") {
		grid[i] = make([]*Point, len(row))
		for j, val := range row {
			if val == '.' || val == 'E' || val == 'S' {
				grid[i][j] = &Point{
					i: i, j: j,
					N: -1, S: -1, E: -1, W: -1,
					val: string(val),
				}
			}
			if val == 'E' {
				grid[i][j].isEnd = true
				end = grid[i][j]
			}
			if val == 'S' {
				grid[i][j].isStart = true
				start = grid[i][j]
			}
		}
	}

	explore(grid, end.i, end.j, "E", 0)
	explore(grid, end.i, end.j, "W", 0)
	explore(grid, end.i, end.j, "N", 0)
	explore(grid, end.i, end.j, "S", 0)

	fmt.Println(start.E)

	path := make(map[vec]bool)
	findPath(grid, path, start, "E")
	fmt.Println(len(path))

	printGrid(grid, path)
	fmt.Println(grid[9][2])
	fmt.Println(grid[11][1])
}
