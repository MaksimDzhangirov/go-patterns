package db

type defaultDatabase struct {
}

func NewDefaultDatabase() *defaultDatabase {
	return &defaultDatabase{}
}

func (db *defaultDatabase) GetUser() string {
	return "user 1"
}
