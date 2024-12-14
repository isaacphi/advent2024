package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vec struct {
	x int
	y int
}

type machine struct {
	a vec
	b vec
	p vec
}

func countTokens(m machine) int {
	aNum := m.p.y*m.b.x - m.p.x*m.b.y
	aDenom := m.a.y*m.b.x - m.a.x*m.b.y
	bNum := m.a.x*m.p.y - m.a.y*m.p.x
	bDenom := m.a.x*m.b.y - m.a.y*m.b.x

	if aNum%aDenom == 0 && bNum%bDenom == 0 {
		return 3*aNum/aDenom + 1*bNum/bDenom
	}

	return 0
}

func main() {
	// file, _ := os.ReadFile("test_data")
	file, _ := os.ReadFile("data")

	machines := make([]machine, 0)

	data := string(file)
	examples := strings.Split(data, "\n\n")
	for _, example := range examples {
		rows := strings.Split(example, "\n")
		machine := machine{}
		for i, row := range rows {
			fields := strings.Fields(row)
			if i == 0 {
				x, _ := strconv.Atoi(fields[2][2 : len(fields[2])-1])
				y, _ := strconv.Atoi(fields[3][2:])
				machine.a = vec{x, y}
			}
			if i == 1 {
				x, _ := strconv.Atoi(fields[2][2 : len(fields[2])-1])
				y, _ := strconv.Atoi(fields[3][2:])
				machine.b = vec{x, y}
			}
			if i == 2 {
				x, _ := strconv.Atoi(fields[1][2 : len(fields[1])-1])
				y, _ := strconv.Atoi(fields[2][2:])
				// machine.p = vec{x, y}
				machine.p = vec{x + 10000000000000, y + 10000000000000}
			}
		}
		machines = append(machines, machine)
	}

	count := 0
	for _, m := range machines {
		count += countTokens(m)
	}
	fmt.Println(count)
}
