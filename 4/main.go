package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func printWordSearch(wordSearch []string) {
	for _, line := range wordSearch {
		fmt.Println(line)
	}
	fmt.Println()
}

func convertToVertical(wordSearch []string) []string {
	H := len(wordSearch)
	L := len(wordSearch[0])

	var verticalSearch []string
	for i := 0; i < L; i++ {
		var verticalString strings.Builder
		verticalString.Grow(H)
		for j := 0; j < H; j++ {
			verticalString.WriteByte(wordSearch[j][i])
		}
		verticalSearch = append(verticalSearch, verticalString.String())
	}
	return verticalSearch
}

func findMatches(s string) {
	matches := 0
	// Go regexp doesn't support positive lookahead
	patterns := []string{`XMAS`, `SAMX`}
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches += len(re.FindAllString(s, -1))
	}
	fmt.Println(matches)
}

func part1(wordSearch []string) {
	// printWordSearch(wordSearch)
	// What might be interesting is using regexp again.
	// Horizontal matches are easy. Just search for XMAS and SAMX
	H := len(wordSearch)
	L := len(wordSearch[0])
	verticalSearch := convertToVertical(wordSearch)
	// printWordSearch(verticalSearch)
	// then search that for XMAS and SAMX

	// for diagonal, it's trickier. Build an offset horizontal offset list then convert to vertical
	var diag1 []string
	for i := 0; i < H; i++ {
		var s strings.Builder
		for range i {
			s.WriteRune(' ')
		}
		s.WriteString(wordSearch[i])
		for range L - i {
			s.WriteRune(' ')
		}
		diag1 = append(diag1, s.String())
	}
	flippedDiag1 := convertToVertical(diag1)
	// printWordSearch(flippedDiag1)

	// and the other diagonal
	var diag2 []string
	for i := 0; i < H; i++ {
		var s strings.Builder
		for range L - i {
			s.WriteRune(' ')
		}
		s.WriteString(wordSearch[i])
		for range i {
			s.WriteRune(' ')
		}
		diag2 = append(diag2, s.String())
	}
	flippedDiag2 := convertToVertical(diag2)
	// printWordSearch(flippedDiag2)

	// Now combine them all into one long string
	var sb strings.Builder
	for _, s := range wordSearch {
		sb.WriteString(s)
		sb.WriteRune(' ')
	}
	for _, s := range verticalSearch {
		sb.WriteString(s)
		sb.WriteRune(' ')
	}
	for _, s := range flippedDiag1 {
		sb.WriteString(s)
		sb.WriteRune(' ')
	}
	for _, s := range flippedDiag2 {
		sb.WriteString(s)
		sb.WriteRune(' ')
	}

	finalString := sb.String()
	findMatches(finalString)
}

func isXMAS(grid [][]string, row, col int) int {
	//  M S
	//   A
	//  M S
	// the rule is that M and S must be on opposite corners
	tl := grid[row-1][col-1]
	tr := grid[row-1][col+1]
	bl := grid[row+1][col-1]
	br := grid[row+1][col+1]
	if (tl == "M" && br == "S" || tl == "S" && br == "M") &&
		(bl == "M" && tr == "S" || bl == "S" && tr == "M") {
		fmt.Println(row, col)
		return 1
	}
	return 0
}

func part2(grid [][]string) {
	fmt.Println(grid)
	L := len(grid[0])
	H := len(grid)
	// traverse the grid looking for A's
	// exclude the outer rows and columns
	// check to see if each A is the center of a MAS
	count := 0
	for i := 1; i < L-1; i++ {
		for j := 1; j < H-1; j++ {
			if grid[i][j] == "A" {
				count += isXMAS(grid, i, j)
			}
		}
	}
	fmt.Println(count)
}

func main() {
	content, _ := os.ReadFile("data")
	wordSearch := strings.Fields(string(content))
	var grid [][]string

	for _, line := range wordSearch {
		grid = append(grid, strings.Split(line, ""))
	}

	// part1(wordSearch)
	part2(grid)
}
