package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNeighbours(t *testing.T) {
	board := [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}

	testgol := Game{board: board}

	result1 := testgol.getNeighbours(3, 2)
	result2 := testgol.getNeighbours(1, 2)
	result3 := testgol.getNeighbours(0, 4)
	result4 := testgol.getNeighbours(4, 0)

	assert.Equal(t, 3, result1, "They should be equal!")
	assert.Equal(t, 5, result2, "They should be equal!")
	assert.Equal(t, 1, result3, "They should be equal!")
	assert.Equal(t, 0, result4, "They should be equal!")

}

func TestValidPosition(t *testing.T) {
	board := [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}

	testgol := Game{board: board}

	result1 := testgol.isValidPosition(0, 0)
	result2 := testgol.isValidPosition(-1, 0)
	result3 := testgol.isValidPosition(6, 0)
	result4 := testgol.isValidPosition(0, -1)
	result5 := testgol.isValidPosition(0, 6)

	assert.True(t, result1)
	assert.False(t, result2)
	assert.False(t, result3)
	assert.False(t, result4)
	assert.False(t, result5)

}

func TestIsAlive(t *testing.T) {

	result1 := isAlive(true, 1)
	result2 := isAlive(true, 0)
	result3 := isAlive(true, 2)
	result4 := isAlive(true, 3)
	result5 := isAlive(true, 4)
	result6 := isAlive(true, 5)
	result7 := isAlive(false, 1)
	result8 := isAlive(false, 4)
	result9 := isAlive(false, 3)

	assert.False(t, result1)
	assert.False(t, result2)
	assert.True(t, result3)
	assert.True(t, result4)
	assert.False(t, result5)
	assert.False(t, result6)
	assert.False(t, result7)
	assert.False(t, result8)
	assert.True(t, result9)

}
