//
package main

import (
	"fmt"
	"sort"
)

func foo(ar []int) []int {
	// Copying
	temp := make([]int, len(ar))
	copy(temp, ar)
	var res []int
	for _, value := range temp {
		if value > 0 {
			for index1, value1 := range temp {
				if value+value1 == 0 {
					temp[index1] = 0
					res = append(res, value)
					break
				}
			}
		}
	}
	return res[:]
}

func foo1(ar []int) []int {
	// Using sorting
	sort.Ints(ar)
	fmt.Println(ar)
	i := len(ar) - 1
	basej := 0
	var res []int
	for {
		val := ar[i]
		if val <= 0 {
			break
		}
		if i == 0 {
			break
		}
		for {
			val1 := ar[basej]
			if val+val1 > 0 {
				break
			}

			if val+val1 == 0 {
				res = append(res, val)
				basej++
				break
			}
		}
		i--
	}

	return res[:]
}

func foo2(ar []int) []int {
	//Using map with counts
	var res []int
	mm := make(map[int]int)
	for _, value := range ar {
		if value < 0 {
			mmval, ok := mm[value]
			if ok {
				mm[value] = mmval + 1
			} else {
				mm[value] = 1
			}
		}
	}
	for _, value := range ar {
		if value <= 0 {
			continue
		}
		mmval, ok := mm[-value]
		if ok && mmval >= 1 {
			res = append(res, value)
			mm[-value] = mmval - 1
		}
	}
	return res
}

func main() {
	//
	var ar = []int{3, 4, 4, -4, -4, 0, -4, 0, -2, 1, 2, -2, 4, 4}
	//var ar = []int{-4, -4, -2, 1, 2, 2, 3, 4, 5} // -> [2, 4]
	result := foo2(ar)
	fmt.Println(result)
}
