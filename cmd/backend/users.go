// users model for the users table, has a direct dependancy on database
package main

import (
	"fmt"
	"regexp"
	"time"
)

// User data structure
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// GetUsers retrieves one element from users
func GetUsers() ([]User, error) {
	users := []User{}
	db := GetDB()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Get retrieves one element from users
func (user User) Get() (User, error) {
	db := GetDB()
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", user.ID)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Add will add a new user
func (user User) Add() error {
	db := GetDB()
	stmt, err := db.Prepare("INSERT INTO users (first_name, last_name, email) VALUES ( ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email)
	return err
}

// Delete will remove a user
func (user User) Delete() error {
	db := GetDB()
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.ID)
	return err
}

// Update will update an existing user
func (user User) Update() error {
	db := GetDB()

	// ? Should we update timestamp upon update?
	t := time.Now().UTC()
	user.CreatedAt = t.Format("2006-01-02 15:04:05")

	stmt, err := db.Prepare("UPDATE users SET first_name = ?, last_name = ?, email = ?, created_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.ID)
	return err
}

// Validate will validate the user data
// ? should we contrain length here?
// Consider: https://github.com/go-playground/validator
func (user User) Validate() error {
	if user.FirstName == "" {
		return fmt.Errorf("first_name is required")
	}
	if user.LastName == "" {
		return fmt.Errorf("last_name is required")
	}
	if user.Email == "" {
		return fmt.Errorf("email is required")
	}
	if ok, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, user.Email); !ok {
		return fmt.Errorf("email is not valid")
	}
	return nil
}
