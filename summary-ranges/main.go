package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
    var result []string;;
    var start *int = nil;

    for i := 0; i < len(nums); i++ {
        if i + 1 < len(nums) && nums[i] + 1 == nums[i + 1] {
            if start == nil {
                start = &nums[i];
            }
        } else if start != nil {
            result = append(result, fmt.Sprintf("%d->%d", *start, nums[i]));

            start = nil;
        } else {
            result = append(result, fmt.Sprintf("%d", nums[i]));
        }
    }

    return result;
}

func main() {
    fmt.Println("Test1:", summaryRanges([]int{0, 1, 2, 4, 5, 7})) // [0->2 4->5 7]

    fmt.Println("Test2:", summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // [0 2->4 6 8->9]

    fmt.Println("Test3:", summaryRanges([]int{-2147483648,-2147483647,2147483647})) // [-2147483648->-2147483647 2147483647]
}
