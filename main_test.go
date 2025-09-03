package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("zero size", func(t *testing.T) {
		data := generateRandomElements(0)
		if len(data) != 0 {
			t.Errorf("ожидался пустой слайс, а получили %v", data)
		}
	})

	t.Run("positive size", func(t *testing.T) {
		size := 10
		data := generateRandomElements(size)

		if len(data) != size {
			t.Errorf("ожидалась длина %d, а получили %d", size, len(data))
		}

		for i, v := range data {
			if v <= 0 {
				t.Errorf("элемент слайса под индексом %d не положительный: %d", i, v)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		_, err := maximum([]int{})
		if err == nil {
			t.Errorf("ожидалась ошибка для пустого слайса")
		}
	})

	t.Run("one element", func(t *testing.T) {
		nums := []int{42}
		got, err := maximum(nums)
		if err != nil {
			t.Fatalf("не ожидалась ошибка, получили %v", err)
		}
		if got != 42 {
			t.Errorf("ожидалось 42, получили %d", got)
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		nums := []int{3, 7, 2, 9, 5}
		got, err := maximum(nums)
		if err != nil {
			t.Fatalf("не ожидалась ошибка, получили %v", err)
		}
		if got != 9 {
			t.Errorf("ожидалось 9, получили %d", got)
		}
	})

	t.Run("all equal elements", func(t *testing.T) {
		nums := []int{5, 5, 5, 5}
		got, err := maximum(nums)
		if err != nil {
			t.Fatalf("не ожидалась ошибка, получили %v", err)
		}
		if got != 5 {
			t.Errorf("ожидалось 5, получили %d", got)
		}
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		_, err := maxChunks([]int{})
		if err == nil {
			t.Errorf("ожидалась ошибка для пустого слайса")
		}
	})

	t.Run("small slice", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got, err := maxChunks(nums)
		if err != nil {
			t.Fatalf("не ожидалась ошибка, получили %v", err)
		}
		if got != 3 {
			t.Errorf("ожидалось 3, получили %d", got)
		}
	})

	t.Run("large slice", func(t *testing.T) {
		nums := []int{10, 200, 30, 400, 50, 60, 70, 800, 90}
		got, err := maxChunks(nums)
		if err != nil {
			t.Fatalf("не ожидалась ошибка, получили %v", err)
		}
		if got != 800 {
			t.Errorf("ожидалось 800, получили %d", got)
		}
	})

	t.Run("all equal elements", func(t *testing.T) {
		nums := []int{5, 5, 5, 5, 5, 5, 5, 5}
		got, err := maxChunks(nums)
		if err != nil {
			t.Fatalf("не ожидалась ошибка, получили %v", err)
		}
		if got != 5 {
			t.Errorf("ожидалось 5, получили %d", got)
		}
	})
}
