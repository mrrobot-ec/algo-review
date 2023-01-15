package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	t.Run("It should be successfully ordered when executing insertionSort function", func(t *testing.T) {
		list := []int{5, 2, 4, 6, 1, 3, -3, -2, -1, 0, 8, 97, -5, 23, 45, 87, 12, 434, 234, 32, 567}
		start := time.Now()
		insertionSort(&list)
		log.Printf("execution time %s\n", time.Since(start))
		expectedResult := []int{-5, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 8, 12, 23, 32, 45, 87, 97, 234, 434, 567}
		assert.Equal(t, expectedResult, list)
	})
}
