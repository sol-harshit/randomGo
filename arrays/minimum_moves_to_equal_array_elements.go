// Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.
//
// In one move, you can increment n - 1 elements of the array by 1.
//
// Example 1:
//
// Input: nums = [1,2,3]
// Output: 3
// Explanation: Only three moves are needed (remember each move increments two elements):
// [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
//
// Example 2:
//
// Input: nums = [1,1,1]
// Output: 0
//
// Constraints:
//
//	n == nums.length
//	1 <= nums.length <= 105
//	-109 <= nums[i] <= 109
//	The answer is guaranteed to fit in a 32-bit integer.
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1000000000}
	moves := minMoves(nums)
	fmt.Println(moves)
}

func minMoves(nums []int) int {
	sort.Ints(nums)
	length := len(nums)

	moves := checkIfEqual(&nums, length, 0)
	return moves
}

func checkIfEqual(nums *[]int, length int, moves int) int {
	temp := (*nums)[0]
	for i := 0; i < length; i++ {
		if temp != (*nums)[i] {
			moves++
			increment(nums, length)
			place(nums, length)
			moves = checkIfEqual(nums, length, moves)
			break
		}
	}
	return moves
}

func increment(nums *[]int, length int) {
	for i := 0; i < length-1; i++ {
		(*nums)[i]++
	}
}

func place(nums *[]int, length int) {
	// since the array is sorted the largest element will be at the end
	largest := (*nums)[length-1]
	for i := length - 1; i > 0; i-- {
		if largest < (*nums)[i] {
			largest = (*nums)[i]
			(*nums)[i], (*nums)[i+1] = (*nums)[i+1], (*nums)[i]
		}
	}
}
