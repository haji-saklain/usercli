package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/haji-saklain/usercli/users"
)

func main() {
	// Define command-line flags for user management operations
	addUser := flag.Bool("adduser", false, "Add a new user")
	username := flag.String("username", "", "Username of the new user")
	email := flag.String("email", "", "Email of the new user")
	deleteUser := flag.Int("deleteuser", 0, "Delete a user by ID")
	updateUser := flag.Int("updateuser", 0, "Update a user by ID")
	getUser := flag.Int("getuser", 0, "Get a user by ID")

	// Parse the command-line flags
	flag.Parse()

	// Load users from file
	userFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println("Error opening users file:", err)
		return
	}
	defer userFile.Close()

	var userList []users.User
	if err := json.NewDecoder(userFile).Decode(&userList); err != nil && err.Error() != "EOF" {
		fmt.Println("Error loading users:", err)
		return
	}

	// Perform user management operations based on flags
	if *addUser {
		if *username == "" || *email == "" {
			fmt.Println("Please provide both username and email for adding a user.")
			return
		}
		userList = users.AddUser(userList, *username, *email)
	}
	if *deleteUser != 0 {
		userList = users.DeleteUser(userList, *deleteUser)
	}
	if *updateUser != 0 {
		if *username == "" || *email == "" {
			fmt.Println("Please provide both username and email for updating a user.")
			return
		}
		userList = users.UpdateUser(userList, *updateUser, *username, *email)
	}
	if *getUser != 0 {
		user := users.GetUserByID(userList, *getUser)
		if user == nil {
			fmt.Println("User not found")
		} else {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Username, user.Email)
		}
	}

	// Save users to file
	userFile, err = os.Create("users.json")
	if err != nil {
		fmt.Println("Error creating users file:", err)
		return
	}
	defer userFile.Close()
	if err := json.NewEncoder(userFile).Encode(userList); err != nil {
		fmt.Println("Error saving users:", err)
		return
	}
}
