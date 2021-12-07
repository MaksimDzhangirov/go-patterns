package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/iterator/code/model"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/iterator/code/user"
)

func main() {
	user1 := model.NewUser("a", 30)
	user2 := model.NewUser("b", 20)
	userCollection := user.NewUserCollection([]*model.User{user1, user2})
	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		currentUser := iterator.GetNext()
		fmt.Printf("User is %+v\n", currentUser)
	}
}
