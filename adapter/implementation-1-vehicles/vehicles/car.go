package vehicles

import "fmt"

type Car struct{}

func (c *Car) DriveToDestination() {
	fmt.Println("car is driving to destination.")
}
