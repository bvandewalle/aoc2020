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

	//var input []int
	//for scanner.Scan() {
	//	v := scanner.Text()
	//	in, _ := strconv.Atoi(v)
	//	input = append(input, in)
	//}

	//[][]INT
	// var input [][]int
	// for scanner.Scan() {
	// 	line := []int{}
	// 	v := scanner.Text()
	//
	// 	p := strings.Split(v, " ")
	// 	for _, e := range p {
	// 		in, _ := strconv.Atoi(e)
	// 		line = append(line, in)
	// 	}
	//
	// 	input = append(input, line)
	// }

	//[]STRINGS:
	var input []string
	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	// []MAP[STRINGS]STRING
	//var input []map[string]string
	//current := map[string]string{}
	//
	//for scanner.Scan() {
	//	v := scanner.Text()
	//
	//	if v == "" {
	//		input = append(input, current)
	//		current = map[string]string{}
	//		continue
	//	}
	//
	//	p := strings.Split(v, " ")
	//	for _, e := range p {
	//		kv := strings.Split(e, ":")
	//		current[kv[0]] = kv[1]
	//	}
	//}
	//input = append(input, current)

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
