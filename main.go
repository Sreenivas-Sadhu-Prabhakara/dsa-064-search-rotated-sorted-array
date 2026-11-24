package main

import "fmt"

// SearchRotated returns the index of target in a sorted array of distinct
// values that was rotated at an unknown pivot, or -1. At every probe, one
// half is guaranteed sorted; check whether the target lies inside it.
func SearchRotated(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // left half is sorted
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right half is sorted
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(SearchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
