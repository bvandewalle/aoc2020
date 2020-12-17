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
	mem := map[int64]int64{}
	var ones int64
	var zeroes int64
	for _, l := range input {
		if strings.Contains(l, "mask = ") {
			currentOnes := ""
			currentZeroes := ""
			for _, c := range strings.Split(l, "mask = ")[1] {
				switch c {
				case '0':
					currentOnes += "0"
					currentZeroes += "1"
				case '1':
					currentOnes += "1"
					currentZeroes += "0"
				case 'X':
					currentOnes += "0"
					currentZeroes += "0"
				}
			}

			//fmt.Println(currentOnes)
			//fmt.Println(currentZeroes)
			ones, _ = strconv.ParseInt(currentOnes, 2, 64)
			zeroes, _ = strconv.ParseInt(currentZeroes, 2, 64)
			continue
		}

		n := strings.Split(strings.TrimPrefix(l, "mem["), "] = ")
		//fmt.Println(n)
		n1, _ := strconv.Atoi(n[0])
		n2, _ := strconv.Atoi(n[1])
		n3 := (int64(n2) | ones) &^ zeroes

		fmt.Printf("-----\n")
		fmt.Printf("ones\t%36b\n", ones)
		fmt.Printf("zeroes\t%36b\n", zeroes)
		fmt.Printf("origi\t%36b\n", n2)
		fmt.Printf("result\t%36b\n", n3)

		mem[int64(n1)] = n3
	}

	//fmt.Println(mem)
	var sum int64
	for _, v := range mem {
		sum += v
	}
	fmt.Println(sum)
}

func generateAllAddresses(address int64, x int64) []int64 {
	generated := []int64{address}

	for i := 0; i < 64; i++ {

	}

	return generated
}

func part2(input []string) {
	mem := map[int64]int64{}
	var m int64
	var x int64
	for _, l := range input {
		if strings.Contains(l, "mask = ") {
			mask := ""
			currentX := ""
			for _, c := range strings.Split(l, "mask = ")[1] {
				switch c {
				case '0':
					mask += "0"
					currentX += "0"
				case '1':
					mask += "1"
					currentX += "0"
				case 'X':
					mask += "0"
					currentX += "1"
				}
			}

			//fmt.Println(currentOnes)
			//fmt.Println(currentZeroes)
			m, _ = strconv.ParseInt(mask, 2, 64)
			x, _ = strconv.ParseInt(currentX, 2, 64)
			continue
		}

		n := strings.Split(strings.TrimPrefix(l, "mem["), "] = ")
		//fmt.Println(n)
		n1, _ := strconv.Atoi(n[0])
		n0 := (int64(n1) | mask)
		n2, _ := strconv.Atoi(n[1])

		allAddresses := generateAllAddresses(n0, x)
		for _, a := range allAddresses {
			mem[a] = n2
		}
	}

	//fmt.Println(mem)

	var sum int64
	for _, v := range mem {
		sum += v
	}

	fmt.Println(sum)

}
