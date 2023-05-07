package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	targetnum := rand.Int31n(100)
	var input uint32

	fmt.Println("Please input you number")
	for n, _ := fmt.Scanln(&input); n != 0; n, _ = fmt.Scanln(&input) {
		if input == uint32(targetnum) {
			fmt.Println("Bingo!")
			break
		} else if input < uint32(targetnum) {
			fmt.Println("Too Small:(")
		} else if input > uint32(targetnum) {
			fmt.Println("Too Large:(")
		}
		fmt.Println("Please input you number")
	}
	i := 0
	for i < 10 {
		i++
		fmt.Print(i, " ")
	}
	fmt.Println()

	wd := time.Now().Weekday().String()
	switch wd {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It is", wd, ", go to work!")
	default:
		fmt.Println("It is", wd, ", just take a rest!")
	}
}
