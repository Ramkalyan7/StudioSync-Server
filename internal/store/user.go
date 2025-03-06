package store

import (
	"context"
	"database/sql"
)

type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}


type UserStore struct{
	db *sql.DB
}


func (u *UserStore)CreateUser(context.Context,*User)error{
	//write query to store user in go 
	//send response.
	//do proper error handling by taking reference from tiago bhai
}
