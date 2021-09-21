package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var minSum int64
	var maxSum int64
	for _, v := range arr[:4] {
		minSum += int64(v)
	}
	for _, v := range arr[1:] {
		maxSum += int64(v)
	}
	fmt.Println(minSum, maxSum)
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}

	miniMaxSum(arr)
}

