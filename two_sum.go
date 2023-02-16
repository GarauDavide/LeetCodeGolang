package main

import "fmt"

func main() {

	v := []int{1, 7, 4, 2, 9, 8, 5}

	fmt.Println(twoSum(v, 9))
}

func twoSum(nums []int, target int) []int {

	ln := len(nums)
	max_len := ln - 1

	for i := 0; i < ln; i++ {

		z := max_len

		for z > i {

			if (nums[i] + nums[z]) == target {

				return []int{i, z}
			}

			z--
		}
	}

	return []int{0, 0}
}
