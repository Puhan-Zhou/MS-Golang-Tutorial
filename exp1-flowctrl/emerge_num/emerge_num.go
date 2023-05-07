package emergenum

import "fmt"

func IsEmerge(target int) {
	if target < 0 {
		panic("Your number is a negative number, emergency!")
	} else {
		fmt.Println("The number is permitted, input your number again:")
	}
}
