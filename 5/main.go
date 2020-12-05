package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func calculateID(l string) int {
	var row float64
	var side float64

	for i, c := range l {
		if c == 'B' {
			row += math.Exp2(6 - float64(i))
		}
		if c == 'R' {
			side += math.Exp2(9 - float64(i))
		}
	}

	return int(row*8 + side)
}

func part1(input []string) {
	max := 0

	for _, l := range input {
		val := calculateID(l)
		if val > max {
			max = val
		}
	}

	fmt.Println(max)
}

func part2(input []string) {
	max := 0
	min := 1000

	mem := map[int]bool{}

	for _, l := range input {
		val := calculateID(l)
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
		mem[val] = true
	}

	for i := min + 1; i < max-1; i++ {
		if _, exists := mem[i]; exists {
			continue
		}

		if _, existsprevious := mem[i-1]; existsprevious {
			if _, existsafter := mem[i+1]; existsafter {
				fmt.Println(i)
			}
		}
	}
}
