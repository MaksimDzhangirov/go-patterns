package user

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/iterator/code/interfaces"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/iterator/code/model"
)

type userCollection struct {
	users []*model.User
}

func NewUserCollection(users []*model.User) *userCollection {
	return &userCollection{
		users: users,
	}
}

func (u *userCollection) CreateIterator() interfaces.Iterator {
	return &userIterator{
		users: u.users,
	}
}
