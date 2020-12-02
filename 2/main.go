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

	part2(input)
}

func part1(input []string) {
	count := 0

	for _, l := range input {
		s := strings.FieldsFunc(l, func(r rune) bool {
			return r == ':' || r == '-' || r == ' '
		})

		c := strings.Count(s[3], s[2])
		v1, _ := strconv.Atoi(s[0])
		v2, _ := strconv.Atoi(s[1])

		if c >= v1 && c <= v2 {
			count++
		}

	}

	fmt.Println(count)
}

func part2(input []string) {
	count := 0

	for _, l := range input {
		s := strings.FieldsFunc(l, func(r rune) bool {
			return r == ':' || r == '-' || r == ' '
		})

		good := false

		v1, _ := strconv.Atoi(s[0])
		v2, _ := strconv.Atoi(s[1])
		if s[2][0] == s[3][v1-1] {
			good = !good
		}
		if s[2][0] == s[3][v2-1] {
			good = !good
		}

		if good {
			count++
		}
	}

	fmt.Println(count)
}
