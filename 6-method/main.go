package main

import (
	"6-method/Animal"
	"fmt"
)

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) Say() {
	fmt.Println("My name is", d.name, ", Woof!")
}

func (d *Dog) Rename(name string) {
	d.name = name
}

func (d Dog) GetName() string {
	return d.name
}

func (c Cat) Say() {
	fmt.Println("My name is", c.name, ", Meow~")
}

func (c *Cat) Rename(name string) {
	c.name = name
}

func (c Cat) GetName() string {
	return c.name
}

func main() {
	dog := Dog{"Lulu"}
	cat := Cat{"Mimi"}

	check_animal(&dog)
	check_animal(&cat)
}

// Pointer can call the structure method
// but the structure need to transform to the pointer type when it call the pointer method
// so the implement for structure can be called by pointer
// but pointer call to the structrue method is not permitted
func check_animal(a Animal.Animal) {
	defer fmt.Println("=================================")
	a.Say()
	a.Rename("Animal")
	fmt.Println(a.GetName())
	a.Say()
}
