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

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users        map[string]*User
	Transactions []Transaction
}

func NewPaymentSystem() *PaymentSystem {
	return &PaymentSystem{
		Users:        make(map[string]*User),
		Transactions: make([]Transaction, 0, 5),
	}
}

func (ps *PaymentSystem) AddUser(user *User) {
	ps.Users[user.ID] = user
	fmt.Printf("User with ID %s added in PaymentSystem!\n", user.ID)
}

func (ps *PaymentSystem) AddTransaction(transaction *Transaction) {
	ps.Transactions = append(ps.Transactions, *transaction)
	fmt.Printf("Transaction from %s to %s users with amount %.2f added in queue\n", transaction.FromID, transaction.ToID, transaction.Amount)
}

func (ps *PaymentSystem) ProcessingTransactions(transaction Transaction) error {
	if _, ok := ps.Users[transaction.FromID]; !ok {
		errString := fmt.Sprintf("user with ID %s not found in payment system", transaction.FromID)
		return errors.New(errString)
	}
	if _, ok := ps.Users[transaction.ToID]; !ok {
		errString := fmt.Sprintf("user with ID %s not found in payment system", transaction.ToID)
		return errors.New(errString)
	}
	if err := ps.Users[transaction.FromID].Withdraw(transaction.Amount); err != nil {
		return err
	}
	ps.Users[transaction.ToID].Deposit(transaction.Amount)

	return nil
}

func main() {
	paymentSystem := NewPaymentSystem()
	user1 := &User{
		ID:      "user1",
		Name:    "John",
		Balance: 100,
	}
	paymentSystem.AddUser(user1)
	user2 := &User{
		ID:      "user2",
		Name:    "Regina",
		Balance: 50.5,
	}
	paymentSystem.AddUser(user2)
	transaction1 := &Transaction{
		FromID: user1.ID,
		ToID:   user2.ID,
		Amount: 49.5,
	}
	paymentSystem.AddTransaction(transaction1)
	transaction2 := &Transaction{
		FromID: user2.ID,
		ToID:   user1.ID,
		Amount: 0.5,
	}
	paymentSystem.AddTransaction(transaction2)
	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)

	for _, t := range paymentSystem.Transactions {
		if err := paymentSystem.ProcessingTransactions(t); err != nil {
			fmt.Printf("Transaction error: %s\n", err)
			continue
		}
		fmt.Printf("Transaction from %s to %s. Sum: %.2f\n", t.FromID, t.ToID, t.Amount)
		fmt.Println(user1.Balance)
		fmt.Println(user2.Balance)
	}
}
