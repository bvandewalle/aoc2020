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

func parseInput(input []string) map[string]map[string]int {
	mem := map[string]map[string]int{}

	for _, l := range input {
		equ := strings.Split(l, " bags contain")
		terms := strings.Split(equ[1], ",")
		equMem := map[string]int{}
		for _, d := range terms {
			e := strings.Split(d, " ")
			i, _ := strconv.Atoi(e[1])
			equMem[e[2]+" "+e[3]] = i
		}
		mem[equ[0]] = equMem
	}

	return mem
}

func recurHelper1(mem map[string]bool, data map[string]map[string]int, current string) {
	for k, v := range data {
		for k2 := range v {
			if k2 == current {
				if _, exists := mem[k]; !exists {
					//fmt.Printf("%s --> %s\n", k, k2)
					mem[k] = true
					recurHelper1(mem, data, k)
				}
			}
		}
	}
}

func part1(input []string) {
	mem := parseInput(input)

	memHold := map[string]bool{
		"shiny gold": true,
	}

	recurHelper1(memHold, mem, "shiny gold")

	fmt.Println(len(memHold) - 1)
}

func recurHelper2(data map[string]map[string]int, current string) int {
	count := 1

	for k, v := range data[current] {
		count += v * recurHelper2(data, k)
	}

	return count
}

func part2(input []string) {
	mem := parseInput(input)

	solution := recurHelper2(mem, "shiny gold")

	fmt.Println(solution - 1)
}
