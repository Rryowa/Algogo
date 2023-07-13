package array

import (
	"fmt"
	"sort"
)

/*
You are given a non-negative integer array nums. In one operation, you must:
Choose a positive integer x such that x is less than or equal to the smallest
non-zero element in nums.
Subtract x from every positive element in nums.
Return the minimum number of operations to make every element in nums equal to 0
Input: nums = [1,5,0,3,5]
Output: 3
*/

func MinimumOperations(nums []int) int {
	var x, c int
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		} else {

			c++
			x = nums[i]
			for j := i; j < len(nums); j++ {
				nums[j] = nums[j] - x
			}
			fmt.Println("sub: ", nums)
		}
	}
	return c
}
