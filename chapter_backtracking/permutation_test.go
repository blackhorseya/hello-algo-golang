package chapter_backtracking

import (
	"fmt"
	"testing"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func TestPermutationI(t *testing.T) {
	/* 全排列 I */
	nums := []int{1, 2, 3}
	fmt.Printf("输入数组 nums = ")
	PrintSlice(nums)

	res := permutationsI(nums)
	fmt.Printf("所有排列 res = ")
	fmt.Println(res)
}

func TestPermutationII(t *testing.T) {
	nums := []int{1, 2, 2}
	fmt.Printf("输入数组 nums = ")
	PrintSlice(nums)

	res := permutationsII(nums)
	fmt.Printf("所有排列 res = ")
	fmt.Println(res)
}
