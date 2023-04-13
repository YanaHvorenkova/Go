package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("До сортировки:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("После сортировки:", arr)
}

func quickSort(arr []int, left, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
