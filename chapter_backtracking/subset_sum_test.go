package chapter_backtracking

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/blackhorseya/hello-algo-golang/pkg"
)

func TestSubsetSumINaive(t *testing.T) {
	nums := []int{3, 4, 5}
	target := 9
	res := subsetSumINaive(nums, target)

	fmt.Printf("target = " + strconv.Itoa(target) + ", 输入数组 nums = ")
	PrintSlice(nums)

	fmt.Println("所有和等于 " + strconv.Itoa(target) + " 的子集 res = ")
	for i := range res {
		PrintSlice(res[i])
	}
	fmt.Println("请注意，该方法输出的结果包含重复集合")
}

func TestSubsetSumI(t *testing.T) {
	nums := []int{3, 4, 5}
	target := 9
	res := subsetSumI(nums, target)

	fmt.Printf("target = " + strconv.Itoa(target) + ", 输入数组 nums = ")
	PrintSlice(nums)

	fmt.Println("所有和等于 " + strconv.Itoa(target) + " 的子集 res = ")
	for i := range res {
		PrintSlice(res[i])
	}
}
