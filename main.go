package main

import "fmt"

func main() {

	board := [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}

	gol := Game{board: board}
	gol.print()
	fmt.Println("new board ")
	gol.play()

}
