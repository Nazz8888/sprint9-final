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
		return nil
	}
	if size == 1 {
		randomSlice := []int{rand.Int()}
		return randomSlice
	}
	randomSlice := make([]int, size)
	for i := 0; i < size; i++ {
		randomSlice[i] = rand.Int()
	}
	return randomSlice

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]

	}
	maxValue := data[0]
	for _, v := range data {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		res := maximum(data)
		return res
	}

	var wg sync.WaitGroup

	maxSlice := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		lenPart := len(data) / CHUNKS
		iStart := i * lenPart
		iFinish := iStart + lenPart
		part := data[iStart:iFinish]
		wg.Add(1)

		go func(localPart []int, i int) {
			defer wg.Done()
			maxInPart := maximum(part)
			maxSlice[i] = maxInPart
		}(part, i)
	}
	wg.Wait()
	maxValue := maximum(maxSlice)
	return maxValue
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	timeStart := time.Now()
	maxOneFlow := maximum(data)
	timeFinish := time.Since(timeStart)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxOneFlow, timeFinish.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	timeStart = time.Now()
	maxMultiFlow := maxChunks(data)
	timeFinish = time.Since(timeStart)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxMultiFlow, timeFinish.Microseconds())
}
