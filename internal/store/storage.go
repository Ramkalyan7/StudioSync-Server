package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct{
	User interface{
		CreateUser(context.Context,*User)(*User,error)
		GetUserWithEmail(context.Context,string)(User,error)
	}
}


func NewStorage(db *sql.DB)Storage{
	return Storage{
		User: &UserStore{db},
	}
}