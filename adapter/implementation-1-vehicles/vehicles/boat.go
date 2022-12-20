package vehicles

import "fmt"

type Boat struct{}

func (c *Boat) NavigateToDestination() {
	fmt.Println("boat is navigating to island.")
}
