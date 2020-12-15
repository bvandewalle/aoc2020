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

	var input []int
	for scanner.Scan() {
		v := scanner.Text()
		w := strings.Split(v, ",")
		for _, c := range w {
			in, _ := strconv.Atoi(c)
			input = append(input, in)
		}
	}

	file.Close()

	part1(input)
	part2(input)
}

func part1(input []int) {
	run(input, 2020)
}

func part2(input []int) {
	run(input, 30000000)
}

func run(input []int, iter int) {
	mem := map[int][]int{}
	last := 0

	for i := 0; i < iter; i++ {
		if i < len(input) {
			mem[input[i]] = append(mem[input[i]], i+1)
			last = input[i]
			continue
		}

		v := mem[last]
		if len(v) >= 2 {
			last = v[len(v)-1] - v[len(v)-2]
		} else {
			last = 0
		}

		v = mem[last]
		if len(v) >= 2 {
			mem[last] = []int{v[len(v)-1], i + 1}
		} else {
			mem[last] = append(mem[last], i+1)
		}
	}

	fmt.Println(iter, last)
}
