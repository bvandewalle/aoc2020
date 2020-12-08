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

func parseInstruction(l string) (string, int) {
	a := strings.Split(l, " ")
	val, _ := strconv.Atoi(a[1])

	return a[0], val
}

func runProgram(input []string, changeLine int) (bool, int) {
	acc := 0
	current := 0
	exists := false

	mem := map[int]bool{}

	for !exists && current < len(input) {
		mem[current] = true
		inst, val := parseInstruction(input[current])
		if current == changeLine {
			if inst == "nop" {
				inst = "jmp"
			} else if inst == "jmp" {
				inst = "nop"
			}
		}

		switch inst {
		case "nop":
			current++
		case "acc":
			acc += val
			current++
		case "jmp":
			current += val
		}

		_, exists = mem[current]
	}

	return exists, acc
}

func part1(input []string) {
	_, val := runProgram(input, -1)
	fmt.Println(val)
}

func part2(input []string) {
	for i, l := range input {
		if strings.Contains(l, "nop") || strings.Contains(l, "jmp") {
			looped, val := runProgram(input, i)
			if !looped {
				fmt.Println(val)
				return
			}
		}
	}
}
