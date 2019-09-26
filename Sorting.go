package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	test := make([]int, 10000)
	randomize(test)

	start := time.Now()
	mergesort(test)
	end := time.Now()
	fmt.Println("The amount of time merge sort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	mergesortConcurrent(test)
	end = time.Now()
	fmt.Println("The amount of time concurrent merge sort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	bubblesort(test)
	end = time.Now()
	fmt.Println("The amount of time bubblesort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	insertionsort(test)
	end = time.Now()
	fmt.Println("The amount of time insertionsort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	selectionsort(test)
	end = time.Now()
	fmt.Println("The amount of time selectionsort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	quicksort(test, 0, len(test)-1)
	end = time.Now()
	fmt.Println("The amount of time quicksort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	quicksortConcurrent(test, 0, len(test)-1)
	end = time.Now()
	fmt.Println("The amount of time concurrent quicksort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	randomQuicksort(test, 0, len(test)-1)
	end = time.Now()
	fmt.Println("The amount of time random quicksort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	randomQuicksortConcurrent(test, 0, len(test)-1)
	end = time.Now()
	fmt.Println("The amount of time concurrent random quicksort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	cocktailsort(test)
	end = time.Now()
	fmt.Println("The amount of time cocktail sort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	heapsort(test)
	end = time.Now()
	fmt.Println("The amount of time heapsort takes is", end.Sub(start))
	randomize(test)

	start = time.Now()
	shellsort(test)
	end = time.Now()
	fmt.Println("The amount of time shell short takes is", end.Sub(start))
	randomize(test)

}

func heapsort(arr []int) []int {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		temp := arr[0]
		arr[0] = arr[i]
		arr[i] = temp
		heapify(arr, i, 0)
	}

	return arr
}

func heapify(arr []int, n int, index int) {
	largest := index
	l := 2*index + 1
	r := 2*index + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != index {
		temp := arr[index]
		arr[index] = arr[largest]
		arr[largest] = temp

		heapify(arr, n, largest)
	}

}

func cocktailsort(arr []int) []int {
	swapped := true
	start := 0
	end := len(arr)
	for swapped {
		swapped = false

		for i := start; i < end-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				swapped = true
			}
		}
		if swapped == false {
			break
		}
		swapped = false
		end = end - 1
		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				swapped = true
			}
		}
		start = start + 1
	}
	return arr
}

func shellsort(arr []int) []int {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			temp := arr[i]
			j := 0
			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = temp
		}
	}
	return arr
}

func isItSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func randomize(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10000)
	}
	return arr
}

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func insertionsort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = value
	}
	return arr
}

func selectionsort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		temp := arr[index]
		arr[index] = arr[i]
		arr[i] = temp
	}
	return arr
}

func quicksort(arr []int, lo int, hi int) []int {
	if lo < hi {
		pivot := partition(arr, lo, hi)

		arr = quicksort(arr, lo, pivot-1)
		arr = quicksort(arr, pivot+1, hi)
	}
	return arr
}

func quicksortConcurrent(arr []int, lo int, hi int) []int {
	if lo < hi {
		pivot := partition(arr, lo, hi)

		WaitGroup := sync.WaitGroup{}
		WaitGroup.Add(2)

		go func() {
			arr = quicksortConcurrent(arr, lo, pivot-1)
			WaitGroup.Done()
		}()

		go func() {
			arr = quicksortConcurrent(arr, pivot+1, hi)
			WaitGroup.Done()
		}()

		WaitGroup.Wait()
	}
	return arr
}

func randomQuicksort(arr []int, lo int, hi int) []int {
	if lo < hi {
		pivot := randomPartition(arr, lo, hi)
		randomQuicksort(arr, lo, pivot-1)
		randomQuicksort(arr, pivot+1, hi)
	}
	return arr
}

func randomQuicksortConcurrent(arr []int, lo int, hi int) []int {
	if lo < hi {
		pivot := randomPartition(arr, lo, hi)

		WaitGroup := sync.WaitGroup{}
		WaitGroup.Add(2)

		go func() {
			arr = randomQuicksortConcurrent(arr, lo, pivot-1)
			WaitGroup.Done()
		}()

		go func() {
			arr = randomQuicksortConcurrent(arr, pivot+1, hi)
			WaitGroup.Done()
		}()

		WaitGroup.Wait()

	}
	return arr
}

func randomPartition(arr []int, lo int, hi int) int {
	randomValue := rand.Intn((hi-lo)+1) + lo
	temp := arr[hi]
	arr[hi] = arr[randomValue]
	arr[randomValue] = temp
	return partition(arr, lo, hi)
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if arr[j] < pivot {
			i++
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
	}
	temp := arr[i+1]
	arr[i+1] = arr[hi]
	arr[hi] = temp
	return i + 1
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	arr1 := arr[:len(arr)/2]
	arr2 := arr[len(arr)/2:]

	arr1 = mergesort(arr1)
	arr2 = mergesort(arr2)
	arr = merge(arr1, arr2)
	return arr
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

func mergesortConcurrent(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	arr1 := arr[:len(arr)/2]
	arr2 := arr[len(arr)/2:]

	WaitGroup := sync.WaitGroup{}
	WaitGroup.Add(2)

	go func() {
		arr1 = mergesortConcurrent(arr1)
		WaitGroup.Done()
	}()

	go func() {
		arr2 = mergesortConcurrent(arr2)
		WaitGroup.Done()
	}()
	WaitGroup.Wait()

	arr = merge(arr1, arr2)
	return arr
}

func printArray(arr []int) {
	fmt.Printf("%v", arr)
}
