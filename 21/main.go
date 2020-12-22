package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	//part1(input)
	part2(input)
}

func intersect(m1, m2 map[string]bool) map[string]bool {
	newMap := map[string]bool{}

	for k1, v1 := range m1 {
		if _, exists := m2[k1]; exists {
			newMap[k1] = v1
		}
	}

	return newMap
}

func part1(input []string) map[string]map[string]bool {

	mem := map[string]map[string]bool{}
	allIngredients := map[string]int{}

	for _, l := range input {
		ingLine := map[string]bool{}

		ll := strings.Split(l, " (contains ")
		for _, ing := range strings.Split(ll[0], " ") {
			ingLine[ing] = true
			allIngredients[ing]++
		}
		for _, al := range strings.Split(strings.TrimSuffix(ll[1], ")"), ", ") {
			if existingIng, exists := mem[al]; exists {
				mem[al] = intersect(ingLine, existingIng)
			} else {
				mem[al] = ingLine
			}
		}
	}

	for _, v1 := range mem {
		for k2 := range v1 {
			delete(allIngredients, k2)
		}
	}

	//fmt.Println(mem)
	//fmt.Println(allIngredients)

	sum := 0
	for _, v := range allIngredients {
		sum += v
	}

	fmt.Println(sum)
	return mem
}

func part2(input []string) {
	mem := part1(input)

	mapping := map[string]string{}
	mappingAl := map[string]string{}

	for len(mem) > 0 {

		for al, ins := range mem {
			for in := range ins {
				if _, exists := mapping[in]; exists {
					//fmt.Println("deleted: ", in)
					delete(ins, in)
				}
			}
			if len(ins) == 1 {
				for in := range ins {
					//fmt.Println("solved: ", al, in)
					mapping[in] = al
					mappingAl[al] = in
				}
				delete(mem, al)
			}
			break
		}
	}

	keys := make([]string, 0, len(mappingAl))

	for k := range mappingAl {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf(mappingAl[k])
		fmt.Printf(",")
	}
	fmt.Printf("\n")
}
