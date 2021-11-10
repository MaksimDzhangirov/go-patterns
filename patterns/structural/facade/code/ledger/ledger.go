package ledger

import "fmt"

type ledger struct {
}

func NewLedger() *ledger {
	return &ledger{}
}

func (l *ledger) MakeEntry(accountID string, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
}
