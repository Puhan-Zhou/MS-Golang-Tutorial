package fizzbuzz

import (
	"fmt"
)

func Fizzbuzz(target int) {
	fmt.Println(target)
	if target%15 == 0 {
		fmt.Println("fizzbuzz")
	} else if target%3 == 0 {
		fmt.Println("fizz")
	} else if target%5 == 0 {
		fmt.Println("buzz")
	} else {
		fmt.Println(target)
	}
}
