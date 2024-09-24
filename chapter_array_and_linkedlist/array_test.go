package chapter_array_and_linkedlist

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// initialize an array
	var arr [5]int
	fmt.Println("array arr =", arr)

	nums := []int{1, 3, 2, 5, 4}
	fmt.Println("slice nums =", nums)
}
