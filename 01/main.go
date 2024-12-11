package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Open the CSV file
	file, err := os.Open("1a.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var a []int
	var b []int
	// Print the CSV data
	for _, row := range data {
		for i, col := range row {
			if i == 0 {
				num, err := strconv.Atoi(col)
				a = append(a, num)
				if err != nil {
					panic(err)
				}
			}
			if i == 1 {
				num, err := strconv.Atoi(col)
				b = append(b, num)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	sort.Ints(a)
	sort.Ints(b)

	sum := 0
	for i := 0; i < len(a); i++ {
		sum += int(math.Abs(float64(a[i]) - float64(b[i])))
	}
	fmt.Println(sum)

	bFrequencies := map[int]int{}
	for _, val := range b {
		count, exists := bFrequencies[val]
		if !exists {
			bFrequencies[val] = 1
		} else {
			bFrequencies[val] = count + 1
		}
	}

	similarity := 0
	for _, val := range a {
		bOccurences, exists := bFrequencies[val]
		if !exists {
			bOccurences = 0
		}
		similarity += bOccurences * val
	}
	fmt.Println(similarity)
}
