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

func helper(input []string, x int, y int) int {
	currenty := y
	currentx := x

	lenx := len(input[0])

	count := 0
	for currenty < len(input) {
		if input[currenty][currentx%lenx] == '#' {
			count++
		}
		currenty += y
		currentx += x
	}

	return count
}

func part1(input []string) {
	fmt.Println(helper(input, 3, 1))
}

func part2(input []string) {
	fmt.Println(helper(input, 3, 1) * helper(input, 1, 1) * helper(input, 5, 1) * helper(input, 7, 1) * helper(input, 1, 2))

}
