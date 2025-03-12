package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password []byte `json:"password"`
}

type UserStore struct{
	db *sql.DB
}


func (u *UserStore)CreateUser(ctx context.Context,user *User)(*User,error){
	//write query to store user in go 
	query := `INSERT INTO users (username,email,password) VALUES ($1,$2,$3) RETURNING username,email`

	//excecute query on the database
	ctx,cancel:=context.WithTimeout(ctx,QueryTimeoutDuration)
	defer cancel()

	err := u.db.QueryRowContext(ctx,query,user.Username,user.Email,user.Password).Scan(&user.Username,&user.Email)

	//send response.

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return nil,ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return nil,ErrDuplicateUsername
		default:
			return nil,err
		}
	}
	return user,nil
}


func (u *UserStore)GetUserWithEmail(ctx context.Context,email string)(User,error){

	var user User;
	//validate the email
	if email==""{
		return user,fmt.Errorf("email is empty");
	}
	//get the user with email
	err := u.db.QueryRowContext(ctx, "SELECT username,email,password FROM users WHERE email = $1", email).Scan(
		&user.Username,
		&user.Email,
		&user.Password,
	)
	if err!=nil{
		switch err {
		case sql.ErrNoRows:
			return user,ErrNotFound
		default:
			return user,err
		}
	}
	//return the user
	return user,nil;
}
