package notification

import "fmt"

type notification struct {
}

func NewNotification() *notification {
	return &notification{}
}

func (n *notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
