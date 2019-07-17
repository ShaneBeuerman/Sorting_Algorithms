package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	test := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		test[i] = rand.Intn(10000)
	}
	start := time.Now()
	mergeSort(test)
	end := time.Now()
	fmt.Println(end.Sub(start))

	start = time.Now()
	mergeSortConcurrent(test)
	end = time.Now()
	fmt.Println(end.Sub(start))

}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	arr1 := arr[0 : len(arr)/2]
	arr2 := arr[len(arr)/2 : len(arr)]

	arr1 = mergeSort(arr1)
	arr2 = mergeSort(arr2)

	return merge(arr1, arr2)
}

func mergeSortConcurrent(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	arr1 := arr[0 : len(arr)/2]
	arr2 := arr[len(arr)/2 : len(arr)]

	WaitGroup := sync.WaitGroup{}
	WaitGroup.Add(2)

	go func() {
		arr1 = mergeSort(arr1)
		WaitGroup.Done()
	}()

	go func() {
		arr2 = mergeSort(arr2)
		WaitGroup.Done()
	}()
	WaitGroup.Wait()

	return merge(arr1, arr2)
}

func merge(arr1 []int, arr2 []int) []int {
	arr3 := make([]int, len(arr1)+len(arr2))
	a, b, c := 0, 0, 0
	for a < len(arr1) && b < len(arr2) {
		if arr1[a] <= arr2[b] {
			arr3[c] = arr1[a]
			c++
			a++
		} else {
			arr3[c] = arr2[b]
			c++
			b++
		}
	}

	for a < len(arr1) {
		arr3[c] = arr1[a]
		c++
		a++
	}

	for b < len(arr2) {
		arr3[c] = arr2[b]
		c++
		b++
	}

	return arr3
}

func printArray(arr []int) {
	fmt.Printf("%v", arr)
}
