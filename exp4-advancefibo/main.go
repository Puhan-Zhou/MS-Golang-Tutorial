package main

import (
	"fmt"
	"time"
)

func fib(number int) int {
	x, y := 1, 1
	for i := 0; i < number; i++ {
		x, y = y, x+y
	}

	return x
}

func cal_fibo(ch chan<- int) {
	i := 0
	for {
		ch <- fib(i)
		i++
	}
}

func get_input(ch chan<- string) {
	var input string
	for {
		fmt.Scanf("%s\n", &input)
		ch <- input
	}
}

func main() {
	start := time.Now()

	fib_ch := make(chan int)
	input_ch := make(chan string)

	go cal_fibo(fib_ch)
	go get_input(input_ch)

	fmt.Println(<-fib_ch)

	for {
		if input := <-input_ch; input == "quit" {
			break
		} else if input == "" {
			fmt.Println(<-fib_ch)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
