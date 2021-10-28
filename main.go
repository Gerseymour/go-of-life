package main

import "fmt"

func main() {
	game := Game{board: board}

	fmt.Println("Initial Board")
	game.print()

	for i := 1; i != 0; i++ {
		game = game.nextStage()
		fmt.Printf("\nBoard %d:\n", i)
		game.print()

		if game.endOfGame {
			fmt.Println("The Game of Life has ended!")
			return
		}

		fmt.Println("Press Enter to continue to next stage:")
		fmt.Scanln()

	}
}
