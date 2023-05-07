package prime

import (
	"fmt"
	"math"
)

func IsPrime(target int) bool {
	if target == 2 || target == 3 {
		fmt.Println(target, "is prime")
		return true
	}

	mod_six := target % 6

	if mod_six != 5 && mod_six != 1 {
		fmt.Println(target, "is not a prime")
		return false
	} else {
		for num := 5; num <= int(math.Sqrt(float64(target))); num += 6 {
			if target%num == 0 || target%(num+2) == 0 {
				fmt.Println(target, "is not prime")
				return false
			}
		}
	}

	fmt.Println(target, "is a prime")
	return true
}
