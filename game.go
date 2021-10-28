package main

import "fmt"

type Game struct {
	board [5][5]int
}

func (g Game) print() {
	for _, row := range g.board {
		fmt.Println(row)
	}
}

func (g Game) getNeighbours(col int, row int) int {
	count := 0

	neighbours := []int{-1, 0, 1}

	for _, rowMod := range neighbours {
		for _, colMod := range neighbours {
			if rowMod == 0 && colMod == 0 {
				break
			}

			rowPosition := row + rowMod
			colPosition := col + colMod
			if g.isValidPosition(colPosition, rowPosition) {
				if g.board[rowPosition][colPosition] == 1 {
					count++
				}
			}
		}
	}

	return count
}

func (g Game) isValidPosition(col int, row int) bool {
	maxRow := len(g.board)
	maxCol := len(g.board[0])

	if col < 0 {
		return false
	}
	if col >= maxCol {
		return false
	}
	if row < 0 {
		return false
	}
	if row >= maxRow {
		return false
	}
	return true
}

func isAlive(startState bool, aliveNeighbours int) bool {

	if startState {
		if aliveNeighbours < 2 {
			return false
		}
		if aliveNeighbours > 3 {
			return false
		}
	} else if aliveNeighbours != 3 {
		return false
	}
	return true

}

func (g Game) play() {
	blackBoard := [5][5]int{}

	for i := range g.board {
		for j := range g.board {
			aliveNeighbours := g.getNeighbours(j, i)
			startState := g.board[j][i] != 0
			if isAlive(startState, aliveNeighbours) {
				blackBoard[j][i] = 1
			}
		}
	}

	newBoard := Game{board: blackBoard}

	newBoard.print()

}
