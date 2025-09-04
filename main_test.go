package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("zero size", func(t *testing.T) {
		data := generateRandomElements(0)
		assert.Len(t, data, 0, "ожидался пустой слайс, а получили %v", data)
	})

	t.Run("positive size", func(t *testing.T) {
		size := 10
		data := generateRandomElements(size)

		assert.Len(t, data, size, "ожидалась длина %d, а получили %d", size, len(data))

		for i, v := range data {
			assert.Greater(t, v, 0, "элемент слайса под индексом %d не положительный: %d", i, v)
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("one element", func(t *testing.T) {
		nums := []int{42}
		got := maximum(nums)
		assert.Equal(t, 42, got)
	})

	t.Run("multiple elements", func(t *testing.T) {
		nums := []int{3, 7, 2, 9, 5}
		got := maximum(nums)
		assert.Equal(t, 9, got)
	})

	t.Run("all equal elements", func(t *testing.T) {
		nums := []int{5, 5, 5, 5}
		got := maximum(nums)
		assert.Equal(t, 5, got)
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("small slice", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := maxChunks(nums)
		require.Equal(t, 3, got)
	})

	t.Run("large slice", func(t *testing.T) {
		nums := []int{10, 200, 30, 400, 50, 60, 70, 800, 90}
		got := maxChunks(nums)
		require.Equal(t, 800, got)
	})

	t.Run("all equal elements", func(t *testing.T) {
		nums := []int{5, 5, 5, 5, 5, 5, 5, 5}
		got := maxChunks(nums)
		require.Equal(t, 5, got)
	})
}
