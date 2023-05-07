package swap

import "fmt"

func Swap_ng(num1, num2 int) {
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println("NG Local:", num1, num2)
}

func Swap_ok(num1, num2 *int) {
	*num1 = *num1 + *num2
	*num2 = *num1 - *num2
	*num1 = *num1 - *num2
	fmt.Println("OK Local:", *num1, *num2)
}
