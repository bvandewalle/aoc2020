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

func part1(input []string) {
	mem := map[int][]int{}

	for i, l := range input {
		in, _ := strconv.Atoi(l)
		if v, exists := mem[in]; exists {
			mem[in] = append(v, i)
		} else {
			mem[in] = []int{i}
		}
	}

	for i, l := range input {
		if i < 25 {
			continue
		}
		in, _ := strconv.Atoi(l)

		correct := false

		for _, l := range input[i-25 : i] {

			jn, _ := strconv.Atoi(l)
			if v, exists := mem[in-jn]; exists {
				for _, possib := range v {
					if possib > i-25 && possib < i {
						correct = true
					}
				}
			}
		}

		if !correct {
			fmt.Println(l)
			return
		}

	}

}

func part2(input []string) {
	start := 0
	end := 0

	currentSum, _ := strconv.Atoi(input[0])
	invalidNumber := 507622668

	for end < len(input)-1 {
		if currentSum == invalidNumber && start != end {
			fmt.Println("found")
			calculatePart2(input, start, end)
			return
		}
		if currentSum < invalidNumber {
			end++
			val, _ := strconv.Atoi(input[end])
			currentSum += val
		} else {
			val, _ := strconv.Atoi(input[start])
			start++
			currentSum -= val
		}
	}

}

func calculatePart2(input []string, start, end int) {
	max := 0
	min, _ := strconv.Atoi(input[len(input)-1])
	for _, l := range input[start:end] {
		val, _ := strconv.Atoi(l)
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	fmt.Println(min + max)
}
