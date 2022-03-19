package main

import "fmt"

func rotate(s []int, n int) []int {

	for start, end := 0, n; end < len(s); start, end = start + 1, end + 1 {
		s[end], s[start] = s[start], s[end]
	}

	return s
}

func main2() {
	s := []int{0,1,2,3,4,5}

	fmt.Println(rotate(s, 2))
}