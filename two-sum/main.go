package main

import (
	"fmt"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func TwoSum(nums []int, target int) []int {
    var visitedSet = map[int]int {};

    for i := 0; i < len(nums); i++ {
        var index, ok = visitedSet[target - nums[i]];

        if ok {
            return []int{index, i}
        }

        visitedSet[nums[i]] = i;
    }

    return nil;
}

func main() {
    arr1 := []int{2, 7, 11, 15}
    fmt.Println("Test Case 1:", TwoSum(arr1, 9)) // Expected: [0,1]

    arr2 := []int{3, 2, 4}
    fmt.Println("Test Case 2:", TwoSum(arr2, 6)) // Expected: [1,2]

    arr3 := []int{3, 3}
    fmt.Println("Test Case 3:", TwoSum(arr3, 6)) // Expected: [0,1]
}
