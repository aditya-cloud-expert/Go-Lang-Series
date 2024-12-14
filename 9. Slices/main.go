package main

import "fmt"

// Slices are one of the most used constructs
// Slices are noting but the dynamic arrays

func main() {
	// Declared but uninitialized
	var num []int
	fmt.Println(num)
	fmt.Println(len(num))

	var nums1 = make([]int, 2, 5)
	fmt.Println(nums1)
	fmt.Println(cap(nums1)) // cap is means capacity.

	copy(num, nums1)
	fmt.Println(nums1, num)
}
