package main

import (
	"errors"
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

	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000) + 1 // от 1 до 1000
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("error: func maximum - slice is null")
	}

	max := nums[0]
	for _, v := range nums[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.

func maxChunks(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("error: func maxChunks - slice is null")
	}

	chunkSize := (len(nums) + CHUNKS - 1) / CHUNKS // округляем вверх
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		// если диапазон пустой (могло получиться на маленьких слайсах) — пропускаем
		if start >= len(nums) {
			wg.Done()
			continue
		}

		go func(idx int, part []int) {
			defer wg.Done()
			localMax, _ := maximum(part) // используем однопоточный максимум
			maxValues[idx] = localMax
		}(i, nums[start:end])
	}

	wg.Wait()

	// ищем максимум среди максимумов
	finalMax, _ := maximum(maxValues)
	return finalMax, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	if len(data) == 0 {
		fmt.Println("Слайс пуст, выход.")
		return
	}
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max, err := maximum(data)
	elapsed := time.Since(start).Microseconds()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	fmt.Printf("Ищем максимальное значение в %d потоков...\n", CHUNKS)
	start = time.Now()
	max, err = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
