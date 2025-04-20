package main

import "fmt"

// most used constructs in Go
// + useful methods

func main() {

	nums := []int{}
	fmt.Println(nums == nil)
	fmt.Println(cap(nums))

	nums = append(nums, 1, 2, 3, 4)
	fmt.Println(nums)
	fmt.Println(cap(nums))

	nums_2 := [][]int{{1, 2, 3}, {4, 5, 6}} // 2D slice or array
	fmt.Println(nums_2)
	fmt.Println(nums_2[0][1])

}
