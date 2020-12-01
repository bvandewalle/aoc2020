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
	var input []int

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		input = append(input, v)
	}

	file.Close()

	part1(input)
	part2(input)
}

func part1(input []int) {
	for i, iv := range input {
		for _, jv := range input[i:] {
			if iv+jv == 2020 {
				fmt.Println(iv * jv)
			}
		}
	}
}

func part2(input []int) {
	mem := map[int][]int{}
	for i, iv := range input {
		for _, jv := range input[i:] {
			mem[iv+jv] = []int{iv, jv}
		}
	}

	for _, iv := range input {
		for k, kv := range mem {
			if iv+k == 2020 {
				fmt.Println(iv * kv[0] * kv[1])
				return
			}
		}
	}
}
