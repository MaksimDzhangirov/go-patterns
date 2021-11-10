package wallet

import "fmt"

type wallet struct {
	balance int
}

func NewWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Println("Wallet balance is sufficient")
	w.balance = w.balance - amount
	return nil
}
