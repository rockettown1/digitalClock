package main

import (
	"fmt"
	"time"

	"github.com/rockettown1/learngo/digitalClock/digits"
	"github.com/rockettown1/learngo/digitalClock/screenutils"
)

func main() {
	//sets up infinite loop
	for {
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]digits.Placeholder{
			//an int when divided will remove the fractional part, use the modulus to get the remainder after dividing by 10
			digits.Numbers[hour/10], digits.Numbers[hour%10],
			digits.Seperator,
			digits.Numbers[min/10], digits.Numbers[min%10],
			digits.Seperator,
			digits.Numbers[sec/10], digits.Numbers[sec%10],
		}

		//as we are redrawing the clock every second we need to clear and move the cursor back to the top.
		screenutils.Clear()
		screenutils.MoveTopLeft()
		fmt.Println("Welcome to a basic Golang clock for the terminal.")
		fmt.Println()
		//this will take all the digits in the clock array, and loop through each line of the digit - one line at a time.
		//so all the first lines will be printed to the screen, then all the second lines etc
		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], " ")
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println("Please press [Control + c] to exit.")
		fmt.Println("Written by Dan Krishnan")
		time.Sleep(time.Second)
	}

}
