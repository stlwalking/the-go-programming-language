// exercise 4.5

package main

import "fmt"

func _remove(s []int, i int)  []int{
	copy(s[i:], s[i+1:])

	// 这里copy时后面的元素会变多
	return s[:len(s)-1]
}


func remove(s []int) []int {
	for i := 0; i < len(s) - 1; i++ {
		if s[i] == s[i+1] {
			s = _remove(s, i)
		}
	}

	return s
}

func main() {
	a := []int{1,1,2,2,3,4,5}

	fmt.Println(remove(a))
}
