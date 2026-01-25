package models

import (
	"errors"
	"go_api/db"
	"go_api/utils"
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

func (user *User) Login() (string, error) {

	query := `
	SELECT id,password FROM users WHERE email = ?
	`

	row := db.DB.QueryRow(query, user.Email)

	var storedPassword string
	err := row.Scan(&user.ID, &storedPassword)

	if err != nil {
		return "", errors.New("Credentials Invalid")
	}

	err = utils.CheckPassword(user.Password, storedPassword)
	if err != nil {
		return "", errors.New("Credentials Invalid")
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		return "", errors.New("Error creatiing a token")
	}

	return token, nil
}
