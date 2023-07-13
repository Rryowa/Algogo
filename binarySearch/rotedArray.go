package binarySearch

func RotedSearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[r] {
			// set nums[i] = T
			// 4 5 6 7 0 1 2
			// F F F F T T T
			// Now we have a border
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
