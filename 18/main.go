package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string
	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	file.Close()

	part1(input)
	part2(input)
}

func operate(val1, val2 int, op rune) int {
	if op == '+' {
		return val1 + val2
	} else {
		return val1 * val2
	}
}

func precedenceP1(o rune) int {
	if o == '+' || o == '*' {
		return 1
	}
	return 0
}

func precedenceP2(o rune) int {
	if o == '+' {
		return 2
	}
	if o == '*' {
		return 1
	}
	return 0
}

// Good implementation based on stacks. Polish inverted...
func evaluate(expression string, precedenceFunc func(rune) int) int {
	stackNumber := []int{}
	stackOp := []rune{}

	for _, c := range expression {

		switch c {
		case ' ':

		case '+':
			for len(stackOp) != 0 && precedenceFunc(stackOp[len(stackOp)-1]) >= precedenceFunc(c) {
				n := len(stackNumber)
				v1 := stackNumber[n-1]
				v2 := stackNumber[n-2]
				stackNumber = stackNumber[:n-2]
				n = len(stackOp)
				op := stackOp[n-1]
				stackOp = stackOp[:n-1]

				stackNumber = append(stackNumber, operate(v1, v2, op))
			}
			stackOp = append(stackOp, '+')

		case '*':
			for len(stackOp) != 0 && precedenceFunc(stackOp[len(stackOp)-1]) >= precedenceFunc(c) {
				n := len(stackNumber)
				v1 := stackNumber[n-1]
				v2 := stackNumber[n-2]
				stackNumber = stackNumber[:n-2]
				n = len(stackOp)
				op := stackOp[n-1]
				stackOp = stackOp[:n-1]

				stackNumber = append(stackNumber, operate(v1, v2, op))
			}
			stackOp = append(stackOp, '*')

		case '(':
			stackOp = append(stackOp, '(')

		case ')':
			for len(stackOp) != 0 && stackOp[len(stackOp)-1] != '(' {
				n := len(stackNumber)
				v1 := stackNumber[n-1]
				v2 := stackNumber[n-2]
				stackNumber = stackNumber[:n-2]
				n = len(stackOp)
				op := stackOp[n-1]
				stackOp = stackOp[:n-1]

				stackNumber = append(stackNumber, operate(v1, v2, op))
			}
			stackOp = stackOp[:len(stackOp)-1]

		default:
			x, _ := strconv.Atoi(string(c))
			stackNumber = append(stackNumber, x)
		}

	}

	for len(stackOp) != 0 {
		n := len(stackNumber)
		v1 := stackNumber[n-1]
		v2 := stackNumber[n-2]
		stackNumber = stackNumber[:n-2]
		n = len(stackOp)
		op := stackOp[n-1]
		stackOp = stackOp[:n-1]

		stackNumber = append(stackNumber, operate(v1, v2, op))
	}

	return stackNumber[0]
}

func part1(input []string) {
	sum := 0

	for _, l := range input {
		sum += evaluate(l, precedenceP1)
	}

	fmt.Println(sum)
}

func part2(input []string) {
	sum := 0

	for _, l := range input {
		sum += evaluate(l, precedenceP2)
	}

	fmt.Println(sum)
}

// first naive solution
func evaluateHelperP1(expression string) (int, int) {
	currentResult := 1
	currentMult := true

	for i := 0; i < len(expression); i++ {
		switch c := expression[i]; c {
		case ' ':

		case '+':
			currentMult = false

		case '*':
			currentMult = true

		case '(':
			jump, x := evaluateHelperP1(expression[i+1:])
			i += jump
			if currentMult {
				currentResult *= x
			} else {
				currentResult += x
			}

		case ')':
			return i + 1, currentResult

		default:
			x, _ := strconv.Atoi(string(c))
			if currentMult {
				currentResult *= x
			} else {
				currentResult += x
			}
		}

	}

	return 0, currentResult
}

// first naive solution
func evaluateP1(expression string) int {
	_, x := evaluateHelperP1(expression)
	return x
}

// first naive solution
func part1naive(input []string) {
	sum := 0

	for _, l := range input {
		sum += evaluateP1(l)
	}

	fmt.Println(sum)
}
