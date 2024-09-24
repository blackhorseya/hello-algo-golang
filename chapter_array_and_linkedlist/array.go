package chapter_array_and_linkedlist

import (
	"math/rand"
)

func randomAccess(nums []int) (randonNum int) {
	return nums[rand.Intn(len(nums))]
}
