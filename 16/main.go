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

func part1(input []string) {
	constraints := map[string][]int{}
	for _, l := range input {

		if l == "" {
			break
		}

		n := strings.Split(l, ": ")
		c := strings.Split(n[1], " or ")
		p1 := strings.Split(c[0], "-")
		p2 := strings.Split(c[1], "-")

		p1p1, _ := strconv.Atoi(p1[0])
		p1p2, _ := strconv.Atoi(p1[1])
		p2p1, _ := strconv.Atoi(p2[0])
		p2p2, _ := strconv.Atoi(p2[1])

		constraints[n[0]] = []int{p1p1, p1p2, p2p1, p2p2}
	}

	sum := 0

	process := false
	for _, l := range input {
		if l == "nearby tickets:" {
			process = true
			continue
		}
		if !process {
			continue
		}

		for _, p := range strings.Split(l, ",") {
			pint, _ := strconv.Atoi(p)
			constraintFailed := true

			for _, v := range constraints {
				if (v[0] <= pint && pint <= v[1]) || (v[2] <= pint && pint <= v[3]) {
					constraintFailed = false
				}
			}

			if constraintFailed {
				sum += pint
			}
		}
	}

	fmt.Println(sum)

	//fmt.Println(constraints)
}

func part2(input []string) {
	constraints := map[string][]int{}
	allTickets := [][]int{}
	myTicket := []int{}

	// processing constraints
	for _, l := range input {
		if l == "" {
			break
		}

		n := strings.Split(l, ": ")
		c := strings.Split(n[1], " or ")
		p1 := strings.Split(c[0], "-")
		p2 := strings.Split(c[1], "-")

		p1p1, _ := strconv.Atoi(p1[0])
		p1p2, _ := strconv.Atoi(p1[1])
		p2p1, _ := strconv.Atoi(p2[0])
		p2p2, _ := strconv.Atoi(p2[1])

		constraints[n[0]] = []int{p1p1, p1p2, p2p1, p2p2}
	}

	// processing my tickets
	process := false
	for _, l := range input {
		if l == "your ticket:" {
			process = true
			continue
		}
		if !process {
			continue
		}

		for _, p := range strings.Split(l, ",") {
			pint, _ := strconv.Atoi(p)
			myTicket = append(myTicket, pint)
		}

		break
	}

	// processing nearby tickets
	process = false
	for _, l := range input {
		if l == "nearby tickets:" {
			process = true
			continue
		}
		if !process {
			continue
		}

		ticketEntry := []int{}
		ticketValid := true

		for _, p := range strings.Split(l, ",") {
			pint, _ := strconv.Atoi(p)
			ticketEntry = append(ticketEntry, pint)

			localConstraintFailed := true

			for _, v := range constraints {
				if (v[0] <= pint && pint <= v[1]) || (v[2] <= pint && pint <= v[3]) {
					localConstraintFailed = false
				}
			}

			if localConstraintFailed {
				ticketValid = false
				break
			}
		}

		if ticketValid {
			allTickets = append(allTickets, ticketEntry)
		}
	}

	constraintsPossible := map[string]map[int]bool{}

	// generating list of possibilities
	for k, v := range constraints {
		possibles := map[int]bool{}

		for i := 0; i < len(allTickets[0]); i++ {
			possible := true
			for _, t := range allTickets {
				if !((v[0] <= t[i] && t[i] <= v[1]) || (v[2] <= t[i] && t[i] <= v[3])) {
					possible = false
					break
				}
			}

			if possible {
				possibles[i] = true
			}
		}

		constraintsPossible[k] = possibles
	}

	// iterating and elimination process
	for pass := 0; pass < 50; pass++ {
		for k, v := range constraintsPossible {
			if len(v) == 1 {
				fmt.Println(k, v)
				for k1 := range v {
					for k2, v2 := range constraintsPossible {
						if k != k2 {
							delete(v2, k1)
						}
					}
				}
			}
		}
	}

	resp := 1

	// calculating result
	for k, v := range constraintsPossible {
		if strings.Contains(k, "departure") {
			fmt.Println(k, v)
			for k1 := range v {
				resp *= myTicket[k1]
			}
		}
	}

	fmt.Println(resp)
}
