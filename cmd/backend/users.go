// users model for the users table, has a direct dependancy on database
package main

// User data structure
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// Get retrieves one element from users
func (user User) Get() (User, error) {
	db := GetDB()
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", user.ID)
	err := row.Scan(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Add will add a new user
func (user User) Add() error {
	db := GetDB()
	stmt, err := db.Prepare("INSERT INTO users (id, first_name, last_name, email, created_at) VALUES (?, ?, ?, ?, ?)")
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
	stmt, err := db.Prepare("UPDATE users SET first_name = ?, last_name = ?, email = ?, created_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.ID)
	return err
}
