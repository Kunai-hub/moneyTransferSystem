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
		Transactions: make([]Transaction, 5, 10),
	}
}

func (ps *PaymentSystem) AddUser(user *User) {
	ps.Users[user.ID] = user
	fmt.Printf("User with ID %s added in PaymentSystem!\n", user.ID)
}

func (ps *PaymentSystem) AddTransaction(transaction *Transaction) {
	ps.Transactions = append(ps.Transactions, *transaction)
	fmt.Printf("Transaction from %s to %s users added in queue\n", transaction.FromID, transaction.ToID)
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
		Balance: 68.5,
	}
	paymentSystem.AddUser(user2)
	transaction1 := &Transaction{
		FromID: user1.ID,
		ToID:   user2.ID,
		Amount: 50.5,
	}
	paymentSystem.AddTransaction(transaction1)
	transaction2 := &Transaction{
		FromID: user2.ID,
		ToID:   user1.ID,
		Amount: 18,
	}
	paymentSystem.AddTransaction(transaction2)
	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)

	for i := range paymentSystem.Transactions {
		fmt.Printf("Transaction from %s to %s. Sum: %.2f", paymentSystem.Transactions[i].FromID, paymentSystem.Transactions[i].ToID, paymentSystem.Transactions[i].Amount)
		fmt.Println(user1.Balance)
		fmt.Println(user2.Balance)
	}
	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)

	// if err := user1.Withdraw(300); err != nil {
	// 	fmt.Printf("Error for %s: %s\n", user1.Name, err)
	// }
	// if err := user2.Withdraw(5.19); err != nil {
	// 	fmt.Printf("Error for %s: %s\n", user2.Name, err)
	// }
	// fmt.Println(user1.Balance)
	// fmt.Println(user2.Balance)
}
