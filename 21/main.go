package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("test_data")
	// data, _ = os.ReadFile("data")
	input := string(data)
	codes := strings.Fields(input)

	codes = []string{"379A"}
	// fmt.Println(codes)

	answer := 0
	for _, code := range codes {
		num, _ := strconv.Atoi(code[:3])

		count := typeCode("A" + code)
		fmt.Println(count, num)
		answer += count * num
	}
	fmt.Println(answer)
}

func typeCode(code string) int {
	count := 0
	fmt.Println("K1: ", code)

	for i := 0; i < len(code)-1; i++ {
		start := string(code[i])
		end := string(code[i+1])
		directionPresses := getKeypadDirectionPresses(start, end)
		fmt.Println(start, end, "D1: ", directionPresses)

		for j := 0; j < len(directionPresses)-1; j++ {
			start := string(directionPresses[j])
			end := string(directionPresses[j+1])
			directionPresses := getDirectionPresses(start, end)
			fmt.Println("-", start, end, "D2: ", directionPresses)

			for k := 0; k < len(directionPresses)-1; k++ {
				start := string(directionPresses[k])
				end := string(directionPresses[k+1])
				directionPresses := getDirectionPresses(start, end)
				fmt.Println("--", start, end, "D3: ", directionPresses, len(directionPresses)-1)
				count += len(directionPresses) - 1
			}
			// fmt.Println()
		}
		fmt.Println()
	}
	return count
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getDirectionPresses(start, end string) string {
	si, sj := getDirectionLocation(start)
	ei, ej := getDirectionLocation(end)

	di := ei - si
	dj := ej - sj

	var presses strings.Builder
	var horizontal, vertical string

	presses.WriteString("A")

	if dj > 0 {
		horizontal = ">"
	} else {
		horizontal = "<"
	}
	if di > 0 {
		vertical = "v"
	} else {
		vertical = "^"
	}

	if dj > 0 {
		for j := 0; j < abs(dj); j++ {
			presses.WriteString(horizontal)
		}
		for i := 0; i < abs(di); i++ {
			presses.WriteString(vertical)
		}
	} else {
		for i := 0; i < abs(di); i++ {
			presses.WriteString(vertical)
		}
		for j := 0; j < abs(dj); j++ {
			presses.WriteString(horizontal)
		}
	}
	presses.WriteString("A")
	return presses.String()
}

func getKeypadDirectionPresses(start, end string) string {
	si, sj := getKeypadLocation(start)
	ei, ej := getKeypadLocation(end)

	di := ei - si
	dj := ej - sj

	var presses strings.Builder
	var horizontal, vertical string

	presses.WriteString("A")

	if dj > 0 {
		horizontal = ">"
	} else {
		horizontal = "<"
	}
	if di > 0 {
		vertical = "v"
	} else {
		vertical = "^"
	}

	if dj > 0 {
		for j := 0; j < abs(dj); j++ {
			presses.WriteString(horizontal)
		}
		for i := 0; i < abs(di); i++ {
			presses.WriteString(vertical)
		}
	} else {
		for i := 0; i < abs(di); i++ {
			presses.WriteString(vertical)
		}
		for j := 0; j < abs(dj); j++ {
			presses.WriteString(horizontal)
		}
	}
	presses.WriteString("A")
	return presses.String()
}

func getKeypadLocation(code string) (int, int) {
	i, j := 0, 0
	switch code {
	case "0":
		i, j = 3, 1
	case "1":
		i, j = 2, 0
	case "2":
		i, j = 2, 1
	case "3":
		i, j = 2, 2
	case "4":
		i, j = 1, 0
	case "5":
		i, j = 1, 1
	case "6":
		i, j = 1, 2
	case "7":
		i, j = 0, 0
	case "8":
		i, j = 0, 1
	case "9":
		i, j = 0, 2
	case "A":
		i, j = 3, 2
	}
	return i, j
}

func getDirectionLocation(code string) (int, int) {
	i, j := 0, 0
	switch code {
	case "<":
		i, j = 1, 0
	case "v":
		i, j = 1, 1
	case ">":
		i, j = 1, 2
	case "^":
		i, j = 0, 1
	case "A":
		i, j = 0, 2
	}
	return i, j
}
