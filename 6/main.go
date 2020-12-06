package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func part1(input []string) {
	count := 0
	mem := map[rune]bool{}

	for _, l := range input {
		if l == "" {
			count += len(mem)
			mem = map[rune]bool{}
			continue
		}
		for _, c := range l {
			mem[c] = true
		}
	}
	count += len(mem)

	fmt.Println(count)
}

func part2(input []string) {
	count := int
	currentGroupMem := map[rune]bool{}
	currentMem := map[rune]bool{}
	first := true

	for _, l := range input {
		if l == "" {
			count += len(currentGroupMem)
			currentGroupMem = map[rune]bool{}
			currentMem = map[rune]bool{}
			first = true
			continue
		}
		for _, c := range l {
			currentMem[c] = true
			if first {
				currentGroupMem[c] = true
			}
		}

		if first {
			first = false
		}

		for k := range currentGroupMem {
			if _, exists := currentMem[k]; !exists {
				delete(currentGroupMem, k)
			}
		}

		currentMem = map[rune]bool{}
	}
	count += len(currentGroupMem)

	fmt.Println(count)
}
