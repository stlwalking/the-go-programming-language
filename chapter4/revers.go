package main

import "fmt"

func reverse(ptr *[10]int) {
	for i,j := 0, len(ptr) - 1; i < j; i,j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func main1() {
	var arr [10]int = [10]int{1,3,5,6,7,8,9,0,11,10}

	reverse(&arr)
	fmt.Println(arr)
	
}