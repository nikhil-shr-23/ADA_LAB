package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Benchmarker struct{}

func (b *Benchmarker) Measure(name string, f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("Algorithm: %-15s | Time: %s\n", name, elapsed)
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return arr
}

func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000)
	}
	return arr
}

func main() {
	b := &Benchmarker{}
	size := 5000
	data := generateRandomArray(size)

	fmt.Printf("Benchmarking with array size: %d\n", size)
	fmt.Println("--------------------------------------------------")

	// Bubble Sort
	arrBubble := make([]int, size)
	copy(arrBubble, data)
	b.Measure("Bubble Sort", func() {
		BubbleSort(arrBubble)
	})

	// Quick Sort
	arrQuick := make([]int, size)
	copy(arrQuick, data)
	b.Measure("Quick Sort", func() {
		QuickSort(arrQuick)
	})

	// Linear Search
	target := -1 // Worst case
	b.Measure("Linear Search", func() {
		LinearSearch(data, target)
	})
}
