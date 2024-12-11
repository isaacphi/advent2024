package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isSafe(row []string) bool {
	increasing := false
	decreasing := false
	safe := true

	for i, col := range row {
		if i == 0 {
			continue
		}
		val, _ := strconv.Atoi(col)
		lastVal, _ := strconv.Atoi(row[i-1])

		diff := float64(val - lastVal)
		if math.Abs(diff) > 3 || diff == 0 {
			safe = false
			break
		}
		if diff > 0 {
			increasing = true
		} else {
			decreasing = true
		}
		if increasing && decreasing {
			safe = false
			break
		}
	}
	return safe
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	data, _ := reader.ReadAll()

	numSafe := 0
	for _, row := range data {
		for i, _ := range row {
			fmt.Println(i)
			// SLICES IN GO ARE WEIRD
			newRow := make([]string, len(row)-1)
			copy(newRow, row[:i])
			copy(newRow[i:], row[i+1:])
			fmt.Println(row)
			fmt.Println(newRow)
			fmt.Println()
			if isSafe(newRow) {
				numSafe += 1
				break
			}
		}
	}

	fmt.Println(numSafe)
}

func part1() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	numSafe := 0
	for _, row := range data {
		increasing := false
		decreasing := false
		safe := true
		for i, col := range row {
			if i == 0 {
				continue
			}
			val, _ := strconv.Atoi(col)
			lastVal, _ := strconv.Atoi(row[i-1])
			diff := float64(val - lastVal)
			if math.Abs(diff) > 3 || diff == 0 {
				safe = false
				break
			}
			if diff > 0 {
				increasing = true
			} else {
				decreasing = true
			}
			if increasing && decreasing {
				safe = false
				break
			}
		}
		if safe {
			numSafe += 1
			fmt.Println(row)
		}
	}
	fmt.Println(numSafe)
}
