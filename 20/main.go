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

func reverse(s string) string {
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
		border1 = reverse(border1)
		border2 = reverse(border2)
		border3 = reverse(border3)
		border4 = reverse(border4)

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

func rotate(in []string) []string {
	out := []string{}
	for i, l := range in {
		s := ""
		for j := range l {
			s += string(in[len(in)-j-1][i])
		}
		out = append(out, s)
	}
	return out
}

func flipH(in []string) []string {
	out := []string{}
	for i := range in {
		out = append(out, in[len(in)-1-i])
	}
	return out
}

func generateAllTiles(mem map[int][][]string, id int) {
	for i := 0; i < 3; i++ {
		mem[id] = append(mem[id], rotate(mem[id][i]))
	}
	for i := 0; i < 4; i++ {
		mem[id] = append(mem[id], flipH(mem[id][i]))
	}
}

func generateAllGrids(original []string) [][]string {
	allGrids := [][]string{original}
	for i := 0; i < 3; i++ {
		allGrids = append(allGrids, rotate(allGrids[i]))
	}
	for i := 0; i < 4; i++ {
		allGrids = append(allGrids, flipH(allGrids[i]))
	}
	return allGrids
}

func displayTiles(tiles [][]string) {
	for _, t := range tiles {
		for _, l := range t {
			fmt.Println(l)
		}
		fmt.Println("--------")
	}
}

func findTopMatch(memTiles map[int][][]string, memBorderToTiles map[string][]int, remainingTiles map[int]bool, topTile []string) ([]string, int) {
	toMatch := topTile[len(topTile)-1]

	tileID := 0
	var tiles [][]string
	for _, v := range memBorderToTiles[toMatch] {
		if _, exists := remainingTiles[v]; exists {
			tiles = memTiles[v]
			tileID = v
			break
		}
	}

	for _, t := range tiles {
		if t[0] == toMatch {
			return t, tileID
		}
	}

	return nil, 0
}

func findLeftMatch(memTiles map[int][][]string, memBorderToTiles map[string][]int, remainingTiles map[int]bool, leftTile []string) ([]string, int) {
	toMatch := ""
	for _, l := range leftTile {
		toMatch += string(l[len(l)-1])
	}

	tileID := 0
	var tiles [][]string
	for _, v := range memBorderToTiles[toMatch] {
		if _, exists := remainingTiles[v]; exists {
			tiles = memTiles[v]
			tileID = v
			break
		}
	}

	for _, t := range tiles {
		column := ""
		for _, l := range t {
			column += string(l[0])
		}
		if column == toMatch {
			return t, tileID
		}
	}
	return nil, 0
}

func findTopLeft(tiles [][]string, memBorderToTiles map[string][]int) []string {
	for _, t := range tiles {
		column := ""

		for _, l := range t {
			column += string(l[0])
		}

		if len(memBorderToTiles[column]) == 1 {
			if len(memBorderToTiles[t[0]]) == 1 {
				return t
			}
		}
	}
	return nil
}

func addGridRight(grid []string, tile []string) []string {
	for i, l := range tile[1 : len(tile)-1] {
		grid[len(grid)-8+i] += l[1 : len(l)-1]
	}

	return grid
}

func addGridBottom(grid []string, tile []string) []string {
	for i, l := range tile {
		if (i == 0) || (i == len(tile)-1) {
			continue
		}
		grid = append(grid, l[1:len(l)-1])
	}

	return grid
}

func part2(input []string) {
	// Mapping of all possible border to a list of TILE ID having them
	memBorderToTiles := map[string][]int{}
	// Mapping of Tile ID to the 8 possible Tiles rotated and flipped
	memTiles := map[int][][]string{}

	i := 0

	// Filling the maps based on input
	for i < len(input) {
		id, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(input[i], "Tile "), ":"))
		border1 := input[i+1]
		border3 := input[i+10]

		border2 := ""
		border4 := ""

		originalTile := []string{}
		for j := 0; j < 10; j++ {
			border2 += string(input[i+1+j][0])
			border4 += string(input[i+1+j][9])
			originalTile = append(originalTile, input[j+i+1])
		}

		// Generating all the tiles
		memTiles[id] = [][]string{originalTile}
		generateAllTiles(memTiles, id)

		// Generating all the borders
		memBorderToTiles[border1] = append(memBorderToTiles[border1], id)
		memBorderToTiles[border2] = append(memBorderToTiles[border2], id)
		memBorderToTiles[border3] = append(memBorderToTiles[border3], id)
		memBorderToTiles[border4] = append(memBorderToTiles[border4], id)

		// we need to account for flipping
		border1 = reverse(border1)
		border2 = reverse(border2)
		border3 = reverse(border3)
		border4 = reverse(border4)

		memBorderToTiles[border1] = append(memBorderToTiles[border1], id)
		memBorderToTiles[border2] = append(memBorderToTiles[border2], id)
		memBorderToTiles[border3] = append(memBorderToTiles[border3], id)
		memBorderToTiles[border4] = append(memBorderToTiles[border4], id)

		i += 12
	}

	// Memcount1 keeps a count of the amount of times a Tile ID has a border that noone else has.
	memcount1 := map[int]int{}
	for _, b := range memBorderToTiles {
		switch len(b) {
		case 1:
			memcount1[b[0]]++
		}
	}

	// dist keeps a distribution per amount of times a tile ID has a non=common border.
	dist := map[int][]int{}
	for k, c := range memcount1 {
		dist[c] = append(dist[c], k)
	}

	// RemainingTiles contains all the tiles initially
	remainingTiles := map[int]bool{}
	for k := range memTiles {
		remainingTiles[k] = true
	}

	// The final grid
	grid := []string{}

	// The Corder is chosen as the first element with 4 possible borders (out of 8) that are unique. That's how we know it's a corner
	topLeftID := dist[4][0]

	// The Correct Orientation still needs to be found for that corner tile.
	topLeftTile := findTopLeft(memTiles[topLeftID], memBorderToTiles)

	// That tile is not available anymore.
	delete(remainingTiles, topLeftID)

	// The tile is added to the grid, on the bottom left
	grid = addGridBottom(grid, topLeftTile)

	//displayTiles([][]string{topLeftTile})

	// The current left and top tile is the corner tile
	leftTile := topLeftTile
	topTile := topLeftTile
	id := topLeftID
	jj := 0

	// The main logic. Iterate line by line. When filling in the line, we attempt to find a tile that matches the left neighbour. (referenced by leftTile)
	// Once this is impossible, we know it's time to go to the next line.
	// To find the first element of the next line, we attemp to match it to the tile above (referenced by topTile).
	// Once this is impossible, we know the grid is finished.
	for {
		jj++
		ii := 0
		for {
			ii++
			leftTile, id = findLeftMatch(memTiles, memBorderToTiles, remainingTiles, leftTile)
			// id is 0 when no match was found. Go to next line
			if id == 0 {
				break
			}
			delete(remainingTiles, id)

			grid = addGridRight(grid, leftTile)

			//displayTiles([][]string{leftTile})
		}
		topTile, id = findTopMatch(memTiles, memBorderToTiles, remainingTiles, topTile)
		leftTile = topTile
		// id is 0 when no match was found. End of grid!
		if id == 0 {
			break
		}
		delete(remainingTiles, id)

		grid = addGridBottom(grid, leftTile)

		//displayTiles([][]string{leftTile})

	}

	//generate the 8 possible rotation and flips from that grid
	allGrids := generateAllGrids(grid)

	for _, g := range allGrids {
		i := findSeaMonster(g)
		// Will only be non zero if at least one was found.
		if i != 0 {
			fmt.Println(i)
			return
		}
	}
}

// Rough logic : We iterate over the whole grid and see if we can match a Sea Monster char by char....
// If we do, we keep count.
// At the end, we count the amount of '#' from which we remove the amount of sea monsters * 15
// That only works because there are no overlapping sea monsters.
func findSeaMonster(grid []string) int {
	found := false
	count := 0
	countc := 0
	sm := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	for i := range grid[:len(grid)-2] {
	mainGrid:
		for j := range grid[0][:len(grid[0])-len(sm[0])] {

			for ii := range sm {
				for jj := range sm[0] {
					if sm[ii][jj] == '#' {
						if !(grid[i+ii][j+jj] == '#' || grid[i+ii][j+jj] == 'O') {
							continue mainGrid
						}
					}
				}
			}
			found = true
			fmt.Println("FOUND SEA MONSTER starting at", i, j)
			count++
		}
	}

	for _, l := range grid {
		for _, c := range l {
			if c == '#' {
				countc++
			}
		}
	}

	if found {
		return countc - (count * 15)
	}

	return 0

}
