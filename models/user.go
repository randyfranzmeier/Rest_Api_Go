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

func (u *User) ChangePassword() error {
	//query to update the password for the given user
	query := "UPDATE users SET password = ? WHERE id = ?"
	//preparing the query for efficiency and to prevent sql injection attacks
	stmt, err := db.DB.Prepare(query)
	//if there's an error we don't want to continue
	if err != nil {
		return err
	}
	//make sure the query statement gets closed when it's not needed
	defer stmt.Close()
	//we want to encrypt the password for security reasons
	hashedPassword, err := utils.HashPasswords(u.Password)
	u.Password = hashedPassword
	//now the values should be ready to be passed in
	_, err = stmt.Exec(u.Password, u.ID)
	return err
}
