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
	z int
	h int
}

func countAdjacent(grid map[point]bool, xx, yy, zz, hh int) int {
	count := 0

	for h := -1; h <= 1; h++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if (z != 0) || (y != 0) || (x != 0) || (h != 0) {
						if grid[point{x + xx, y + yy, z + zz, h + hh}] {
							count++
						}
					}
				}
			}
		}
	}

	return count
}

func part1(input []string) {
	mem := map[point]bool{}

	for y, l := range input {
		curr := []bool{}
		for x, c := range l {
			if c == '#' {
				mem[point{x, y, 0, 0}] = true
				curr = append(curr, true)
			}
		}
	}

	minx := 0
	maxx := len(input[0]) - 1
	miny := 0
	maxy := len(input) - 1
	minz := 0
	maxz := 0

	for cycle := 0; cycle < 6; cycle++ {
		minx--
		miny--
		minz--
		maxx++
		maxy++
		maxz++
		newMem := map[point]bool{}

		for z := minz; z <= maxz; z++ {
			for y := miny; y <= maxy; y++ {
				for x := minx; x <= maxx; x++ {

					count := countAdjacent(mem, x, y, z, 0)

					active := mem[point{x, y, z, 0}]

					if active {
						if !(count == 2 || count == 3) {
							active = false
						}
					} else {
						if count == 3 {
							active = true
						}
					}

					newMem[point{x, y, z, 0}] = active
				}
			}
		}

		mem = newMem
	}

	count := 0
	for _, v := range mem {
		if v {
			count++
		}
	}

	fmt.Println(count)
}

func part2(input []string) {
	mem := map[point]bool{}

	for y, l := range input {
		curr := []bool{}
		for x, c := range l {
			if c == '#' {
				mem[point{x, y, 0, 0}] = true
				curr = append(curr, true)
			}
		}
	}

	minx := 0
	maxx := len(input[0]) - 1
	miny := 0
	maxy := len(input) - 1
	minz := 0
	maxz := 0
	minh := 0
	maxh := 0

	for cycle := 0; cycle < 6; cycle++ {
		minx--
		miny--
		minz--
		minh--
		maxx++
		maxy++
		maxz++
		maxh++
		newMem := map[point]bool{}

		for h := minh; h <= maxh; h++ {
			for z := minz; z <= maxz; z++ {
				for y := miny; y <= maxy; y++ {
					for x := minx; x <= maxx; x++ {

						count := countAdjacent(mem, x, y, z, h)

						active := mem[point{x, y, z, h}]

						if active {
							if !(count == 2 || count == 3) {
								active = false
							}
						} else {
							if count == 3 {
								active = true
							}
						}

						newMem[point{x, y, z, h}] = active
					}
				}
			}
		}

		mem = newMem
	}

	count := 0
	for _, v := range mem {
		if v {
			count++
		}
	}

	fmt.Println(count)
}
