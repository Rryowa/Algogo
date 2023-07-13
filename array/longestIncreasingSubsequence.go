package array

/*
Given an integer array nums, return the length of the longest strictly increasing
subsequence.

We will try to create the longest possible tail!
so to speak if we have [10, 5, 8, 3] Answer is 2 [5, 8]
1) Add 10 to tail, len of tail is now = 1, so len of LIS is now 1
2) Replace 10 with 5, cause its not forming increasing subsequence
3) Add 8, len of tail is now = 2, so len of LIS is now 2
4) Replace 5 with 3, so len of LIS is now 2
return 2, cause we care only about len of tail

*/

func LengthOfLIS(nums []int) int {
	var tail []int
	idx := 0
	for _, num := range nums {
		idx = binarySearch(tail, num)
		if idx == len(tail) {
			tail = append(tail, num)
		} else {
			tail[idx] = num
		}
	}
	return len(tail)
}

func binarySearch(tail []int, target int) int {
	n := len(tail)
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if tail[mid] < target {
			left = mid + 1
		} else if tail[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return left
}

func LengthOfLIS2(nums []int) int {
	// Create a list to store the lengths of increasing subsequences
	dp := make([]int, len(nums))

	// Initialize all lengths to 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	// Iterate through the array
	for i := 1; i < len(nums); i++ {
		// Check all previous elements
		for j := 0; j < i; j++ {
			// If the current eleme3nt is greater than the previous element
			// and the length of the increasing subsequence ending at the previous element plus 1
			// is greater than the length of the increasing subsequence ending at the current element
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				// Update the length of the increasing subsequence ending at the current element
				dp[i] = dp[j] + 1
			}
		}
	}

	// Find the maximum length
	maxLength := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}

	return maxLength
}
