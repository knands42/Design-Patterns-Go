package template_methods

import "fmt"

func PlayGame(start, takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}
