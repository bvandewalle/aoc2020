package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func isSimilar(first, second []string) bool {
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func countOccupied(input []string) int {
	count := 0
	for _, v := range input {
		count += strings.Count(v, "#")
	}
	return count
}

func countAdjacentOccupied(input []string, x, y int) int {
	adjacentx := []int{0, 1, 1, 1, 0, -1, -1, -1}
	adjacenty := []int{1, 1, 0, -1, -1, -1, 0, 1}

	count := 0
	for i := range adjacentx {
		if (adjacentx[i]+x >= 0) && (adjacentx[i]+x < len(input)) {
			if (adjacenty[i]+y >= 0) && (adjacenty[i]+y < len(input[0])) {
				if input[adjacentx[i]+x][adjacenty[i]+y] == '#' {
					count++
				}
			}
		}
	}

	return count
}

func calculateNextPart1(input []string) []string {
	output := make([]string, len(input))

	for x, vi := range input {
		output[x] = ""

		for y, vj := range vi {
			switch vj {
			case '#':
				if countAdjacentOccupied(input, x, y) >= 4 {
					output[x] += string('L')
				} else {
					output[x] += string('#')
				}
			case 'L':
				if countAdjacentOccupied(input, x, y) == 0 {
					output[x] += string('#')
				} else {
					output[x] += string('L')
				}
			case '.':
				output[x] += string('.')
			}
		}
	}

	return output
}

func part1(input []string) {
	output := calculateNextPart1(input)
	for {
		output = calculateNextPart1(input)
		if isSimilar(input, output) {
			break
		}
		input = output
	}
	fmt.Println(countOccupied(input))
}

func countDirectionOccupied(input []string, x, y, directionx, directiony int) int {

	for {
		x += directionx
		y += directiony
		if (x >= 0) && (x < len(input)) {
			if (y >= 0) && (y < len(input[0])) {
				if input[x][y] == '#' {
					return 1
				}
				if input[x][y] == 'L' {
					return 0
				}
			} else {
				return 0
			}
		} else {
			return 0
		}
	}
}

func countAllDirectionOccupied(input []string, x, y int) int {
	adjacentx := []int{0, 1, 1, 1, 0, -1, -1, -1}
	adjacenty := []int{1, 1, 0, -1, -1, -1, 0, 1}

	count := 0

	for i := range adjacentx {
		count += countDirectionOccupied(input, x, y, adjacentx[i], adjacenty[i])
	}

	return count
}

func calculateNextPart2(input []string) []string {
	output := make([]string, len(input))

	for x, vi := range input {
		output[x] = ""

		for y, vj := range vi {
			switch vj {
			case '#':
				if countAllDirectionOccupied(input, x, y) >= 5 {
					output[x] += string('L')
				} else {
					output[x] += string('#')
				}
			case 'L':
				if countAllDirectionOccupied(input, x, y) == 0 {
					output[x] += string('#')
				} else {
					output[x] += string('L')
				}
			case '.':
				output[x] += string('.')
			}
		}
	}

	return output
}

func part2(input []string) {
	output := calculateNextPart2(input)
	for {
		output = calculateNextPart2(input)
		if isSimilar(input, output) {
			break
		}
		input = output
	}
	fmt.Println(countOccupied(input))
}
