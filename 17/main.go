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

func (c *computer) getLiteralOperand() int {
	return c.instructions[c.instructionPtr+1]
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
		op := c.getLiteralOperand()
		c.B = c.B ^ op
	case 2:
		op, err := c.getComboOperand()
		if err != nil {
			return err
		}
		c.B = op % 8
	case 3:
		if c.A != 0 {
			op := c.getLiteralOperand()
			c.instructionPtr = op
			return nil
		}
	case 4:
		_ = c.getLiteralOperand()
		c.B = c.B ^ c.C
	case 5:
		op, err := c.getComboOperand()
		if err != nil {
			return err
		}
		c.output = append(c.output, op%8)
		if len(c.output) > len(c.instructions) || c.output[len(c.output)-1] != c.instructions[len(c.output)-1] {
			// return errors.New("wrong value for A")
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

func (c *computer) run() {
	for {
		fmt.Println(*c)
		if err := c.process(); err != nil {
			fmt.Println(err)
			break
		}
	}
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

func runOnce(A int) (newA, newB, newC, output int) {
	B := A & 0b111
	B = B ^ 0b001
	C := A >> B
	A = A >> 3
	B = B ^ C
	B = B ^ 0b110
	output = B & 0b111
	return A, B, C, output
}

func getSpot(A, i int, input []int) {
	for a := 0; a < 8; a++ {
		newA := A | a
		_, _, _, output := runOnce(newA)
		if output == input[i] {
			if i == 0 {
				fmt.Println(newA)
			} else {
				getSpot(newA<<3, i-1, input)
			}
		}
	}
}

func main() {
	// data, _ := os.ReadFile("test_data")
	// data, _ := os.ReadFile("test_data2")
	data, _ := os.ReadFile("data")

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

	// c.A = 117440
	// c.run()

	fmt.Println(c)

	getSpot(0, 15, c.instructions)
	// A, B, C := c.A, c.B, c.C
	// var output int
	// for A > 0 {
	// 	A, B, C, output = runOnce(A)
	// 	fmt.Println(output)
	// }

}
