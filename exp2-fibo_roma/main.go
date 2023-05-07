package main

import (
	"fmt"
	"strings"
)

var special_roma = map[string]int{
	"CM": 900,
	"CD": 400,
	"XC": 90,
	"XL": 40,
	"IX": 9,
	"IV": 4,
}

var roma_table = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func fibo(rank int) int {
	if rank == 1 || rank == 2 {
		return 1
	} else {
		return fibo(rank-1) + fibo(rank-2)
	}
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d ", fibo(i))
	}
	fmt.Println()

	var roma_number string
	var s string
	result := 0
	fmt.Println("Please input a roma number:")
	fmt.Scanln(&roma_number)
	s = strings.Clone(roma_number)
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("[Error]", handler)
		}
	}()
	for len(s) != 0 {
		if len(s) >= 2 {
			target := s[:2]
			if roma_table[string(target[0])] < roma_table[string(target[1])] {
				v, err := roma_table[string(target[0])]
				if err {
					result -= v
				} else {
					panic("Unsupported Roma Number Format!")
				}
			} else {
				v, err := roma_table[string(target[0])]
				if err {
					result += v
				} else {
					panic("Unsupported Roma Number Format!")
				}
			}
			s = s[1:]
		} else {
			v, err := roma_table[s]
			if err {
				result += v
				s = s[1:]
			} else {
				panic("Unsupported Roma Number Format!")
			}
		}
	}
	fmt.Println("Result:", result)
}
