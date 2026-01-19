package models

import (
	"go_api/db"
	"go_api/utils"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {

	query := `
	INSERT INTO users(email, password)
	VALUES (?,?)
	`

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, user.Email, hashedPassword)

	result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func (user User) Login() error {

	var tempUser User

	query := `
	SELECT * FROM users WHERE email = ?
	`

	row := db.DB.QueryRow(query, user.Email)

	err := row.Scan(&tempUser.ID, &tempUser.Email, &tempUser.Password)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(tempUser.Password), []byte(user.Password))

	if err != nil {
		return err
	}

	return nil
}
