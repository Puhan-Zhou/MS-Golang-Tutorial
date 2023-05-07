package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var (
		firstName, lastName string = "Puhan", "Zhou"
		age                 int    = 23
	)
	const (
		birthday string = "10-22"
	)
	five := "5"
	fmt.Println("Hello", firstName, lastName, "Age: ", age)
	firstName = "Puu"
	fmt.Println("Hello", firstName, lastName, "Age: ", age)
	fmt.Println("My birthday is", birthday)
	fmt.Println(math.MaxInt64)
	num, err := strconv.Atoi(five)
	if err != nil {
		fmt.Println("Error!")
	}
	fmt.Println(num)
}
