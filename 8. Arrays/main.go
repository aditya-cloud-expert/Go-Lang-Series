package main

import "fmt"

func main() {
	var nums [4]int

	nums[0] = 1

	// array length
	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)

	var vals [4]bool
	vals[1] = true
	fmt.Println(vals)

	var names [4]string
	fmt.Println(names)

	nums1 := [3]int{12, 23, 42}
	fmt.Println(nums1)

	nums2 := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(nums2)
}
