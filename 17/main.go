package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type computer struct {
	A, B, C        int
	instructionPtr int
	instructions   []int
	output         []int
}

func (c *computer) getLiteralOperand() (int, error) {
	if c.instructionPtr == len(c.instructions)-1 {
		return 0, errors.New("invalid program: no operand")
	}
	return c.instructions[c.instructionPtr+1], nil
}

func (c *computer) getComboOperand() (int, error) {
	if c.instructionPtr == len(c.instructions)-1 {
		return 0, errors.New("invalid program: no operand")
	}
	op := c.instructions[c.instructionPtr+1]
	if op <= 3 {
		return op, nil
	} else if op == 4 {
		return c.A, nil
	} else if op == 5 {
		return c.B, nil
	} else if op == 6 {
		return c.C, nil
	}
	return 0, errors.New("invalid program: combo operand 7")
}

func printInstructions(instructions []int) string {
	var sb strings.Builder
	for i, n := range instructions {
		sb.WriteString(strconv.Itoa(n))
		if i != len(instructions)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func pow2(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 2
	}
	return result
}

func (c *computer) process() error {
	if c.instructionPtr >= len(c.instructions) {
		return errors.New("program halt")
	}

	switch c.instructions[c.instructionPtr] {
	case 0:
		op, err := c.getComboOperand()
		if err != nil {
			return err
		}
		c.A = c.A / pow2(op)
	case 1:
		op, _ := c.getLiteralOperand()
		c.B = c.B ^ op
	case 2:
		op, err := c.getComboOperand()
		if err != nil {
			return err
		}
		c.B = op % 8
	case 3:
		if c.A != 0 {
			op, _ := c.getLiteralOperand()
			c.instructionPtr = op
			return nil
		}
	case 4:
		_, _ = c.getLiteralOperand()
		c.B = c.B ^ c.C
	case 5:
		op, err := c.getComboOperand()
		if err != nil {
			return err
		}
		c.output = append(c.output, op%8)
		if len(c.output) > len(c.instructions) || c.output[len(c.output)-1] != c.instructions[len(c.output)-1] {
			return errors.New("wrong value for A")
		}
	case 6:
		op, err := c.getComboOperand()
		if err != nil {
			return err
		}
		c.B = c.A / pow2(op)
	case 7:
		op, err := c.getComboOperand()
		if err != nil {
			return err
		}
		c.C = c.A / pow2(op)
	}

	c.instructionPtr += 2
	return nil
}

func parseInstructions(instructions string) []int {
	output := make([]int, 0)
	for _, e := range strings.Split(instructions, ",") {
		inst, _ := strconv.Atoi(e)
		output = append(output, inst)
	}
	return output
}

func parseRegister(row string) int {
	reg, _ := strconv.Atoi(strings.Split(row, " ")[2])
	return reg
}

func main() {
	data, _ := os.ReadFile("test_data")
	// data, _ = os.ReadFile("data")

	input := string(data)
	c := computer{instructionPtr: 0}
	var startInstructions string
	for i, row := range strings.Split(input, "\n") {
		switch i {
		case 0:
			c.A = parseRegister(row)
		case 1:
			c.B = parseRegister(row)
		case 2:
			c.C = parseRegister(row)
		case 4:
			startInstructions = strings.Split(row, " ")[1]
			c.instructions = parseInstructions(startInstructions)
		}
	}

	for i := 0; i < 10000000000; i++ {
		c.A = i
		c.B = 0
		c.C = 0
		c.instructionPtr = 0
		c.output = make([]int, 0)
		for {
			if err := c.process(); err != nil {
				break
			}
		}
		if printInstructions(c.output) == startInstructions {
			fmt.Println(i)
			break
		}
	}
}
