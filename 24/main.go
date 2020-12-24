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

type point struct {
	x int
	y int
}

func part1(input []string) {
	instructions := [][]int{}

	moveX := []int{2, 1, -1, -2, -1, 1}
	moveY := []int{0, -1, -1, 0, 1, 1}

	for _, l := range input {
		lineInput := []int{}
		for i := 0; i < len(l); i++ {
			if l[i] == 's' {
				if i+1 < len(l) && l[i+1] == 'e' {
					lineInput = append(lineInput, 1)
					i++
				} else if i+1 < len(l) && l[i+1] == 'w' {
					lineInput = append(lineInput, 2)
					i++
				} else {
					fmt.Println("ISSUE S")
				}
			} else if l[i] == 'n' {
				if i+1 < len(l) && l[i+1] == 'e' {
					lineInput = append(lineInput, 5)
					i++
				} else if i+1 < len(l) && l[i+1] == 'w' {
					lineInput = append(lineInput, 4)
					i++
				} else {
					fmt.Println("ISSUE W")
				}
			} else if l[i] == 'e' {
				lineInput = append(lineInput, 0)
			} else if l[i] == 'w' {
				lineInput = append(lineInput, 3)
			} else {
				fmt.Println("ISSUE nOTHING")
			}
		}
		instructions = append(instructions, lineInput)
	}

	mem := map[point]bool{}

	for _, ins := range instructions {
		x := 0
		y := 0
		for _, m := range ins {
			x += moveX[m]
			y += moveY[m]
		}
		fmt.Println(x, y)

		p := point{
			x: x,
			y: y,
		}
		if v, exists := mem[p]; exists {
			mem[p] = !v
		} else {
			mem[p] = true
		}
	}

	count := 0
	for _, my := range mem {
		if my {
			count++
		}
	}

	fmt.Println(count)
}

func part2(input []string) {
	instructions := [][]int{}

	moveX := []int{2, 1, -1, -2, -1, 1}
	moveY := []int{0, -1, -1, 0, 1, 1}

	for _, l := range input {
		lineInput := []int{}
		for i := 0; i < len(l); i++ {
			if l[i] == 's' {
				if i+1 < len(l) && l[i+1] == 'e' {
					lineInput = append(lineInput, 1)
					i++
				} else if i+1 < len(l) && l[i+1] == 'w' {
					lineInput = append(lineInput, 2)
					i++
				} else {
					fmt.Println("ISSUE S")
				}
			} else if l[i] == 'n' {
				if i+1 < len(l) && l[i+1] == 'e' {
					lineInput = append(lineInput, 5)
					i++
				} else if i+1 < len(l) && l[i+1] == 'w' {
					lineInput = append(lineInput, 4)
					i++
				} else {
					fmt.Println("ISSUE W")
				}
			} else if l[i] == 'e' {
				lineInput = append(lineInput, 0)
			} else if l[i] == 'w' {
				lineInput = append(lineInput, 3)
			} else {
				fmt.Println("ISSUE nOTHING")
			}
		}
		instructions = append(instructions, lineInput)
	}

	mem := map[point]bool{}

	for _, ins := range instructions {
		x := 0
		y := 0
		for _, m := range ins {
			x += moveX[m]
			y += moveY[m]
		}

		p := point{
			x: x,
			y: y,
		}
		if v, exists := mem[p]; exists {
			mem[p] = !v
		} else {
			mem[p] = true
		}
	}

	for i := 0; i < 100; i++ {
		newMem := map[point]bool{}
		for p, currentlyBlack := range mem {
			x := p.x
			y := p.y
			c := adjacentBlacks(mem, x, y)
			if currentlyBlack {
				if !(c == 0 || c > 2) {
					newMem[point{x: x, y: y}] = true
				}
			} else {
				if c == 2 {
					newMem[point{x: x, y: y}] = true
				}
			}
			for i := range moveX {

				xx := x + moveX[i]
				yy := y + moveY[i]
				currentlyBlack = mem[point{x: xx, y: yy}]
				c := adjacentBlacks(mem, xx, yy)
				if currentlyBlack {
					if !(c == 0 || c > 2) {
						newMem[point{x: xx, y: yy}] = true
					}
				} else {
					if c == 2 {
						newMem[point{x: xx, y: yy}] = true
					}
				}
			}
		}
		mem = newMem

		count := 0
		for _, my := range mem {
			if my {
				count++
			}
		}

		fmt.Println(i+1, count)
	}

}

func adjacentBlacks(mem map[point]bool, x, y int) int {

	moveX := []int{2, 1, -1, -2, -1, 1}
	moveY := []int{0, -1, -1, 0, 1, 1}

	count := 0
	for i := range moveX {
		p := point{x: x + moveX[i], y: y + moveY[i]}
		if v, exists := mem[p]; exists {
			if v {
				count++
			}
		}
	}

	return count
}
