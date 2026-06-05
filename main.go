package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(sum float64) {
	u.Balance += sum
	fmt.Printf("Top up balance at %.2f. Current balance: %.2f\n", sum, u.Balance)
}

func (u *User) Withdraw(sum float64) error {
	if u.Balance < sum {
		return errors.New("insufficient funds!")
	}
	u.Balance -= sum
	fmt.Printf("Write-off from balance %.2f. Current balance: %.2f\n", sum, u.Balance)

	return nil
}

func main() {
	user1 := &User{
		ID:      "user1",
		Name:    "John",
		Balance: 100,
	}
	user2 := &User{
		ID:      "user2",
		Name:    "Regina",
		Balance: 68.5,
	}

	user1.Deposit(50.5)
	user2.Deposit(10.10)
	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)

	if err := user1.Withdraw(300); err != nil {
		fmt.Printf("Error for %s: %s\n", user1.Name, err)
	}
	if err := user2.Withdraw(5.19); err != nil {
		fmt.Printf("Error for %s: %s\n", user2.Name, err)
	}
	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)
}
