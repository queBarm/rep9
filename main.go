package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	// Ставим seed, чтобы числа менялись при каждом запуске
	rand.Seed(time.Now().UnixNano())

	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in chunks.
func maxChunks(data []int) int {
	n := len(data)
	if n == 0 {
		return 0
	}

	workers := CHUNKS
	if n < workers {
		workers = n
	}

	// равномерное деление
	q, r := n/workers, n%workers
	maxValues := make([]int, workers)

	var wg sync.WaitGroup
	wg.Add(workers)

	start := 0
	for i := 0; i < workers; i++ {
		size := q
		if i < r {
			size++
		}
		end := start + size
		part := data[start:end]

		go func(idx int, p []int) {
			defer wg.Done()
			maxValues[idx] = maximum(p)
		}(i, part)

		start = end
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
