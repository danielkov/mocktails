package main

import "github.com/danielkov/mocktail/mocktail"

type User struct {
	Username string `mock:"username"`
	Password string `mock:"password"`
	Email string `mock:"email"`
	DateOfBirth string `mock:"dob"`
	Company string `mock:"company"`
	Pets string `mock:"number"`
}

func main() {
	mocktail.Mock(User)
}