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

	var input []int
	for scanner.Scan() {
		v := scanner.Text()
		for _, c := range v {
			val, _ := strconv.Atoi(string(c))
			input = append(input, val)
		}
	}

	file.Close()

	part1(input)
	part2(input)
}

type elem struct {
	val  int
	next *elem
}

func isInSelected(val int, selected *elem) bool {
	if selected.val == val || selected.next.val == val || selected.next.next.val == val {
		return true
	}
	return false
}

func part1(input []int) {
	mem := map[int]*elem{}
	var first *elem
	var previous *elem

	for i, v := range input {
		e := &elem{
			val: v,
		}
		if i == 0 {
			first = e
		} else {
			previous.next = e
		}
		previous = e
		mem[v] = e
	}
	previous.next = first

	current := first

	for i := 0; i < 100; i++ {
		selected := current.next
		current.next = selected.next.next.next
		val := current.val - 1
		if val < 1 {
			val = 9
		}
		for isInSelected(val, selected) {
			val--
			if val < 1 {
				val = 9
			}
		}
		destination := mem[val]
		destination.next, selected.next.next.next = selected, destination.next
		current = current.next
	}
	pr := mem[1].next
	for i := 1; i < 9; i++ {
		fmt.Printf("%d", pr.val)
		pr = pr.next
	}
	fmt.Println()
}

func part2(input []int) {
	mem := map[int]*elem{}
	var first *elem
	var previous *elem

	for i := 0; i < 1000000; i++ {
		e := &elem{}
		if i < len(input) {
			e.val = input[i]
		} else {
			e.val = i + 1
		}
		if i == 0 {
			first = e
		} else {
			previous.next = e
		}
		previous = e
		mem[e.val] = e
	}

	previous.next = first
	current := first

	for i := 0; i < 10000000; i++ {
		selected := current.next
		current.next = selected.next.next.next
		val := current.val - 1
		if val < 1 {
			val = 1000000
		}
		for isInSelected(val, selected) {
			val--
			if val < 1 {
				val = 1000000
			}
		}
		destination := mem[val]
		destination.next, selected.next.next.next = selected, destination.next
		current = current.next
	}
	one := mem[1]
	result := 1
	for i := 1; i <= 3; i++ {
		result *= one.val
		one = one.next
	}

	fmt.Println(result)
}
