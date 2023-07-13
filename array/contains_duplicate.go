package array

/*
Given an integer array nums, return true if any value
appears at least twice in the array, and return false
if every element is distinct.
Input: nums = [1,2,3,1]
Output: true
*/

type void any

func СontainsDuplicate(nums []int) bool {
	set := make(map[int]any)
	var v any

	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = v
	}
	return false
}
