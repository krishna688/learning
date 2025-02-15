package main

import "log"

func main() {

	log.Println("rotated sorted array")

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	for i, v := range nums {
		log.Printf("index %d -> %d", i, v)
	}

	log.Println(findTarget(nums, 2))
}

func findTarget(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // left is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
