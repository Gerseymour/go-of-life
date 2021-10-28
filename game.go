package main

import (
	"fmt"
	"reflect"
)

type Game struct {
	board     [][]int
	endOfGame bool
}

func (g Game) numOfRows() int {
	return len(g.board)
}

func (g Game) numOfCols() int {
	return len(g.board[0])
}

func (g Game) print() {
	for _, row := range g.board {
		fmt.Println(row)
	}
}

func (g Game) getNeighbours(col int, row int) int {
	count := 0

	neighbours := [][]int{
		{-1, -1}, {-1, 0}, {-1, +1},
		{0, -1}, {0, +1},
		{+1, -1}, {+1, 0}, {+1, +1},
	}

	for _, neighbour := range neighbours {
		x := neighbour[0] + row
		y := neighbour[1] + col

		if g.isValidPosition(y, x) && g.board[x][y] == 1 {
			count++
		}
	}

	return count
}

func (g Game) isValidPosition(col int, row int) bool {
	validCol := (col >= 0) && (col < g.numOfCols())
	validRow := (row >= 0) && (row < g.numOfRows())

	return validCol && validRow
}

func (g Game) nextStage() Game {
	newBoard := makeBlankBoard(g.numOfRows(), g.numOfCols())

	for row := range g.board {
		for col := range g.board[0] {
			aliveNeighbours := g.getNeighbours(col, row)
			isCurrentCellAlive := g.board[row][col] != 0

			if isAlive(isCurrentCellAlive, aliveNeighbours) {
				newBoard[row][col] = 1
			}
		}
	}

	newGame := Game{board: newBoard, endOfGame: areBoardsEqual(g.board, newBoard)}

	return newGame
}

func isAlive(isCurrentCellAlive bool, aliveNeighbours int) bool {
	return (isCurrentCellAlive && aliveNeighbours == 2) || (aliveNeighbours == 3)
}

func makeBlankBoard(numOfRows int, numOfCols int) [][]int {
	blankBoard := make([][]int, numOfRows)

	for i := range blankBoard {
		blankBoard[i] = make([]int, numOfCols)
	}

	return blankBoard
}

func areBoardsEqual(board1 [][]int, board2 [][]int) bool {
	return reflect.DeepEqual(board1, board2)
}
