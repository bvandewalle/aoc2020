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
	k1, _ := strconv.Atoi(input[0])
	k2, _ := strconv.Atoi(input[1])

	v := 1
	l1 := 1
	for {
		v *= 7
		v %= 20201227
		if v == k1 {
			break
		}
		l1++
	}

	e := 1
	for j := 0; j < l1; j++ {
		e *= k2
		e %= 20201227
	}

	fmt.Println(l1)
	fmt.Println(e)
}

func part2(input []string) {

}
