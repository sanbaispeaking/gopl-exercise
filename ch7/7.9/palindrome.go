package main

import (
	"fmt"
	"sort"
)

func isPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i <= j; {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(sort.IntSlice([]int{1, 2, 3, 4, 3, 2, 1})))
	fmt.Println(isPalindrome(sort.IntSlice([]int{1, 4, 2, 1})))
	fmt.Println(isPalindrome(sort.IntSlice([]int{})))
}
