package chapter_array_and_linkedlist

import (
	"math/rand"
)

func randomAccess(nums []int) (randonNum int) {
	// generate a random index
	randomIndex := rand.Intn(len(nums))

	// get the random number
	randomNum := nums[randomIndex]

	return randomNum
}

func insert(nums []int, num int, index int) {
	// move the numbers after the index
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}

	// insert the number
	nums[index] = num
}

func remove(nums []int, index int) {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}
