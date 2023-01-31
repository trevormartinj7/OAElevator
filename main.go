package main

import (
	"fmt"
)

// This elevator object contains its current floor, how long the journey has taken, and an array of integers to keep track of floors it's visited
type elevator struct {
	currentFloor  int
	timeElapsed   int
	visitedFloors []int
}

// This is a method attached to the elevator struct. It can only be called when an elevator object exists.
func (elevator *elevator) visitNextFloor(nextFloor int, travelSpeed int) {

	//Checking to make sure we're actually traveling
	if elevator.currentFloor == nextFloor {
		fmt.Println("We're already on that floor. Enter the next floor, or -1 to end your journey.")
		return
	}

	//Calculating how many seconds it takes to reach the desired floor
	difference := elevator.currentFloor - nextFloor

	//Accounting for negatives when going down in level
	if difference < 0 {
		difference = -difference
	}

	//This allows us to change the const TRAVELSPEED to see how long different speeds of elevators would take
	travelTime := difference * travelSpeed

	//Incrementing total travel time with new value
	elevator.timeElapsed += travelTime

	//Keeping track of visited floors this allows for repeat visits of previous floors
	elevator.visitedFloors = append(elevator.visitedFloors, nextFloor)

	//Setting the elevator up for future floor requests
	fmt.Println("Great, traveling to ", nextFloor)
	fmt.Println("It took us ", travelTime, " seconds to reach that floor. Enter the next floor, or -1 to end your journey.")
	elevator.currentFloor = nextFloor
}

// Validating user input
func validateInput(floor int) bool {
	//Golang will assign a 0 to the floor if anything but an int is returned.
	//Because 0 is not a floor in most buildings, we can use it to check for valid integer inputs
	for floor != 0 {
		return false
	}
	return true
}

func main() {
	//Initializing my elevator struct with null values
	myElevator := elevator{}
	//Adjustable travel speed, though for this problem I was told to use a constant travel speed of 10
	const TRAVELSPEED = 10

	//Getting initial floor from user
	fmt.Println("Welcome to Trevor's elevator. What is the starting floor?")
	fmt.Scanln(&myElevator.currentFloor)

	// Validating user input
	for validateInput(myElevator.currentFloor) {
		fmt.Scanln(&myElevator.currentFloor)
		fmt.Println("Invalid input. Please input an integer.")
	}

	//Adding first floor to list of visited floors
	myElevator.visitedFloors = append(myElevator.visitedFloors, myElevator.currentFloor)
	fmt.Println("We'll start at floor", myElevator.currentFloor)
	fmt.Println("Let's begin! Keep entering floors to visit, or enter '-1' to end your journey.")

	//Continually visits new floors until the user is done
	for myElevator.currentFloor >= 0 {
		//Creates and assigns new variables from user input
		var newFloor int
		fmt.Scanln(&newFloor)

		//Validates input
		for validateInput(newFloor) {
			fmt.Scanln(&newFloor)
			fmt.Println("Invalid input. Please input an integer.")
		}

		//Checks for exit case when input is -1, then visits the next floor
		if newFloor >= 0 {
			myElevator.visitNextFloor(newFloor, TRAVELSPEED)
		} else {
			break
		}
	}

	//Outputting final stats
	fmt.Println("We finished traveling and visited ", len(myElevator.visitedFloors), " floors: ", myElevator.visitedFloors)
	fmt.Println("It took us a total of ", myElevator.timeElapsed, " seconds to travel all that way. Press enter to close the program!")

	//This pauses the program to wait for the user to read the output and press enter.
	fmt.Scanln(&myElevator.currentFloor)

}
