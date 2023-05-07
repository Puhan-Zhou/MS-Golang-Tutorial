package main

import (
	"fmt"
)

func main() {
	// array
	var a1 [3]int
	a2 := [4]string{"Puhan", "Zhou"}
	a3 := [...]string{"Puhan", "Zhou", "2000-10-22"}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)

	// slice
	// len: the length
	// cap: from slice[0] to the end of basic array
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
	quarter := make([]string, 4)
	copy(quarter, quarter1)
	fmt.Println(quarter)

	// map
	student := map[string]int{
		"Puhan": 1,
		"Zhou":  2,
	}

	fmt.Println(student)
	fmt.Println("Puhan:", student["Puhan"])
	student["Test"] = 3
	fmt.Println(student)
	delete(student, "Test")
	fmt.Println(student)
	for name, index := range student {
		fmt.Println(name, index)
	}
}
