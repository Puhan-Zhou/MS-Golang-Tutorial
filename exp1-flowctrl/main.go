package main

import (
	emergenum "exp1-flowctrl/emerge_num"
	"exp1-flowctrl/fizzbuzz"
	"exp1-flowctrl/prime"
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100)
	fizzbuzz.Fizzbuzz(target)
	prime.IsPrime(target)
	fmt.Println("Input your number:")
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("[Panic]", handler)
		}
	}()
	for {
		fmt.Scanln(&target)
		emergenum.IsEmerge(target)
	}
}
