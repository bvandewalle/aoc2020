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
	x := 0
	y := 0

	dirx := []int{1, 0, -1, 0}
	diry := []int{0, -1, 0, 1}

	currentDir := 0

	for _, l := range input {
		value, _ := strconv.Atoi(l[1:])
		switch l[0] {
		case 'N':
			y += value
		case 'S':
			y -= value
		case 'E':
			x += value
		case 'W':
			x -= value
		case 'L':
			turn := value / 90
			currentDir -= turn
			currentDir %= 4

			if currentDir < 0 {
				currentDir += 4
			}

		case 'R':
			turn := value / 90
			currentDir += turn
			currentDir %= 4
			if currentDir < 0 {
				currentDir += 4
			}

		case 'F':
			x += (dirx[currentDir] * value)
			y += (diry[currentDir] * value)

		}
		//fmt.Println(x, y, currentDir)
	}

	if y < 0 {
		y = -y
	}
	if x < 0 {
		x = -x
	}
	fmt.Println(x + y)
}

func part2(input []string) {
	x := 0
	y := 0

	wx := 10
	wy := 1

	for _, l := range input {
		value, _ := strconv.Atoi(l[1:])
		switch l[0] {
		case 'N':
			wy += value
		case 'S':
			wy -= value
		case 'E':
			wx += value
		case 'W':
			wx -= value
		case 'L':
			turn := value / 90

			for i := 0; i < turn; i++ {
				wx, wy = -wy, wx
			}

		case 'R':
			turn := value / 90

			for i := 0; i < turn; i++ {
				wx, wy = wy, -wx
			}

		case 'F':
			x += value * wx
			y += value * wy
		}
		//fmt.Println(x, y, wx, wy)
	}

	if y < 0 {
		y = -y
	}
	if x < 0 {
		x = -x
	}
	fmt.Println(x + y)
}
