package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomZero(t *testing.T) {
	res := generateRandomElements(0)
	assert.Nil(t, res)
}

func TestGenerateRandomNegative(t *testing.T) {
	res := generateRandomElements(-1)
	assert.Nil(t, res)
}

func TestGenerateRandomLen(t *testing.T) {
	res := generateRandomElements(1)
	assert.Len(t, res, 1)
}
func TestGenerateRandomQuantity(t *testing.T) {
	rnd := rand.Intn(SIZE - 1)
	res := generateRandomElements(rnd)
	assert.Equal(t, rnd, len(res))
}

func TestMaximumChunksNil(t *testing.T) {
	res := maxChunks(nil)
	assert.Empty(t, res)
}

func TestMaximumChunksOne(t *testing.T) {
	r := rand.Int()
	testSlice := []int{r}
	res := maxChunks(testSlice)
	assert.Equal(t, r, res)
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    int
		wantErr bool
	}{
		{
			name:  "пустой слайс",
			input: []int{},
		},
		{
			name:  "один элемент",
			input: []int{42},
			want:  42,
		},
		{
			name:  "все положительные числа",
			input: []int{1, 3, 2, 5, 4},
			want:  5,
		},
		{
			name:  "все отрицательные числа",
			input: []int{-10, -3, -7, -1},
			want:  -1,
		},
		{
			name:  "положительные и отрицательные числа",
			input: []int{-2, 0, 5, -1, 3},
			want:  5,
		},
		{
			name:  "повторяющиеся числа",
			input: []int{2, 7, 7, 3},
			want:  7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximum(tt.input)
			assert.Equal(t, tt.want, res)
		})
	}
}
