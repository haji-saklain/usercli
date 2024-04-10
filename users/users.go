package users

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// AddUser adds a new user to the user list
func AddUser(users []User, username, email string) []User {
	newID := generateNextID(users)
	newUser := User{
		ID:       newID,
		Username: username,
		Email:    email,
	}
	return append(users, newUser)
}

// DeleteUser deletes a user from the user list by ID
func DeleteUser(users []User, id int) []User {
	for i, user := range users {
		if user.ID == id {
			return append(users[:i], users[i+1:]...)
		}
	}
	return users
}

// UpdateUser updates a user's information in the user list by ID
func UpdateUser(users []User, id int, username, email string) []User {
	for i, user := range users {
		if user.ID == id {
			users[i].Username = username
			users[i].Email = email
			break
		}
	}
	return users
}

// GetUserByID retrieves a user from the user list by ID
func GetUserByID(users []User, id int) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil // Return nil if user is not found
}

func generateNextID(users []User) int {
	if len(users) == 0 {
		return 1
	}
	return users[len(users)-1].ID + 1
}
