package chapter_dynamic_programming

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amt := 4

	// 动态规划
	res := coinChangeDP(coins, amt)
	fmt.Printf("凑到目标金额所需的最少硬币数量为 %d\n", res)

	// 空间优化后的动态规划
	res = coinChangeDPComp(coins, amt)
	fmt.Printf("凑到目标金额所需的最少硬币数量为 %d\n", res)
}

func TestCoinChangeII(t *testing.T) {
	coins := []int{1, 2, 5}
	amt := 5

	// 动态规划
	res := coinChangeIIDP(coins, amt)
	fmt.Printf("凑到目标金额的硬币组合数为 %d\n", res)

	// 空间优化后的动态规划
	res = coinChangeIIDPComp(coins, amt)
	fmt.Printf("凑到目标金额的硬币组合数为 %d\n", res)
}
