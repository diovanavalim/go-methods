package main

import (
	"fmt"
	"methods/user"
)

func main() {
	user1 := user.CreateUser("Fabiana", "Souza", 21, "fabianasouze@email.com", "123")

	name := user1.GetName()
	lastName := user1.GetLastName()
	age := user1.GetAge()
	email := user1.GetEmail()

	fmt.Printf("User's name: %s\n", name)
	fmt.Printf("User's last name: %s\n", lastName)
	fmt.Printf("User's age: %d\n", age)
	fmt.Printf("User's e-mail: %s\n", email)

	user1.SetName("Fábio")
	user1.SetLastName("Oliveira")
	user1.SetAge(22)
	user1.SetEmail("fabiooliveira@email.com")

	fmt.Printf("User's updated name: %s\n", user1.GetName())
	fmt.Printf("User's updated last name: %s\n", user1.GetLastName())
	fmt.Printf("User's updated age: %d\n", user1.GetAge())
	fmt.Printf("User's updated e-mail: %s\n", user1.GetEmail())
}
