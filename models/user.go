package models

import (
	"Rest_Api_Go/db"
	"Rest_Api_Go/utils"
	"errors"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close() //if this is closed prematurely, error will occur
	hashedPassword, err := utils.HashPasswords(u.Password)

	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//update user id
	ID, err := res.LastInsertId()
	u.ID = ID
	return err
}

func (u *User) ValidateLogin() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPswd string
	err := row.Scan(&u.ID, &retrievedPswd)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPswd)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
