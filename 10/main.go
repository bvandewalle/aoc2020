package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		v := scanner.Text()
		in, _ := strconv.Atoi(v)
		input = append(input, in)
	}

	file.Close()

	part1(input)
	part2(input)
}

func part1(input []int) {
	sort.Ints(input)
	count1 := 0
	count3 := 1
	previous := 0

	for _, i := range input {
		if i-previous == 1 {
			count1++
		} else if i-previous == 3 {
			count3++
		}
		previous = i
	}

	fmt.Println(count1 * count3)
}

func part2(input []int) {
	sort.Ints(input)
	accum := map[int]int{0: 1}

	for _, i := range input {
		accum[i] = accum[i-1] + accum[i-2] + accum[i-3]
		//fmt.Println(i, accum[i])
	}

	fmt.Println(accum[input[len(input)-1]])
}
