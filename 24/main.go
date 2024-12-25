package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("test_data_2")
	data, _ = os.ReadFile("test_data_3")
	// data, _ = os.ReadFile("data")
	input := string(data)

	wires := make(map[string]*Wire)
	gates := make([]Gate, 0)

	for _, row := range strings.Split(input, "\n") {
		fields := strings.Fields(row)
		if len(fields) == 2 {
			// initial wire values
			wireName := fields[0][0:3]
			value, _ := strconv.Atoi(fields[1])
			wires[wireName] = &Wire{
				name:  wireName,
				value: &value,
			}
		}
		if len(fields) == 5 {
			// gates
			operation := fields[1]
			aWireName := fields[0]
			bWireName := fields[2]
			outWireName := fields[4]

			a, exists := wires[aWireName]
			if !exists {
				a = &Wire{name: aWireName}
				wires[aWireName] = a
			}
			b, exists := wires[bWireName]
			if !exists {
				b = &Wire{name: bWireName}
				wires[bWireName] = b
			}
			out, exists := wires[outWireName]
			if !exists {
				out = &Wire{name: outWireName}
				wires[outWireName] = out
			}

			gate := Gate{operation, a, b, out}
			gates = append(gates, gate)
		}
	}
	// part 1
	// for !outputReady(wires) {
	// 	tickAll(gates)
	// }
	// fmt.Println(getValue(wires))

	// part 2
	// swap 4 pairs of output gates
}

func outputReady(wires map[string]*Wire) bool {
	for _, w := range wires {
		if w.name[0] == 'z' {
			if w.value == nil {
				return false
			}
		}
	}
	return true
}

func tickAll(gates []Gate) {
	for i := range gates {
		gates[i].tick()
	}
}

func getValue(wires map[string]*Wire) int {
	num := 0
	for i := range wires {
		if wires[i].name[0] == 'z' {
			digit, _ := strconv.Atoi(wires[i].name[1:])
			num = num | (*wires[i].value << digit)
		}
	}
	return num
}

type Wire struct {
	name  string
	value *int
}

func (w Wire) print() {
	fmt.Print(w.name, " ")
	if w.value != nil {
		fmt.Print(*w.value)
	} else {
		fmt.Print(w.value)
	}
	fmt.Println()
}

type Gate struct {
	operation string
	a, b, out *Wire
}

func (g Gate) print() {
	fmt.Print(g.operation)
	g.a.print()
	g.b.print()
	g.out.print()
}

func (g *Gate) tick() {
	if g.a.value == nil || g.b.value == nil {
		return
	}
	if g.out.value == nil {
		g.out.value = new(int)
	}
	output := 0

	switch g.operation {
	case "AND":
		output = *g.a.value & *g.b.value
	case "OR":
		output = *g.a.value | *g.b.value
	case "XOR":
		output = *g.a.value ^ *g.b.value
	}

	*g.out.value = output
}
