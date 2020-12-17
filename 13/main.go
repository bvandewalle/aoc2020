package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func part1(input []string) {
	startTime, _ := strconv.Atoi(input[0])
	busses := strings.Split(input[1], ",")
	bestWait := 100000
	bestID := 0
	for _, b := range busses {
		if b == "x" {
			continue
		}
		id, _ := strconv.Atoi(b)

		if id-(startTime%id) < bestWait {
			bestID = id
			bestWait = id - (startTime % id)
		}
	}

	fmt.Println(bestID * bestWait)
}

func part2stupid(input []string) {
	busses := strings.Split(input[1], ",")
	firstID, _ := strconv.Atoi(busses[0])

	startTime := 0
	for {
		startTime += firstID
		fmt.Println(startTime)
		checkTime := startTime
		for i, b := range busses[1:] {
			checkTime++
			if b == "x" {
				continue
			}
			toCheck, _ := strconv.Atoi(b)
			if checkTime%toCheck != 0 {
				break
			}

			if i == len(busses)-2 {
				fmt.Println(startTime)
			}
		}
	}
}

func part2stupid2(input []string) {
	mem := map[int]int{}
	listInts := []int{}
	busses := strings.Split(input[1], ",")

	for i, b := range busses {
		if b == "x" {
			continue
		}

		bInt, _ := strconv.Atoi(b)
		mem[bInt] = i
		listInts = append(listInts, bInt)
	}
	sort.Ints(listInts)

	fmt.Println(mem)
	fmt.Println(listInts)

	idToCheck := listInts[len(listInts)-1]
	currentTime := 0
	offset := mem[idToCheck]

	fmt.Println(idToCheck)
	fmt.Println(offset)
	for {
		currentTime += idToCheck
		//fmt.Println(currentTime)

		for i := (len(listInts) - 2); i >= 0; i-- {
			if ((currentTime - offset + mem[listInts[i]]) % listInts[i]) != 0 {
				break
			}

			if i == 0 {
				fmt.Println(currentTime - offset)
				return
			}
		}
	}
}

func inv(a int, m int) int {
	m0 := m
	var t, q int
	x0 := 0
	x1 := 1

	if m == 1 {
		return 0
	}

	// Apply extended Euclid Algorithm
	for a > 1 {
		// q is quotient
		q = a / m

		t = m

		// m is remainder now, process
		// same as euclid's algo
		m = a % m
		a = t

		t = x0

		x0 = x1 - q*x0

		x1 = t
	}

	// Make x1 positive
	if x1 < 0 {
		x1 += m0
	}

	return x1
}

// After googling a lot, I recognized this is the chinese reminder theorem!
func part2(input []string) {
	mem := map[int]int{}
	listInts := []int{}
	busses := strings.Split(input[1], ",")

	for i, b := range busses {
		if b == "x" {
			continue
		}

		bInt, _ := strconv.Atoi(b)
		mem[bInt] = bInt - i
		listInts = append(listInts, bInt)
	}

	fmt.Println(mem)
	fmt.Println(listInts)

	prod := 1
	for _, n := range listInts {
		prod *= n
	}
	fmt.Println(prod)

	result := 0
	for k, v := range mem {
		pp := prod / k
		result += v * pp * inv(pp, k)
	}

	fmt.Println(result % prod)
}
