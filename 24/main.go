package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("test_data")
	data, _ = os.ReadFile("test_data_2")
	data, _ = os.ReadFile("test_data_3")
	data, _ = os.ReadFile("data")
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
	// fmt.Println(getValue(wires, 'z'))

	// part 2
	// swap 4 pairs of output gates
	run(wires, gates)
	c := wires["wjg"] // carry bit for i = 2
	swaps := make([]string, 0)
	// var possibleErrors []*Gate
	for i := 2; i < 45; i++ {
		// there is no carry bit for i = 0, 1. Those subcircuits are correct
		// i is the current bit

		testResult := testDigit(wires, gates, i, c)
		fmt.Println("i =", i, testResult)

		if testResult == false {
			_, possibleErrors := getNextC(wires, gates, i, c)
			peGates := make([]string, 0)
			for _, pe := range possibleErrors {
				fmt.Println(pe.out.name)
				peGates = append(peGates, pe.out.name)
			}
			fmt.Println("possibleErrors", possibleErrors)

			for _, badWire := range peGates {
				foundSwap := false
				for j := range gates {
					if gates[j].out.name == badWire {
						for k := range gates {
							fmt.Println("trying swap", badWire, gates[k].out.name)
							gates[j].out, gates[k].out = gates[k].out, gates[j].out
							// todo: only check current z
							if !testDigit(wires, gates, i, c) || gates[j].out.name == "z05" {
								fmt.Println("failed")
								gates[j].out, gates[k].out = gates[k].out, gates[j].out
							} else {
								foundSwap = true
								swaps = append(swaps, gates[k].out.name, gates[j].out.name)
								fmt.Println("worked", gates[j].out.name)
								break
							}
						}
						if foundSwap {
							break
						}
					}
				}
				if foundSwap {
					break
				}
			}
		}

		c, _ = getNextC(wires, gates, i, c)
	}
	sort.Strings(swaps)
	for i, s := range swaps {
		fmt.Print(s)
		if i < len(swaps)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}

// test: x, y, c, z expected values for subcircuit i
type t [4]int

func testDigit(wires map[string]*Wire, gates []Gate, digit int, c *Wire) bool {
	digitString := strconv.Itoa(digit)
	if digit < 10 {
		digitString = "0" + digitString
	}
	x := wires["x"+digitString]
	y := wires["y"+digitString]
	// z := wires["z"+digitString]

	setValue(wires, 'x', 0)
	setValue(wires, 'y', 0)

	for _, v := range []t{
		t{0, 0, 0, 0},
		t{0, 0, 1, 0},
		t{0, 1, 0, 1},
		t{0, 1, 1, 1},
		t{1, 0, 0, 1},
		t{1, 0, 1, 1},
		t{1, 1, 0, 2},
		t{1, 1, 1, 2},
	} {
		setValue(wires, 'x', 0)
		setValue(wires, 'y', 0)
		clearOutput(wires)
		*x.value = v[0]
		*y.value = v[1]
		c.value = new(int)
		// *c.value = v[2]
		run(wires, gates)
		if digit == 36 {
			fmt.Println(v)
			fmt.Println(getValue(wires, 'z') >> digit)
		}
		isCorrectValue := (getValue(wires, 'z') >> digit) == v[3]
		if !isCorrectValue {
			return false
		}
	}
	return true
}

func getNextC(wires map[string]*Wire, gates []Gate, digit int, c *Wire) (nextC *Wire, possibleErrors []*Gate) {
	digitString := strconv.Itoa(digit)
	if digit < 10 {
		digitString = "0" + digitString
	}
	x := wires["x"+digitString]
	y := wires["y"+digitString]
	z := wires["z"+digitString]

	and1, and1Err := getGateFromInput(x, y, gates, "AND")
	if and1Err != nil {
		fmt.Println("can't find and1", and1)
		panic("can't find add1")
		return nil, []*Gate{}
	}

	xor1, xor1Err := getGateFromInput(x, y, gates, "XOR")
	if xor1Err != nil {
		fmt.Println("can't find xor1", xor1)
		panic("can't find xor1")
		return nil, []*Gate{}
	}

	and2, and2Err := getGateFromInput(c, xor1.out, gates, "AND")
	if and2Err != nil {
		fmt.Println("can't find and2", and2, c.name, xor1.out.name)
		lastCGate, _ := getGatesFromOutput(c, gates, "OR")
		return nil, []*Gate{xor1, lastCGate}
	}

	xor2, xor2Err := getGateFromInput(c, xor1.out, gates, "XOR")
	// xor2, xor2Err = getGatesFromOutput(z, gates, "XOR")
	if xor2Err != nil {
		fmt.Println("can't find xor2", xor2)
		lastCGate, _ := getGatesFromOutput(c, gates, "OR")
		// xor2, xor2Err = getGateFromInput(c, xor1.out, gates, "XOR")
		return nil, []*Gate{xor1, lastCGate}
	}

	or, orErr := getGateFromInput(and1.out, and2.out, gates, "OR")
	if orErr != nil {
		fmt.Println("can't find or", or)
		return nil, []*Gate{and2, and1}
	}

	zGate, zGateErr := getGatesFromOutput(z, gates, "XOR")
	if zGateErr != nil {
		fmt.Println("can't find output", zGate)
		return nil, []*Gate{xor2}
	}

	if orErr == nil {
		fmt.Println("ok")
		return or.out, []*Gate{}
	}
	panic("reached end but didn't return")
	return nil, []*Gate{}
}

func getGateFromInput(a, b *Wire, gates []Gate, operation string) (*Gate, error) {
	var gate *Gate
	count := 0
	for _, g := range gates {
		if g.operation == operation && (g.a == a && g.b == b || g.a == b && g.b == a) {
			gate = &g
			count++
		}
	}
	if count == 1 {
		return gate, nil
	}
	return nil, errors.New("Gate error")
}

func getGatesFromOutput(out *Wire, gates []Gate, operation string) (*Gate, error) {
	for _, g := range gates {
		if g.out == out && g.operation == operation {
			return &g, nil
		}
	}
	return nil, errors.New("Gate error")
}

func clearOutput(wires map[string]*Wire) {
	for wireName := range wires {
		if wires[wireName].name[0] != 'x' && wires[wireName].name[0] != 'y' {
			wires[wireName].value = nil
		}
	}
}

func setValue(wires map[string]*Wire, c byte, v int) {
	for wireName := range wires {
		if wires[wireName].name[0] == c {
			wires[wireName].value = new(int)
			*wires[wireName].value = 0
		}
	}
	for i := 0; i < 45; i++ {
		for wireName := range wires {
			if wires[wireName].name[0] == c {
				digit, _ := strconv.Atoi(wires[wireName].name[1:])
				if digit == i {
					*wires[wireName].value = *wires[wireName].value | ((v >> i) & 1)
				}
			}
		}
	}
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

func run(wires map[string]*Wire, gates []Gate) {
	i := 0
	for !outputReady(wires) && i < 1000 {
		tickAll(gates)
		i++
	}
	if i == 1000 {
		setValue(wires, 'z', 0)
	}
}

func tickAll(gates []Gate) {
	for i := range gates {
		gates[i].tick()
	}
}

func getValue(wires map[string]*Wire, c byte) int {
	num := 0
	for i := range wires {
		if wires[i].name[0] == c {
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
