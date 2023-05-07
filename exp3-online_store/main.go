package main

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) ChangeName(FirstName string, LastName string) {
	a.FirstName = FirstName
	a.LastName = LastName
}

type Employee struct {
	amount float64
	Account
}

func (e *Employee) AddCredits(number float64) (float64, error) {
	if number > 0 {
		e.amount += number
		return e.amount, nil
	}
	return 0, errors.New("Invalid add!\n")
}

func (e *Employee) RemoveCredits(number float64) (float64, error) {
	if number > 0 {
		if e.amount >= number {
			e.amount -= number
			return e.amount, nil
		} else {
			return 0, errors.New("The amount that you want to remove is larger than credits")
		}
	}
	return 0, errors.New("Invalid remove!\n")
}

func (e Employee) CheckCredits() float64 {
	return e.amount
}

func (e Employee) ChangeName(first_name string, last_name string) {
	e.Account.FirstName = first_name
	e.Account.LastName = last_name
}

func (e Employee) String() string {
	return fmt.Sprintf("%v %v's credits is %v", e.Account.FirstName, e.Account.LastName, e.amount)
}

func CreateEmployee(FirstName, LastName string, amount float64) (*Employee, error) {
	return &Employee{amount: amount, Account: Account{FirstName: FirstName, LastName: LastName}}, nil
}

func main() {
	bruce, _ := CreateEmployee("Bruce", "Lee", 500)
	fmt.Println(bruce.CheckCredits())
	credits, err := bruce.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	_, err = bruce.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	}

	bruce.ChangeName("Mark", "Lee")

	fmt.Println(bruce)
}
