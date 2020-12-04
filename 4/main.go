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
	var input []map[string]string
	current := map[string]string{}

	for scanner.Scan() {
		v := scanner.Text()

		if v == "" {
			input = append(input, current)
			current = map[string]string{}
			continue
		}

		p := strings.Split(v, " ")
		for _, e := range p {
			kv := strings.Split(e, ":")
			current[kv[0]] = kv[1]
		}
	}
	input = append(input, current)

	file.Close()

	part1(input)
	part2(input)
}

func part1(input []map[string]string) {
	count := 0

	for _, l := range input {
		if validate1(l) {
			count++
		}
	}

	fmt.Println(count)
}

func part2(input []map[string]string) {
	count := 0

	for _, l := range input {
		if validate1(l) {
			if validate2(l) {
				count++
			}
		}
	}

	fmt.Println(count)
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

func validate2(input map[string]string) bool {
	validatingFuncs := map[string]func(v string) bool{
		"byr": func(v string) bool { return validateInterval(1920, 2002, v) },
		"iyr": func(v string) bool { return validateInterval(2010, 2020, v) },
		"eyr": func(v string) bool { return validateInterval(2020, 2030, v) },
		"hgt": func(v string) bool {
			if strings.Contains(v, "cm") {
				return validateInterval(150, 193, strings.TrimSuffix(v, "cm"))
			} else if strings.Contains(v, "in") {
				return validateInterval(59, 76, strings.TrimSuffix(v, "in"))
			} else {
				return false
			}
		},
		"hcl": func(v string) bool {
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
			return true
		},
		"ecl": func(v string) bool {
			p := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			for _, pp := range p {
				if pp == v {
					return true
				}
			}
			return false
		},
		"pid": func(v string) bool {
			if len(v) != 9 {
				return false
			}
			return validateInterval(0, 999999999, v)
		},
	}

	for k, v := range input {
		if f, exists := validatingFuncs[k]; exists {
			if !f(v) {
				return false
			}
		}
	}

	return true
}

func validateInterval(low int, up int, v string) bool {
	v1, err := strconv.Atoi(v)
	if err != nil {
		return false
	}

	if v1 < low || v1 > up {
		return false
	}

	return true
}
