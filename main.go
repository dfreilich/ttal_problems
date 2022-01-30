package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	if err := test(); err != nil {
		fmt.Println(err)
	}
}

type testCase struct {
	p    []int
	q    []int
	resp int
}

func test() error {
	for _, c := range []testCase{
		{[]int{1, 4, 1}, []int{1, 5, 1}, 2},
		{[]int{4, 4, 2, 4}, []int{5, 5, 2, 5}, 3},
		{[]int{2, 3, 4, 2}, []int{2, 5, 7, 2}, 2},
	} {
		el := Solution(c.p, c.q)
		if c.resp != el {
			return errors.New(fmt.Sprintf("Expected %d, got %d, from %+v", c.resp, el, c))
		}
	}
	return nil
}

func Solution(P []int, S []int) int {
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
		if key > 0 && {
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

// String test
// type testCase struct {
// 	message string
// 	k       int
// 	resp    string
// }
//
// func test() error {
// 	for _, c := range []testCase{
// 		{"Codility We test coders", 14, "Codility We"},
// 		{"Why not   ", 100, "Why not"},
// 		{"The quick brown fox jumps over the lazy dog", 39, "The quick brown fox jumps over the lazy"},
// 		{"To crop or      not to crop", 21, "To crop or not"},
// 	} {
// 		el := Solution(c.message, c.k)
// 		if c.resp != el {
// 			return errors.New(fmt.Sprintf("Expected %s, got %s, from message %s with k %d", c.resp, el, c.message, c.k))
// 		}
// 	}
// 	return nil
// }
//
// func Solution(message string, K int) string {
// 	if K > len(message) {
// 		return strings.TrimSpace(message)
// 	}
// 	arr := strings.Split(message, " ")
// 	buf := &bytes.Buffer{}
// 	written := 0
// 	for i := 0; i < len(arr) && written+len(arr[i]) < K; i++ {
// 		if i > 0 {
// 			buf.WriteString(" ")
// 		}
// 		buf.WriteString(arr[i])
// 	}
// 	return strings.TrimSpace(buf.String())
// }
