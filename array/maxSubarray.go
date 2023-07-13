package array

/*Given an integer array nums, find the subarray
with the largest sum, and return its sum.*/

func MaxSubarray(nums []int) int {
	sum := nums[0]
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > m {
			m = sum
		}
	}
	return m
}
