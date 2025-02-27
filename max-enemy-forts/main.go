package main

import (
	"fmt"
)

/* You are given a 0-indexed integer array forts of length n representing the positions of several forts. forts[i] can be -1, 0, or 1 where:

-1 represents there is no fort at the ith position.
0 indicates there is an enemy fort at the ith position.
1 indicates the fort at the ith the position is under your command.
Now you have decided to move your army from one of your forts at position i to an empty position j such that:

0 <= i, j <= n - 1
The army travels over enemy forts only. Formally, for all k where min(i,j) < k < max(i,j), forts[k] == 0.
While moving the army, all the enemy forts that come in the way are captured.

Return the maximum number of enemy forts that can be captured. In case it is impossible to move your army, or you do not have any fort under your command, return 0. */

func captureForts(forts []int) int {
    var maxCaptured = 0
    var lastPos = -1
    var lastVal = 0

    for i, fort := range forts {

        if fort == 0 {
            continue
        }

        if lastPos != -1 && fort != lastVal {
            var captured = i - lastPos - 1

            if captured > maxCaptured {
                maxCaptured = captured
            }
        }

        lastPos = i
        lastVal = fort
    }

    return maxCaptured
}

func main() {
    tests := [][]int{
        {1, 0, 0, 0, -1},
        {0, 1, 0, 1, 0, 0, -1},
        {1, -1, 0, 0, 0, 1},
        {0, 0, 0, 0, 0},
        {1, 1, 1},
        {0, 0, 1, -1},
        {1, 0, 0, -1, 0, 0, 0, 0, 1},
    }

    for _, test := range tests {
        result := captureForts(test)
        fmt.Printf("Input: %v, captured enemy forts: %d\n", test, result)
    }
}
