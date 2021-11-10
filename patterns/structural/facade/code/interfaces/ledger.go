package interfaces

type Ledger interface {
	MakeEntry(string, string, int)
}
