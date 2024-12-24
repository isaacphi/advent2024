package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("data")
	input := string(data)
	codes := strings.Fields(input)

	// codes = []string{"379A"}

	answer := 0
	for _, code := range codes {
		num, _ := strconv.Atoi(code[:len(code)-1])
		count := typeCode("A" + code)
		fmt.Println("code:", num, count)
		answer += count * num
	}
	fmt.Println(answer)
}

func typeCode(code string) int {
	count := 0

	for i := 0; i < len(code)-1; i++ {
		start := string(code[i])
		end := string(code[i+1])
		count += getKeypadPresses(start, end)
	}

	return count
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getPath(di, dj int, firstDirection string) string {
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

	if firstDirection == "horizontal" {
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

	out := presses.String()
	return out
}

func getDirectionPresses(start, end string, n int) int {
	si, sj := getDirectionLocation(start)
	ei, ej := getDirectionLocation(end)

	paths := make([]string, 0)
	if !(si == 0 && ej == 0) {
		paths = append(paths, getPath(ei-si, ej-sj, "horizontal"))
	}
	if !(sj == 0 && ei == 0) {
		paths = append(paths, getPath(ei-si, ej-sj, "vertical"))
	}
	if len(paths) == 2 && paths[0] == paths[1] {
		paths = []string{paths[0]}
	}

	minPresses := 99999999999

	for _, path := range paths {
		presses := 0
		for j := 0; j < len(path)-1; j++ {
			s := string(path[j])
			e := string(path[j+1])
			if n == 0 {
				// last keypad
				presses += 1
			} else {
				presses += getDirectionPresses(s, e, n-1)
			}
		}
		minPresses = min(minPresses, presses)
	}
	if n >= 0 {
		// fmt.Println(n, "direction pad from", start, "to", end, paths, minPresses)
	}

	return minPresses
}

func getKeypadPresses(start, end string) int {
	si, sj := getKeypadLocation(start)
	ei, ej := getKeypadLocation(end)

	paths := make([]string, 0)
	if !(si == 3 && ej == 0) {
		paths = append(paths, getPath(ei-si, ej-sj, "horizontal"))
	}
	if !(sj == 0 && ei == 3) {
		paths = append(paths, getPath(ei-si, ej-sj, "vertical"))
	}
	if len(paths) == 2 && paths[0] == paths[1] {
		paths = []string{paths[0]}
	}

	minPresses := 99999999999

	for _, path := range paths {
		presses := 0
		for j := 0; j < len(path)-1; j++ {
			s := string(path[j])
			e := string(path[j+1])
			presses += getDirectionPresses(s, e, 1)
		}
		minPresses = min(minPresses, presses)
	}
	fmt.Println("keypad from", start, "to", end, paths, minPresses)

	return minPresses
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
