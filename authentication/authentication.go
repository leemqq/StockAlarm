package authentication

import (
	"bufio"
	"fmt"
	"os"
)

// User represents a user with a username and password.
type User struct {
	Username string
	Password string
}

// predefined users for simplicity
var users = []User{
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Please fill the password", name)
	return message
}

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func validate(username, password string) bool {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

// Login validates the username and password.
func Login(username string, password string) bool {
	// Prompt for username.

	if validate(username, password) {
		return true
	}
	return false
}
