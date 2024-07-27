package main

import "fmt"

// type name should start with capital letters
type User struct {
	Email    string
	Password string
}

// type name should start with capital letters
type UserRepo struct {
	//DB is a user database slice of the User struct type to store the emails and passwords of users
	DB []User
}

// Register method adds a new user to the repository
func (r *UserRepo) Register(u User) {
	if u.Email == "" || u.Password == "" {
		fmt.Println("register failed")
	}
	// Append the new user to the DB slice
	r.DB = append(r.DB, u)
}

// Login method checks if the user's credentials exist in the repository
func (r *UserRepo) Login(u User) string {
	if u.Email == "" || u.Password == "" {
		fmt.Println("login failed")
	}

	//instead of using "us" as variable we used "existinguser" to clearly define what that variable is
	for _, existinguser := range r.DB {
		if existinguser.Email == u.Email && existinguser.Password == u.Password {
			return "auth token"
		}
	}

	return ""
}
