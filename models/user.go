package models

import (
	"database/sql"
)

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) CreateUser(db *sql.DB) error {
	query := `INSERT INTO users (name, gender, email, phone_number, username, password) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, u.Name, u.Gender, u.Email, u.PhoneNumber, u.Username, u.Password)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	query := `SELECT id, name, gender, email, phone_number, username, password FROM users WHERE username = ?`
	row := db.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Gender, &user.Email, &user.PhoneNumber, &user.Username, &user.Password)
	return &user, err
}
