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
	mem := map[string]string{}
	count := 0

	for _, l := range input {
		if l == "" {
			if validate1(mem) {
				count++
			}
			mem = map[string]string{}
			continue
		}

		p := strings.Split(l, " ")
		//fmt.Println(p)
		for _, e := range p {
			kv := strings.Split(e, ":")
			mem[kv[0]] = kv[1]
		}
	}
	// Off by one: need to validate one more time
	if validate1(mem) {
		count++
	}

	fmt.Println(count)
}

func part2(input []string) {
	mem := map[string]string{}
	count := 0

	for _, l := range input {
		if l == "" {
			fmt.Println("-----")

			if validate1(mem) {
				if validate2(mem) {
					count++
				}
			}

			mem = map[string]string{}
			continue
		}

		p := strings.Split(l, " ")
		for _, e := range p {
			kv := strings.Split(e, ":")
			mem[kv[0]] = kv[1]
		}
	}

	// Off by one: need to validate one more time
	if validate1(mem) {
		if validate2(mem) {
			count++
		}
	}

	fmt.Printf("final count: %d\n", count)
}

func validate1(input map[string]string) bool {
	toValidate := map[string]bool{}
	toValidate["byr"] = false
	toValidate["iyr"] = false
	toValidate["eyr"] = false
	toValidate["hgt"] = false
	toValidate["hcl"] = false
	toValidate["ecl"] = false
	toValidate["pid"] = false

	for k := range input {
		if _, exists := toValidate[k]; exists {
			toValidate[k] = true
		}
	}

	for _, v := range toValidate {
		if !v {
			return false
		}
	}

	return true
}

func validateInterval(low int, up int, v string) bool {
	v1, err := strconv.Atoi(v)
	if err != nil {
		return false
	}

	if v1 < low {
		return false
	}
	if v1 > up {
		return false
	}

	return true
}

func validate2(input map[string]string) bool {
	for k, v := range input {
		fmt.Println(k, v)

		switch k {
		case "byr":
			if !validateInterval(1920, 2002, v) {
				return false
			}
		case "iyr":
			if !validateInterval(2010, 2020, v) {
				return false
			}
		case "eyr":
			if !validateInterval(2020, 2030, v) {
				return false
			}
		case "hgt":
			if strings.Contains(v, "cm") {
				if !validateInterval(150, 193, strings.TrimSuffix(v, "cm")) {
					return false
				}
			} else if strings.Contains(v, "in") {
				if !validateInterval(59, 76, strings.TrimSuffix(v, "in")) {
					return false
				}
			} else {
				return false
			}
		case "hcl":
			for i, c := range v {
				if i == 0 && c != '#' {
					return false
				}
				if i > 0 && (c < '0') && (c > '9') && (c < 'a') && (c > 'f') {
					return false
				}
				if i > 6 {
					return false
				}
			}
		case "ecl":
			p := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			good := false
			for _, pp := range p {
				if pp == v {
					good = true
					break
				}
			}
			if !good {
				return false
			}

		case "pid":
			if len(v) != 9 {
				return false
			}
			if !validateInterval(0, 999999999, v) {
				return false
			}
		}

	}

	return true
}
