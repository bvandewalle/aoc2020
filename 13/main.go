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

func part2(input []string) {
	listInts := []int{}
	busses := strings.Split(input[1], ",")

	for i, b := range busses {
		if b == "x" {
			continue
		}

		bInt, _ := strconv.Atoi(b)
		listInts = append(listInts, bInt-i)
	}
	sort.Ints(listInts)

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
