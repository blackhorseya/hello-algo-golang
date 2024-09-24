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

func traverse(nums []int) {
	count := 0

	// 透過索引走訪陣列
	for i := 0; i < len(nums); i++ {
		count++
	}

	count = 0
	// 直接走訪陣列元素
	for _, num := range nums {
		count += num
	}

	// 同時走訪資料索引和元素
	for index, num := range nums {
		count += nums[index]
		count += num
	}
}

func find(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
