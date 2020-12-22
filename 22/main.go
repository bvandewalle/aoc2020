package main

import "fmt"

func main() {

	part1()
	part2()
}

func part1() {
	q1 := []int{26, 16, 33, 8, 5, 46, 12, 47, 39, 27, 50, 10, 34, 20, 23, 11, 43, 14, 18, 1, 48, 28, 31, 38, 41}
	q2 := []int{45, 7, 9, 4, 15, 19, 49, 3, 36, 25, 24, 2, 21, 37, 35, 44, 29, 13, 32, 22, 17, 30, 42, 40, 6}

	for len(q1) > 0 && len(q2) > 0 {
		c1 := q1[0]
		c2 := q2[0]
		q1 = q1[1:]
		q2 = q2[1:]

		if c1 > c2 {
			q1 = append(q1, c1, c2)
		} else {
			q2 = append(q2, c2, c1)
		}
	}

	wq := q1
	if len(q1) == 0 {
		wq = q2
	}

	result := 0
	for i := range wq {
		result += (i + 1) * wq[len(wq)-1-i]
	}

	fmt.Println(result)
}

func alreadyPlayed(mem *[][][]int, q1, q2 []int) bool {
	q1 = append([]int(nil), q1...)
	q2 = append([]int(nil), q2...)

Exit:
	for _, previous := range *mem {
		if len(previous[0]) == len(q1) && len(previous[1]) == len(q2) {
			for i, e := range previous[0] {
				if q1[i] != e {
					continue Exit
				}

			}
			for i, e := range previous[1] {
				if q2[i] != e {
					continue Exit
				}

			}
			return true
		}
	}

	*mem = append(*mem, [][]int{q1, q2})
	return false
}

//return true if p1 wins
func playRecursive(q1, q2 []int, first bool) bool {
	fmt.Println("subgame", q1, q2)
	mem := [][][]int{}

	for len(q1) > 0 && len(q2) > 0 {
		//fmt.Println(q1, q2)

		c1 := q1[0]
		c2 := q2[0]

		q1 = q1[1:]
		q2 = q2[1:]

		if alreadyPlayed(&mem, q1, q2) {
			fmt.Println("already played")
			return true
		}

		if c1 > len(q1) || c2 > len(q2) {
			if c1 > c2 {
				q1 = append(q1, c1, c2)
			} else {
				q2 = append(q2, c2, c1)
			}
		} else {
			q1subgame := append([]int(nil), q1[:c1]...)
			q2subgame := append([]int(nil), q2[:c2]...)
			if playRecursive(q1subgame, q2subgame, false) {
				q1 = append(q1, c1, c2)
			} else {
				q2 = append(q2, c2, c1)
			}
		}
	}

	if first {
		wq := q1
		if len(q1) == 0 {
			wq = q2
		}
		result := 0
		for i := range wq {
			result += (i + 1) * wq[len(wq)-1-i]
		}
		fmt.Println(result)
	}

	if len(q2) == 0 {
		return true
	}
	return false
}

// copies

func part2() {
	q1 := []int{26, 16, 33, 8, 5, 46, 12, 47, 39, 27, 50, 10, 34, 20, 23, 11, 43, 14, 18, 1, 48, 28, 31, 38, 41}
	q2 := []int{45, 7, 9, 4, 15, 19, 49, 3, 36, 25, 24, 2, 21, 37, 35, 44, 29, 13, 32, 22, 17, 30, 42, 40, 6}

	//q1 := []int{9, 2, 6, 3, 1}
	//q2 := []int{5, 8, 4, 7, 10}

	playRecursive(q1, q2, true)
}
