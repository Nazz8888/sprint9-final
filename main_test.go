package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
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

func TestMaximumNil(t *testing.T) {
	res := maximum(nil)
	assert.Empty(t, res)
}

func TestMaximumNOne(t *testing.T) {
	r := rand.Int()
	testSlice := []int{r}
	res := maximum(testSlice)
	assert.Equal(t, r, res)
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
