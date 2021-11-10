package wallet

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/facade/code/account"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/facade/code/interfaces"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/facade/code/ledger"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/facade/code/notification"
)

type walletFacade struct {
	account      interfaces.Account
	wallet       interfaces.Wallet
	securityCode interfaces.SecurityCode
	notification interfaces.Notification
	ledger       interfaces.Ledger
}

func NewWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacade := &walletFacade{
		account:      account.NewAccount(accountID),
		securityCode: account.NewSecurityCode(code),
		wallet:       NewWallet(),
		notification: notification.NewNotification(),
		ledger:       ledger.NewLedger(),
	}
	fmt.Println("Account created")
	return walletFacade
}

func (w *walletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(accountID, "debit", amount)
	return nil
}
