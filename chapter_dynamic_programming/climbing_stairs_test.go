package chapter_dynamic_programming

import (
	"fmt"
	"testing"
)

func TestClimbingStairsBacktrack(t *testing.T) {
	n := 9
	res := climbingStairsBacktrack(n)
	fmt.Printf("爬 %d 阶楼梯共有 %d 种方案\n", n, res)
}

func TestClimbingStairsDFS(t *testing.T) {
	n := 9
	res := climbingStairsDFS(n)
	fmt.Printf("爬 %d 阶楼梯共有 %d 种方案\n", n, res)
}

func TestClimbingStairsDFSMem(t *testing.T) {
	n := 9
	res := climbingStairsDFSMem(n)
	fmt.Printf("爬 %d 阶楼梯共有 %d 种方案\n", n, res)
}

func TestClimbingStairsDP(t *testing.T) {
	n := 9
	res := climbingStairsDP(n)
	fmt.Printf("爬 %d 阶楼梯共有 %d 种方案\n", n, res)
}

func TestClimbingStairsDPComp(t *testing.T) {
	n := 9
	res := climbingStairsDPComp(n)
	fmt.Printf("爬 %d 阶楼梯共有 %d 种方案\n", n, res)
}

func TestMinCostClimbingStairsDPComp(t *testing.T) {
	cost := []int{0, 1, 10, 1, 1, 1, 10, 1, 1, 10, 1}
	fmt.Printf("输入楼梯的代价列表为 %v\n", cost)

	res := minCostClimbingStairsDP(cost)
	fmt.Printf("爬完楼梯的最低代价为 %d\n", res)

	res = minCostClimbingStairsDPComp(cost)
	fmt.Printf("爬完楼梯的最低代价为 %d\n", res)
}

func TestClimbingStairsConstraintDP(t *testing.T) {
	n := 9
	res := climbingStairsConstraintDP(n)
	fmt.Printf("爬 %d 阶楼梯共有 %d 种方案\n", n, res)
}
