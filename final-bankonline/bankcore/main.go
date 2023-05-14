package bank

import (
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(num float64) error {
	if num <= 0 {
		return errors.New("the amount should greater than 0")
	}

	a.Balance += num

	return nil
}

func (a *Account) Withdraw(num float64) error {
	if num <= 0 {
		return errors.New("the amount should greater than 0")
	}

	if a.Balance < num {
		return errors.New("balance is not enough")
	}

	a.Balance -= num

	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(target *Account, num float64) error {
	a.Withdraw(num)
	target.Deposit(num)

	return nil
}
