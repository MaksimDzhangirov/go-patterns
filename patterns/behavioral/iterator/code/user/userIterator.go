package user

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/iterator/code/model"

type userIterator struct {
	index int
	users []*model.User
}

func (u *userIterator) HasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *userIterator) GetNext() *model.User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
