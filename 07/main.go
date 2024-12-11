package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Equation struct {
	result  int
	numbers []int
}

func (e Equation) isPossible() bool {
	a, b := e.numbers[0], e.numbers[1]
	if len(e.numbers) == 2 {
		return e.result == a*b || e.result == a+b
	}
	canAdd := true
	canMultiply := true
	if e.result%a != 0 {
		canMultiply = false
	}
	if e.result-a < 0 {
		canAdd = false
	}
	result := false
	if canAdd {
		newEq := Equation{
			result:  e.result - a,
			numbers: e.numbers[1:len(e.numbers)],
		}
		result = result || newEq.isPossible()
	}
	if canMultiply {
		newEq := Equation{
			result:  e.result / a,
			numbers: e.numbers[1:len(e.numbers)],
		}
		result = result || newEq.isPossible()
	}
	return result
}

func (e Equation) isPossibleConcat() bool {
	a, b := e.numbers[0], e.numbers[1]
	rStr := strconv.Itoa(e.result)
	aStr := strconv.Itoa(a)
	if len(e.numbers) == 2 {
		// order?
		str := strconv.Itoa(b) + strconv.Itoa(a)
		return e.result == a*b || e.result == a+b || str == rStr
	}
	canAdd := true
	canMultiply := true
	canConcat := true
	if e.result%a != 0 {
		canMultiply = false
	}
	if e.result-a < 0 {
		canAdd = false
	}
	if len(aStr) > len(rStr) {
		canConcat = false
	}
	if len(rStr) > len(aStr) && rStr[len(rStr)-len(aStr):] != aStr {
		canConcat = false
	}
	result := false
	if canAdd {
		newEq := Equation{
			result:  e.result - a,
			numbers: e.numbers[1:len(e.numbers)],
		}
		result = result || newEq.isPossibleConcat()
	}
	if canMultiply {
		newEq := Equation{
			result:  e.result / a,
			numbers: e.numbers[1:len(e.numbers)],
		}
		result = result || newEq.isPossibleConcat()
	}
	if canConcat {
		newResult, _ := strconv.Atoi(rStr[0 : len(rStr)-len(aStr)])
		newEq := Equation{
			result:  newResult,
			numbers: e.numbers[1:len(e.numbers)],
		}
		result = result || newEq.isPossibleConcat()
	}
	return result
}

func main() {
	file, _ := os.Open("data")
	// file, _ := os.Open("data")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	equations := make([]Equation, 0)
	for scanner.Scan() {
		line := scanner.Text()
		result, _ := strconv.Atoi(strings.Split(line, ": ")[0])
		numbersStr := strings.Split(line, ": ")[1]
		numbers := make([]int, 0)
		for _, val := range strings.Split(numbersStr, " ") {
			number, _ := strconv.Atoi(val)
			numbers = append(numbers, number)
		}
		slices.Reverse(numbers)
		equations = append(equations, Equation{result: result, numbers: numbers})
	}

	sum := 0
	for _, e := range equations {
		possible := e.isPossibleConcat()
		if possible {
			sum += e.result
		}
		fmt.Println(e)
		fmt.Println(possible)
	}
	fmt.Println(sum)
}
