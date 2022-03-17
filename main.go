package main

import (
	"Structs/users"
	"fmt"
)

func main() {

	pedro := users.User{Id: 2, Name: "Pedro", Age: 21}
	john := users.User{Id: 3, Name: "john", Age: 21}

	marta := users.User{Id: 1, Name: "Marta", Age: 20}
	marta.SayHello()
	fmt.Printf("%+v", marta)

	marta.AddFriend(pedro)

	marta.AddFriend(john)

	marta.ListFriends()

	marta.RemoveFriend(john.Id)

	fmt.Println("================")

	marta.ListFriends()

}
