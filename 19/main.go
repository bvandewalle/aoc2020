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

func evaluateHelper(l string, puzzle map[int][][]int, currentRule int) (int, bool) {

}

func evaluate(l string, puzzle map[int][][]int, currentRule []int) (bool, int) {
	success := false

	for _, id := range currentRule {
		if id >= 0 {

		}
	}
}

func part1(input []string) {
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

	count := 0

	for i < len(input) {
		if evaluate(input[i], mem) {
			count++
		}

		i++
	}

	fmt.Println(mem)
	fmt.Println(count)
}

func part2(input []string) {

}
