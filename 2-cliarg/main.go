package main

import (
	"2-cliarg/swap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input 2 args")
		os.Exit(-1)
	}
	num1, _ := strconv.Atoi(os.Args[1])
	num2, _ := strconv.Atoi(os.Args[2])

	fmt.Println(num1, num2)

	swap.Swap_ng(num1, num2)
	fmt.Println("NG Main:", num1, num2)
	swap.Swap_ok(&num1, &num2)
	fmt.Println("OK Main:", num1, num2)
}

// func swap_ng(num1, num2 int) {
// 	num1 = num1 + num2
// 	num2 = num1 - num2
// 	num1 = num1 - num2
// 	fmt.Println("NG Local:", num1, num2)
// }

// func swap_ok(num1, num2 *int) {
// 	*num1 = *num1 + *num2
// 	*num2 = *num1 - *num2
// 	*num1 = *num1 - *num2
// 	fmt.Println("OK Local:", *num1, *num2)
// }
