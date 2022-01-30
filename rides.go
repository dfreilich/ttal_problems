package ttal

import (
	"fmt"
	"sort"
)

func GetMinRides(P []int, S []int) int {
	// write your code in Go 1.4
	sizeMap := map[int][]int{}
	keys := []int{}
	cars := 0
	for i, el := range S {
		val, ok := sizeMap[el]
		if ok {
			sizeMap[el] = append(val, i)
		} else {
			sizeMap[el] = []int{i}
		}
		keys = append(keys, el)
	}
	sort.Ints(keys)

	fmt.Println(P)
	fmt.Println(S)
	fmt.Println(sizeMap)
	fmt.Println(keys)
	for _, key := range keys {
		orig := key
		fmt.Printf("Dealing with key %d\n", key)
		for key > 0 && len(keys) > 0 {
			fmt.Printf("'''val %d\n", key)
			currHi := keys[len(keys)-1]
			els := sizeMap[currHi]
			for i, hi := range els {
				for P[hi] < S[hi] {
					key--
					P[hi]++
				}
				if P[hi] == S[hi] {
					fmt.Printf("Upping car because %d topped off\n", hi)
					cars++
					// pop off list
					els = remove(els, i)
				}
			}

			if len(els) > 0 {
				sizeMap[currHi] = els
			} else {
				delete(sizeMap, currHi)
				keys = keys[:len(keys)-1]
			}
		}
		if len(keys) > 0 {
			keys = keys[1:]
		}
		if key > 0 {
			fmt.Println(key)
			return cars + 1
		} else {
			spliced := remove(sizeMap[orig], 0)
			if len(spliced) > 0 {
				sizeMap[orig] = spliced
			} else {
				delete(sizeMap, orig)
			}
		}
	}
	return cars
}

func remove(s []int, i int) []int {
	if len(s) == 1 {
		return []int{}
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
