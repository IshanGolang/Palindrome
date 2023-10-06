package main

import (
	"fmt"
)

func isPalindrome(arr []int) bool {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		if arr[i] != arr[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	list1 := []int{8, 7, 4, 7, 8}
	list2 := []int{1, 2, 3, 2, 1}

	fmt.Print("List 1 is a Palindrome: ")
	fmt.Println(isPalindrome(list1))

	fmt.Print("List 2 is a Palindrome: ")
	fmt.Println(isPalindrome(list2))
}
