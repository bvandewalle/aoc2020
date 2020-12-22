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

	part1(input)
	part2(input)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func part1(input []string) {
	memBorder := map[string][]int{}
	memTile := map[int][]string{}

	i := 0

	for i < len(input) {
		id, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(input[i], "Tile "), ":"))
		border1 := input[i+1]
		border3 := input[i+10]

		border2 := ""
		border4 := ""

		for j := 0; j < 10; j++ {
			border2 += string(input[i+1+j][0])
			border4 += string(input[i+1+j][9])
		}

		memTile[id] = []string{border1, border2, border3, border4}
		memBorder[border1] = append(memBorder[border1], id)
		memBorder[border2] = append(memBorder[border2], id)
		memBorder[border3] = append(memBorder[border3], id)
		memBorder[border4] = append(memBorder[border4], id)

		// we need to account for flipping
		border1 = Reverse(border1)
		border2 = Reverse(border2)
		border3 = Reverse(border3)
		border4 = Reverse(border4)

		memTile[id] = append(memTile[id], border1, border2, border3, border4)
		memBorder[border1] = append(memBorder[border1], id)
		memBorder[border2] = append(memBorder[border2], id)
		memBorder[border3] = append(memBorder[border3], id)
		memBorder[border4] = append(memBorder[border4], id)

		i += 12
	}

	memcount1 := map[int]int{}

	// Calculating all the tiles that have common borders
	for _, b := range memBorder {
		switch len(b) {
		case 1:
			memcount1[b[0]]++
		}
	}

	// finding the tiles that have only 2 common borders
	m := 1
	for id, c := range memcount1 {
		if c == 4 {
			m *= id
		}
	}

	fmt.Println(m)
}

func part2(input []string) {
	memBorderToTiles := map[string][]int{}
	memTileToBorders := map[int][]string{}
	memFullTiles := map[int][]string{}

	i := 0

	for i < len(input) {
		id, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(input[i], "Tile "), ":"))
		border1 := input[i+1]
		border3 := input[i+10]

		border2 := ""
		border4 := ""

		for j := 0; j < 10; j++ {
			border2 += string(input[i+1+j][0])
			border4 += string(input[i+1+j][9])
			memFullTiles[id] = append(memFullTiles[j], input[j])
		}

		memTileToBorders[id] = []string{border1, border2, border3, border4}
		memBorderToTiles[border1] = append(memBorderToTiles[border1], id)
		memBorderToTiles[border2] = append(memBorderToTiles[border2], id)
		memBorderToTiles[border3] = append(memBorderToTiles[border3], id)
		memBorderToTiles[border4] = append(memBorderToTiles[border4], id)

		// we need to account for flipping
		border1 = Reverse(border1)
		border2 = Reverse(border2)
		border3 = Reverse(border3)
		border4 = Reverse(border4)

		memTileToBorders[id] = append(memTileToBorders[id], border1, border2, border3, border4)
		memBorderToTiles[border1] = append(memBorderToTiles[border1], id)
		memBorderToTiles[border2] = append(memBorderToTiles[border2], id)
		memBorderToTiles[border3] = append(memBorderToTiles[border3], id)
		memBorderToTiles[border4] = append(memBorderToTiles[border4], id)

		i += 12
	}

	memcount1 := map[int]int{}

	// Calculating all the tiles that have common borders
	for _, b := range memBorderToTiles {
		switch len(b) {
		case 1:
			memcount1[b[0]]++
		}
	}

	dist := map[int][]int{}
	for k, c := range memcount1 {
		dist[c] = append(dist[c], k)
	}

	fmt.Println(dist)

	remainingTiles := map[int]bool{}
	for k := range memTileToBorders {
		remainingTiles[k] = true
	}

	currentElem := dist[4][0]
	delete(remainingTiles, currentElem)

	grid := []string{}

	fmt.Println(remainingTiles)
}

func findNext(N, E, S, W string, memBorderToTiles map[int][]string, remainingTiles map[int]bool) {

}
