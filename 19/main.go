package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func evaluateHelper(l string, puzzle map[int][][]int, currentRule []int) bool {
	if len(currentRule) == 0 && len(l) == 0 {
		return true
	}

	if len(currentRule) == 0 || len(l) == 0 {
		return false
	}

	toEvaluate := currentRule[0]

	for _, possibility := range puzzle[toEvaluate] {
		switch possibility[0] {
		case -1:
			if l[0] == 'a' {
				if evaluateHelper(l[1:], puzzle, currentRule[1:]) {
					return true
				}
			}
		case -2:
			if l[0] == 'b' {
				if evaluateHelper(l[1:], puzzle, currentRule[1:]) {
					return true
				}
			}
		default:
			if evaluateHelper(l, puzzle, append(possibility, currentRule[1:]...)) {
				return true
			}
		}
	}

	return false
}

func part1(input []string) {
	part1and2(input, false)
}

func part2(input []string) {
	part1and2(input, true)
}

func part1and2(input []string, part2 bool) {
	mem := map[int][][]int{}

	i := 0
	for j, l := range input {
		if l == "" {
			i = j + 1
			break
		}
		s := strings.Split(l, ": ")

		id, _ := strconv.Atoi(s[0])
		c := strings.Split(s[1], " | ")

		entry := [][]int{}

		for _, e := range c {
			choice := []int{}
			n := strings.Split(e, " ")
			for _, nn := range n {
				x := 0
				if nn == "\"a\"" {
					x = -1
				} else if nn == "\"b\"" {
					x = -2
				} else {
					x, _ = strconv.Atoi(nn)
				}
				choice = append(choice, x)
			}
			entry = append(entry, choice)
		}
		mem[id] = entry
	}

	if part2 {
		mem[8] = [][]int{[]int{42}, []int{42, 8}}
		mem[11] = [][]int{[]int{42, 31}, []int{42, 11, 31}}
	}

	count := 0

	for i < len(input) {
		if evaluateHelper(input[i], mem, []int{0}) {
			count++
		}

		i++
	}

	fmt.Println(count)
}
