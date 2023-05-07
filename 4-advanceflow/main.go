package main

import "fmt"

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}

func main() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}

//	func main() {
//		for i := 1; i <= 4; i++ {
//			// defer will be executed when the parent function is complete
//			// The i in defer will pop from a stack, LIFO(Last In, First Out)
//			// The output is:
//			// // regular 1
//			// // regular 2
//			// // regular 3
//			// // regular 4
//			// // deferred -4
//			// // deferred -3
//			// // deferred -2
//			// // deferred -1
//			defer fmt.Println("deferred", -i)
//			fmt.Println("regular", i)
//		}
//	}
